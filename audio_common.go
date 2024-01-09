package glimmer

import (
	"time"
	"sync"
)

// AudioBuffer lets you play sound.
type AudioBuffer struct {
	SamplesPerSecond int
	BitsPerSample    int
	ChannelCount     int

	buf []byte
    bufMutex sync.Mutex

	output audioOutput
    outputBufDuration time.Duration

    // ReadLenNotifier is used at the end of every time the audio device
    // consumes the buffer, sending the number of bytes consumed
    // in the buffer consumption
    ReadLenNotifier chan int
    bufUnderflows int

	prevReadLen int
	maxWaited time.Duration
}

// GetPrevReadLen returns the size of the last buffer chunk the system asked for in its audio callback
func (ab *AudioBuffer) GetPrevReadLen() int {
	return ab.prevReadLen
}

// GetMaxWaited is to learn the max amount WaitForPlaybackIfAhead has waited in between this and the last call of
// GetMaxWaited. It's for debugging.
func (ab *AudioBuffer) GetMaxWaited() time.Duration {
	out := ab.maxWaited
	ab.maxWaited = time.Duration(0)
	return out
}

// WaitForPlaybackIfAhead waits until one buffer chunk length is yet to be processed by the OS.
func (ab *AudioBuffer) WaitForPlaybackIfAhead() {
	start := time.Now()
	for ab.GetLenUnplayedData() > ab.prevReadLen {
		ab.prevReadLen = <-ab.ReadLenNotifier
	}
	audioDiff := time.Now().Sub(start)
	if audioDiff > ab.maxWaited {
		ab.maxWaited = audioDiff
	}
}

type OpenAudioBufferOptions struct {
    OutputBufDuration time.Duration
    SamplesPerSecond int
    BitsPerSample int
    ChannelCount int
}

// OpenAudioBuffer creates and returns a new playing buffer
func OpenAudioBuffer(opts OpenAudioBufferOptions) (*AudioBuffer, error) {
	ab := AudioBuffer{
		SamplesPerSecond: opts.SamplesPerSecond,
		BitsPerSample:    opts.BitsPerSample,
		ChannelCount:     opts.ChannelCount,
		outputBufDuration: opts.OutputBufDuration,
        ReadLenNotifier: make(chan int),
	}

	err := ab.output.init(&ab)
	if err != nil {
		return nil, err
	}

	// wait to be sure audio is ready, and to get exact len data
	ab.prevReadLen = <-ab.ReadLenNotifier

	return &ab, nil
}

func (ab *AudioBuffer) GetLatestUnderflowCount() int {
    ab.bufMutex.Lock()
    count := ab.bufUnderflows
    ab.bufUnderflows = 0
    ab.bufMutex.Unlock()
    return count
}

func (ab *AudioBuffer) GetLenUnplayedData() int {
    ab.bufMutex.Lock()
    bLen := len(ab.buf)
    ab.bufMutex.Unlock()
    return bLen
}

func (ab *AudioBuffer) Write(data []byte) (int, error) {
    ab.bufMutex.Lock()
    ab.buf = append(ab.buf, data...)
    ab.bufMutex.Unlock()
    return len(data), nil
}

func (ab *AudioBuffer) Read(buf []byte) (int, error) {
    ab.bufMutex.Lock()
    if len(buf) > len(ab.buf) {
        // fmt.Println("[glimmer-audio] buf underflow:", len(buf), "vs", len(ab.buf))
        ab.bufUnderflows++
    }
    n := copy(buf, ab.buf)
    ab.buf = ab.buf[n:]
    ab.bufMutex.Unlock()

    // fill with zero if wrote nothing, otherwise fill with last val to give the buffer a chance to avoid a click if it can catch up
    if n == 0 {
        for i := range buf {
            buf[i] = 0
        }
    } else {
        for i := n+1; i < len(buf); i++ {
            buf[i] = buf[i-1]
        }
    }

    select {
    case ab.ReadLenNotifier<-len(buf):
    default:
    }

    return len(buf), nil
}

// Close closes the buffer and releases all resourses.
func (ab *AudioBuffer) Close() error {
    return ab.output.close()
}


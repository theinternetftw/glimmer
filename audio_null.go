// +build !windows
// +build !linux

package glimmer

import "time"

type audioOutput struct {
	ab *AudioBuffer

	fakeBufs []fakeBuf
}

type fakeBuf struct {
	bufConsumed    int
	lastAvailCheck time.Time
}

func (ao *audioOutput) init(ab *AudioBuffer) error {
	ao.ab = ab
	ao.fakeBufs = make([]fakeBuf, ao.ab.BlockCount)
	for i := range ao.fakeBufs {
		ao.fakeBufs[i].lastAvailCheck = time.Now()
	}
	return nil
}

func (ao *audioOutput) close() error {
	return nil
}

func (ao *audioOutput) noLongerBusy(blockIndex int) bool {
	buf := &ao.fakeBufs[blockIndex]
	lastCheck := buf.lastAvailCheck
	samplesPlayed := int(time.Now().Sub(lastCheck).Seconds() * float64(ao.ab.SamplesPerSecond))
	bytesWritten := samplesPlayed * int(ao.ab.ChannelCount) * int(ao.ab.BitsPerSample/8)
	buf.bufConsumed -= bytesWritten
	if buf.bufConsumed < 0 {
		buf.bufConsumed = 0
	}
	buf.lastAvailCheck = time.Now()
	return buf.bufConsumed == 0
}

func (ao *audioOutput) write(bytes []byte, i int) {
	ao.fakeBufs[i].bufConsumed += len(bytes)
	if ao.fakeBufs[i].bufConsumed > ao.ab.BufferSize() {
		ao.fakeBufs[i].bufConsumed = ao.ab.BufferSize()
	}
}

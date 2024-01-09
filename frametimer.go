package glimmer

import (
	"fmt"
	"time"
)

const frameTimeLogSize = 300

// FrameTimer allows for delaying at the end of
// a frame, and for checking frame stats
type FrameTimer struct {
	frameCount uint
    rtimeLog []time.Duration
    ftimeLog []time.Duration

	statsTimer int

	lastFlipTime time.Time
	lastRenderTime time.Time
}
// Makes a frame timer
func MakeFrameTimer() FrameTimer {
	return FrameTimer{
		lastFlipTime: time.Now(),
	}
}

// WaitForFrametime sleeps until the frametime matches the desired interval
func (f *FrameTimer) WaitForFrametime(timeGoalSecs float64) {

    f.MarkRenderComplete()

    frametimeGoal := time.Duration(float64(time.Second)*timeGoalSecs)

	rDiff := time.Now().Sub(f.lastFlipTime)

	// hack to get better accuracy, could do
	// a two stage wait with spin at the end,
	// but that really adds to the cycles
	fudge := 2 * time.Millisecond
	toWait := frametimeGoal - rDiff - fudge
	if toWait > time.Duration(0) {
		<-time.NewTimer(toWait).C
	}

    f.MarkFrameComplete()
}

// TickFrame updates frame render stats as if the render was just completed
func (f *FrameTimer) MarkRenderComplete() {
	rdiff := time.Now().Sub(f.lastFlipTime)
    f.rtimeLog = append(f.rtimeLog, rdiff)
    if len(f.rtimeLog) > frameTimeLogSize {
        f.rtimeLog = f.rtimeLog[1:]
    }
}

// TickFrame updates frame render stats as if the render was just completed
func (f *FrameTimer) MarkFrameComplete() {
	frameStart := f.lastFlipTime
	f.lastFlipTime = time.Now()

	fdiff := time.Now().Sub(frameStart)
    f.ftimeLog = append(f.ftimeLog, fdiff)
    if len(f.ftimeLog) > frameTimeLogSize {
        f.ftimeLog = f.ftimeLog[1:]
    }
	f.frameCount++
	f.statsTimer++
}

// PrintStatsEveryXFrames prints to stdout the presumed render time,
// as well as the full frame time after sleep. The argument
// determines how often (in frames) to issue a printout.
func (f *FrameTimer) PrintStatsEveryXFrames(numFrames int) {
	if f.statsTimer >= numFrames {

        var maxRTime time.Duration
        var maxFTime time.Duration
        var meanRTime time.Duration
        var meanFTime time.Duration
        for _, d := range f.rtimeLog {
            if d > maxRTime {
                maxRTime = d
            }
            meanRTime += d
        }
        for _, d := range f.ftimeLog {
            if d > maxFTime {
                maxFTime = d
            }
            meanFTime += d
        }
        meanRTime /= time.Duration(len(f.rtimeLog))
        meanFTime /= time.Duration(len(f.ftimeLog))
		fmt.Printf("meanRTime %.4f, meanFTime %.4f, maxRTime %.4f, maxFTime %.4f\n", meanRTime.Seconds(), meanFTime.Seconds(), maxRTime.Seconds(), maxFTime.Seconds())
		f.statsTimer = 0
	}
}

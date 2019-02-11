package glimmer

import (
	"fmt"
	"time"
)

// TODO: As a bonus, could check and see if
// just watching when the main glimmer thread
// publishes is enough to get wait for vsync
// support.
//
// However, right now for my purposes I need
// every-x-ms more than I need every vsync.

// FrameTimer allows for delaying at the end of
// a frame, and for checking frame stats
type FrameTimer struct {
	frameCount uint
	maxRDiff time.Duration
	maxFDiff time.Duration

	statsTimer int

	frametimeGoal time.Duration

	lastFlipTime time.Time
}
// Makes a frame timer for a certain goal framerate (in seconds)
func MakeFrameTimer(timeGoalSecs float64) FrameTimer {
	return FrameTimer{
		frametimeGoal: time.Duration(float64(time.Second)*timeGoalSecs),
		lastFlipTime: time.Now(),
	}
}

// WaitForFrametime sleeps until the frametime matches the desired interval
func (f *FrameTimer) WaitForFrametime() {

	rDiff := time.Now().Sub(f.lastFlipTime)

	// hack to get better accuracy, could do
	// a two stage wait with spin at the end,
	// but that really adds to the cycles
	fudge := 2 * time.Millisecond

	toWait := f.frametimeGoal - rDiff - fudge
	if toWait > time.Duration(0) {
		<-time.NewTimer(toWait).C
	}

	frameStart := f.lastFlipTime
	f.lastFlipTime = time.Now()

	fDiff := time.Now().Sub(frameStart)
	if rDiff > f.maxRDiff {
		f.maxRDiff = rDiff
	}
	if fDiff > f.maxFDiff {
		f.maxFDiff = fDiff
	}

	f.frameCount++
	f.statsTimer++
}

// PrintStatsEveryXFrames prints to stdout the presumed render time,
// as well as the full frame time after sleep. The argument
// determines how often (in frames) to issue a printout.
func (f *FrameTimer) PrintStatsEveryXFrames(numFrames int) {
	if f.statsTimer >= numFrames {

		maxRTime := f.maxRDiff.Seconds()
		maxFTime := f.maxFDiff.Seconds()
		fmt.Printf("maxRTime %.4f, maxFTime %.4f\n", maxRTime, maxFTime)

		f.statsTimer = 0
		f.maxRDiff = 0
		f.maxFDiff = 0
	}
}

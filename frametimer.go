package glimmer

import (
	"fmt"
	"time"
)

const timeLogSize = 300

type FrameTimer struct {
	frameCount    uint
	renderTimeLog []time.Duration
	frameTimeLog  []time.Duration

	statsTimer int

	lastFlipTime   time.Time
	lastRenderTime time.Time
}

func MakeFrameTimer() FrameTimer {
	return FrameTimer{
		lastFlipTime: time.Now(),
	}
}

func (f *FrameTimer) MarkRenderComplete() {
	rdiff := time.Now().Sub(f.lastFlipTime)
	f.renderTimeLog = append(f.renderTimeLog, rdiff)
	if len(f.renderTimeLog) > timeLogSize {
		f.renderTimeLog = f.renderTimeLog[1:]
	}
}

func (f *FrameTimer) MarkFrameComplete() {
	frameStart := f.lastFlipTime
	f.lastFlipTime = time.Now()

	fdiff := f.lastFlipTime.Sub(frameStart)
	f.frameTimeLog = append(f.frameTimeLog, fdiff)
	if len(f.frameTimeLog) > timeLogSize {
		f.frameTimeLog = f.frameTimeLog[1:]
	}
	f.frameCount++
	f.statsTimer++
}

type FrameStats struct {
	MeanRenderTime time.Duration
	MaxRenderTime  time.Duration
	MaxFrameTime   time.Duration
	MeanFrameTime  time.Duration
}

func (f *FrameTimer) ResetStats() {
	f.frameTimeLog = f.frameTimeLog[:0]
	f.renderTimeLog = f.renderTimeLog[:0]
}

func (f *FrameTimer) GetStats() FrameStats {
	stats := FrameStats{}
	for _, d := range f.renderTimeLog {
		if d > stats.MaxRenderTime {
			stats.MaxRenderTime = d
		}
		stats.MeanRenderTime += d
	}
	for _, d := range f.frameTimeLog {
		if d > stats.MaxFrameTime {
			stats.MaxFrameTime = d
		}
		stats.MeanFrameTime += d
	}
	stats.MeanRenderTime /= time.Duration(len(f.renderTimeLog))
	stats.MeanFrameTime /= time.Duration(len(f.frameTimeLog))
	return stats
}

// PrintStatsEveryXFrames prints to stdout the recorded mean/max render time,
// as well as the same for full frame time. The argument determines how often
// (in frames) to issue a printout. (Mean/max are for the past timeLogSize frames)
func (f *FrameTimer) PrintStatsEveryXFrames(numFrames int) {
	if f.statsTimer >= numFrames {
		stats := f.GetStats()

		fmt.Printf("meanRender %.4f, meanFrame %.4f, maxRender %.4f, maxFrame %.4f\n",
			stats.MeanRenderTime.Seconds(), stats.MeanFrameTime.Seconds(),
			stats.MaxRenderTime.Seconds(), stats.MaxFrameTime.Seconds())
		f.statsTimer = 0
	}
}

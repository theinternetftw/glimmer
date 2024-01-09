package glimmer

import (
    "fmt"

    "github.com/hajimehoshi/ebiten/v2/audio"
)

type audioOutput struct {
	context *audio.Context
	player *audio.Player
}

func (ao *audioOutput) init(ab *AudioBuffer) error {
    if ab.BitsPerSample != 16 || ab.ChannelCount != 2 {
        return fmt.Errorf("todo: use oto directly instead of ebiten, should be almost the same code")
    }
    ao.context = audio.NewContext(int(ab.SamplesPerSecond))
    player, err := ao.context.NewPlayer(ab)
    if err != nil {
        return err
    }
    player.SetBufferSize(ab.outputBufDuration)
    player.Play()

	ao.player = player
	return nil
}

func (ao *audioOutput) close() error {
	return ao.player.Close()
}

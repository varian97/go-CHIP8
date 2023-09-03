package emulator

import (
	"github.com/veandco/go-sdl2/sdl"
)

type audio struct {
	spec          *sdl.AudioSpec
	isAudioOpened bool
	audioDevice   sdl.AudioDeviceID
	wave          []byte
}

func MakeAudio(freq int32, sample uint16) *audio {
	spec := &sdl.AudioSpec{
		Freq:     freq,
		Format:   sdl.AUDIO_S16LSB,
		Channels: 2,
		Samples:  sample,
	}

	wave := make([]int16, sample)
	var i uint16 = 0
	for i = 0; i < sample; i++ {
		t := float64(i) / float64(spec.Freq)
		wave[i] = int16(32767 * 0.5 * (1.0 + 0.9*(2.0*3.14159265*float64(freq)*t)))
	}

	byteWave := make([]byte, len(wave)*2)
	for i, s := range wave {
		byteWave[i*2] = byte(s)
		byteWave[i*2+1] = byte(s >> 8)
	}

	audio := &audio{
		spec:          spec,
		isAudioOpened: false,
		wave:          byteWave,
	}

	return audio
}

func (audio *audio) OpenAudio() error {
	if audio.isAudioOpened {
		return nil
	}
	audioDevice, err := sdl.OpenAudioDevice("", false, audio.spec, nil, 0)
	if err != nil {
		return err
	}
	audio.audioDevice = audioDevice
	audio.isAudioOpened = true
	return nil
}

func (audio *audio) PlayAudio() {
	if audio.isAudioOpened {
		sdl.PauseAudioDevice(audio.audioDevice, false)
		sdl.QueueAudio(audio.audioDevice, audio.wave)
	}
}

func (audio *audio) StopAudio() {
	if audio.isAudioOpened {
		sdl.PauseAudioDevice(audio.audioDevice, true)
	}
}

func (audio *audio) Cleanup() {
	sdl.CloseAudioDevice(audio.audioDevice)
}

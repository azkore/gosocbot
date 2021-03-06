package main

import (
	"bytes"
	"github.com/zhulik/margelet"
	"gopkg.in/telegram-bot-api.v4"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

type LissajousHandler struct {
}

func (h LissajousHandler) HandleCommand(message margelet.Message) error {
	var buffer bytes.Buffer
	message.SendUploadPhotoAction()

	lissajous(&buffer)

	msg := tgbotapi.NewVideoUpload(message.Message().Chat.ID,
		tgbotapi.FileBytes{Name: "lissajous.gif", Bytes: buffer.Bytes()})
	msg.ReplyToMessageID = message.Message().MessageID

	message.Bot().Send(msg)

	return nil
}

func (h LissajousHandler) HelpMessage() string {
	return "Send Lissajous figure"
}

func lissajous(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())

	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 128   // number of animation frames
		delay   = 8     // delay between frames in 10ms units
		maxcolors = 16  // max number of colors in the palette
	)
	ncolors := rand.Intn(maxcolors-2)+2
	palette := make([]color.Color, ncolors)
	b := make([]byte, 3)
	for i := 0; i < ncolors; i++ {
		palette[i] = color.RGBA{R: b[0], G: b[1], B: b[2], A: 1}
		rand.Read(b)
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		colorIndex := 1;
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex++;
			if colorIndex == ncolors {
				colorIndex = 1
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

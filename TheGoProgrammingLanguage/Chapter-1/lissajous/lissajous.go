// Lissajous generates random lissajous figures and prints them
// to stdout.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	//color.White,
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xEE, 0x7F, 0xEE, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete x oscillator revolutions
		res     = 0.001 // Angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames in gif
		delay   = 8     // delay between frames in units of 10ms
	)

	freq := rand.Float64() * 3.0 // Relative frecuency of Y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			randColor := rand.Int()%(len(palette)-1) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randColor))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	lissajous(os.Stdout)
}

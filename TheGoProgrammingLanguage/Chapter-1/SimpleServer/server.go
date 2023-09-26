// Runs a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			cycles  = 5     // Number of complete x oscillator revolutions
			res     = 0.001 // Angular resolution
			size    = 100   // image canvas covers [-size..+size]
			nframes = 64    // number of animation frames in gif
			delay   = 8     // delay between frames in units of 10ms
		)

		params := r.URL.Query()
		if par, ok := params["cycles"]; ok {
			arg, err := strconv.Atoi(par[0])
			if err == nil {
				cycles = arg
			}
		}

		if par, ok := params["size"]; ok {
			arg, err := strconv.Atoi(par[0])
			if err == nil {
				size = arg
			}
		}

		if par, ok := params["delay"]; ok {
			arg, err := strconv.Atoi(par[0])
			if err == nil {
				delay = arg
			}
		}

		if par, ok := params["nframes"]; ok {
			arg, err := strconv.Atoi(par[0])
			if err == nil {
				nframes = arg
			}
		}

		lissajous(w, cycles, size, nframes, delay, res)
	})
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "Forma[%q] = %q\n", key, value)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Hits so far: %d\n", count)
	mu.Unlock()
}

// Lissajous code below
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

func lissajous(out io.Writer, cycles, size, nframes, delay int, res float64) {

	freq := rand.Float64() * 4.5 // Relative frecuency of Y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			randColor := rand.Int()%(len(palette)-1) + 1
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(randColor))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

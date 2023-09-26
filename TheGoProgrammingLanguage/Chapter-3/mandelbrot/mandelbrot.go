// Mandelbrot emits a png image of the mandelbrot set.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func BlackAndWhiteMandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// Shading depends on number of iterations it took to "escape"
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func BadlyColoredMandelbrot(z complex128) color.RGBA {
	const iterations = 1000
	const contrast = 200
	const RGB uint32 = 0xFFFFFF

	var v complex128
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// Shading depends on number of iterations it took to "escape"
			offset := n * contrast
			k := RGB - uint32(offset)
			r, b, g := uint8(k&0xFF), uint8((k>>0x8)&0xFF), uint8((k>>0x10)&0xFF)
			return color.RGBA{r, g, b, 0xFF}
		}
	}
	return color.RGBA{0, 0, 0, 0xFF}
}

// Used as a Helper Function to calculate the color at different places
// in each pixel and then take the average.
func SupersamplingComputeColor(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n * contrast
		}
	}
	return 0xFF
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
		samplingFactor         = 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y0 := (float64(py)+0.25)/height*(ymax-ymin) + ymin
		y1 := (float64(py)-0.25)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := (float64(px)+0.25)/width*(xmax-xmin) + xmin
			x1 := (float64(px)-0.25)/width*(xmax-xmin) + xmin

			var cmplxPoints = []complex128{complex(x0, y0), complex(x0, y1), complex(x1, y0), complex(x1, y1)}
			var avg, counter uint32
			for _, zx := range cmplxPoints {
				avg += uint32(SupersamplingComputeColor(zx))
				counter++
			}
			avg /= counter

			img.Set(px, py, color.Gray{255 - uint8(avg)})
		}
	}

	png.Encode(os.Stdout, img) // ignoring possible errors
}

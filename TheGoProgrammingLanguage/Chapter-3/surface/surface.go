// Computes an SVG rendering of a 3D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 900, 600            // canvas size in pixels
	cells         = 300                 // number of grid cells
	xyrange       = 30.0                // axis range (-xyrange, +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axis in radians (30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // measures the distance between x and y
	return (math.Sin(r) * math.Cos(r)) / r
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute z from mathmatical function f we're displaying
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale

	return sx, sy
}

func polygonHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color := colorBasedOnZ(colorZ(i, j))

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func colorZ(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	return f(x, y)
}

func colorBasedOnZ(z float64) string {
	// Interpolate between blue and red based on Z-values.
	// Use the HSL color space.
	// Lesser values (bluer) correspond to 240 (blue in HSL), and greater values (redder) correspond to 0 (red in HSL).
	h := (1 - (z+1)/2) * 240
	return "hsl(" + strconv.FormatFloat(h, 'f', 2, 64) + ", 100%, 50%)"
}

func main() {
	http.HandleFunc("/", polygonHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

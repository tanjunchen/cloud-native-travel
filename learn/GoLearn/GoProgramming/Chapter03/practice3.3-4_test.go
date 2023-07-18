package Chapter03

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"testing"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func imageMain(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	min, max := getMinMax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, r, g, b := corner(i+1, j, min, max)
			bx, by, r, g, b := corner(i, j, min, max)
			cx, cy, r, g, b := corner(i, j+1, min, max)
			dx, dy, r, g, b := corner(i+1, j+1, min, max)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%x%x%x'/>n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func getColor(min, max, current float64) (int, int, int) {
	step := (max - min) / 255
	v := int((current - min) / step)
	r := v
	g := 0
	b := 255 - v
	return r, g, b
}

func getMinMax() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := f(x, y)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return min, max
}

func corner(i, j int, min, max float64) (float64, float64, int, int, int) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	r, g, b := getColor(min, max, z)
	return sx, sy, r, g, b
}

func f(x float64, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func test334() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		imageMain(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8877", nil))
	return
}

func Test3_3_4(t *testing.T) {
	test334()
}

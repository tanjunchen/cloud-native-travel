package Chapter03

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
	"testing"
)

func test039main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-png")
		xv := r.URL.Query().Get("x")
		x, err := strconv.Atoi(xv)
		if err != nil {
			x = 2
		}
		yv := r.URL.Query().Get("y")
		y, err := strconv.Atoi(yv)
		if err != nil {
			y = 2
		}
		generate(x, y, w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8877", nil))
	return
}

func generate(_x, _y int, w io.Writer) {
	var xmin, ymin, xmax, ymax float64 = float64(-1 * _x), float64(-1 * _x), float64(_y), float64(_y)
	const (
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			v := 255 - contrast*n
			return color.YCbCr{v, 255 - v, 255}
		}
	}
	return color.Black
}

func Test039(t *testing.T) {
	test039main()
}

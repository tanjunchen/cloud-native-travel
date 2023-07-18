package Chapter01

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

var palette2 []color.Color

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 256; i++ {
		palette2 = append(palette2, color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 0xFF,
		})
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[4] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous2(os.Stdout)

}
func lissajous2(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 //	角度分辨率
		size    = 100   //	图像画布包含 [-size..+size]
		nframes = 64    //	动画中的帧数
		delay   = 8     //	以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// 然后点击访问 localhost:8000  运行这个函数  传递参数 web
func Test002(t *testing.T) {
	main()
}

package Chapter02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

type Celsius float64
type Fahrenheit float64
type Kelvins float64

func main021() {
	fmt.Printf("Brrrrr! %v\n", AbsoluteC)
	fmt.Println(CToF(BoilingC))
	fmt.Println(CToK(FreezingC))
}

const (
	AbsoluteC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (k Kelvins) String() string    { return fmt.Sprintf("%g `K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func CToK(c Celsius) Kelvins { return Kelvins(c - AbsoluteC) }

func Test021(t *testing.T) {
	main021()
}

type Inch float64
type Meter float64
type Ib float64
type Kilogram float64

func (i Inch) String() string     { return fmt.Sprintf("%g in", i) }
func (m Meter) String() string    { return fmt.Sprintf("%g m", m) }
func (i Ib) String() string       { return fmt.Sprintf("%g ib", i) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func IToM(i Inch) Meter { return Meter(i * 0.0254) }

func MToI(m Meter) Inch { return Inch(m / 0.0254) }

func IToK(i Ib) Kilogram { return Kilogram(i * 0.4535) }

func KToI(k Kilogram) Ib { return Ib(k / 0.4535) }

func main022() {
	for _, str := range os.Args[4:] {
		t, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
		}
		c := Celsius(t)
		f := Fahrenheit(t)

		fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))

		in := Inch(t)
		m := Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", in, IToM(in), m, MToI(m))

		ib := Ib(t)
		k := Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", ib, IToK(ib), k, KToI(k))
	}
}

func Test022(t *testing.T) {
	main022()
}

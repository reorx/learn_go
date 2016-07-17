package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strings"
)

type Axes struct {
	x, y int
}

func (a Axes) Linear() int {
	fmt.Printf("address of a in Linear %p \n", &a)
	return a.x + a.y
}

func (a *Axes) Curve1() int {
	fmt.Printf("address of a in Axes %p \n", a)
	return a.x * a.y
}

type MyInt int

func (i MyInt) Negative() MyInt {
	return MyInt(-i)
}

func main() {
	a := Axes{3, 2}
	fmt.Printf("address of a %p \n", &a)
	fmt.Println(a.Linear(), a.Curve1())

	i := MyInt(1)
	fmt.Println(i.Negative(), i.Negative().Negative())

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	var m image.Image
	m = Image{200, 100}
	fmt.Printf("\n%v %T \n", m, m)
}

// Exercise : rot13Reader
const alpstart byte = 'A'
const alpend byte = 'z'

func rot13b(a byte) byte {
	b := a + 13
	if b > alpend {
		b = b - alpend - 1 + alpstart
	}
	return b
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	i, e := r.r.Read(b)
	if i > 0 {
		for j := 0; j < i; j++ {
			b[j] = rot13b(b[j])
		}
	}
	return i, e
}

// Exercise: Images
type Image struct {
	W int
	H int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

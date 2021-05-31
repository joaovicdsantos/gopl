// Lissajous gera animações GIF de figuras de Lissajous aleatórias
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x2E, 0xCC, 0x71, 0xFF}}

const (
	backIndex  = 0
	frontIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Número de revoluções completas do oscilador x
		res     = 0.001 // Resolução angular
		size    = 100   // Canvas da imagem cobre de  [-size..+size]
		nframes = 64    // Números de quadros da animação
		delay   = 8     // Tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // Frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), frontIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Ignorando erros de codificação
}

// Lissajous gera animações GIF de figuras de Lissajous aleatórias
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x2E, 0xCC, 0x71, 0xFF}}

const (
	backIndex  = 0
	frontIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conf := make(map[string]float64)

		// Default Values
		conf["cycles"] = 5
		conf["res"] = 0.001
		conf["size"] = 100
		conf["nframes"] = 64
		conf["delay"] = 8

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for k, v := range r.Form {
			value, err := strconv.ParseFloat(strings.Join(v, ""), 64)
			if err != nil {
				log.Print(err)
			}
			conf[k] = value
		}
		lissajous(w, conf)
	})
	http.ListenAndServe("localhost:8000", nil)
}

func lissajous(out io.Writer, conf map[string]float64) {
	// Valores padrões
	var (
		cycles  = conf["cycles"]       // Número de revoluções completas do oscilador x
		res     = conf["res"]          // Resolução angular
		size    = conf["size"]         // Canvas da imagem cobre de  [-size..+size]
		nframes = int(conf["nframes"]) // Números de quadros da animação
		delay   = int(conf["delay"])   // Tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // Frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5), frontIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Ignorando erros de codificação
}

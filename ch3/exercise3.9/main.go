// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getImageArg(r *http.Request, key string, defaultValue float64) float64 {
	valueStr := r.URL.Query().Get(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := getImageArg(r, "x", 0.0)
	y := getImageArg(r, "y", 0.0)
	zoom := getImageArg(r, "zoom", 1.0)
	var xmin, ymin, xmax, ymax = -2 * zoom, -2 * zoom, +2 * zoom, +2 * zoom
	var width, height = 1024 * zoom, 1024 * zoom
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := y + float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := x + float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex64(complex(x/zoom, y/zoom))
			img.Set(px, py, mandelbrotByComplex64(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrotByComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if real(v)*real(v)-imag(v)*imag(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

type complexBigFloat struct {
	real big.Float
	imag big.Float
}

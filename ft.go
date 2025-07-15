package main

import (
	"math"
	"math/cmplx"
)

func Fft(y []complex128, length int) []complex128 {

	if length <= 1 {
		return y
	}

	y_even, y_odd := Split(y, length)

	w_even := Fft(y_even, length/2)
	w_odd := Fft(y_odd, length/2)
	
	w := make([]complex128, length)

	for j:=0; j < length / 2; j++ {
		angle := math.Pi * float64(j) / float64(length)
		arg := -2.0i * complex(angle, 0)
		z := complex128(cmplx.Exp(arg)) * w_odd[j]
		w[j] = w_even[j] + z
		w[j + length/2] = w_even[j] - z
	}

	return w
}

func Pad(list []complex128, to int) []complex128 {

	if (to < len(list)) {
		panic("Can't shrink array in padding.")
	}

	padded := make([]complex128, to)

	for i:=0; i < len(list); i++ {
		padded[i] = list[i]
	}

	return padded
}

func Split(list []complex128, size int) (evens []complex128, odds []complex128) {
	evens = make([]complex128, size/2)
	odds = make([]complex128, size/2)

	for i:=0; i < size/2; i++ {
		evens[i] = list[2*i]
		odds[i] = list[2*i+1]
	}

	return
}

func Flatten(signal []complex128) (re []float64, im []float64) {

	re = make([]float64, 0, len(signal))
	for _, e := range signal {
		re = append(re, real(e))
	}

	im = make([]float64, 0, len(signal))
	for _, e := range signal {
		im = append(im, imag(e))
	}

	return
}

func Zip(re []float64, im []float64) []complex128 {
	
	zipped := make([]complex128, 0, len(re))
	for j, _ := range re {
		zipped = append(zipped, complex(re[j],im[j]))
	}

	return zipped
}



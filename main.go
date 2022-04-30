package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"golang.org/x/exp/constraints"
)

func mapRange[T constraints.Float](x0, x1, y0, y1, n T) T {
	return y0 + (n-x0)*(y1-y0)/(x1-x0)
}

func convex(im image.Image) image.Image {
	out := image.NewGray(im.Bounds())
	bounds := im.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			x_ := mapRange(0, float64(bounds.Max.X), 0.589924771587608, -0.589924771587608, float64(x))
			y_ := mapRange(0, float64(bounds.Max.Y), 0.589924771587608, -0.589924771587608, float64(y))

			theta_x := math.Asin(x_)
			theta_y := math.Asin(y_)

			x_out := x_ + (1-math.Cos(theta_x))*math.Tan(2*theta_x)
			y_out := y_ + (1-math.Cos(theta_y))*math.Tan(2*theta_y)

			x_out = mapRange(1, -1, 0, float64(bounds.Max.X), float64(x_out))
			y_out = mapRange(1, -1, 0, float64(bounds.Max.Y), float64(y_out))

			out.Set(int(x_out), int(y_out), im.At(x, y))
		}
	}

	return out
}

func main() {
	imFile, err := os.Open("input.png")
	if err != nil {
		log.Fatal(err)
	}
	imFileOut, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imFile.Close()
	defer imFileOut.Close()

	im, err := png.Decode(imFile)
	if err != nil {
		log.Fatal(err)
	}

	// applies the convex effect.
	imOut := convex(im)

	png.Encode(imFileOut, imOut)
}

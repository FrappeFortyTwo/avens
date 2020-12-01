package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func logoGen() {

	img := image.NewRGBA(image.Rect(0, 0, 64, 64)) // x1,y1,x2,y2

	// backfill alpha
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0}}, image.ZP, draw.Src)

	// draw rectangles
	draw.Draw(img, image.Rect(12, 0, 64, 51), &image.Uniform{color.RGBA{57, 62, 70, 255}}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(0, 12, 51, 64), &image.Uniform{color.RGBA{255, 211, 105, 255}}, image.ZP, draw.Src)

	// save file
	logofile, err := os.Create(ficon)
	checkErr("unable to create missing file : logo.ico", err)
	png.Encode(logofile, img)

}

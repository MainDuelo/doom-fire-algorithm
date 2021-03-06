package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 100
	screenHeight = 50
)

var (
	screenSize  = screenWidth * screenHeight
	pixels      = make([]byte, screenSize*4)
	firePixels  = make([]byte, screenSize)
	firePalette = []color.RGBA{
		{R: 7, G: 7, B: 7},       //  0
		{R: 31, G: 7, B: 7},      //  1
		{R: 47, G: 15, B: 7},     //  2
		{R: 71, G: 15, B: 7},     //  3
		{R: 87, G: 23, B: 7},     //  4
		{R: 103, G: 31, B: 7},    //  5
		{R: 119, G: 31, B: 7},    //  6
		{R: 143, G: 39, B: 7},    //  7
		{R: 159, G: 47, B: 7},    //  8
		{R: 175, G: 63, B: 7},    //  9
		{R: 191, G: 71, B: 7},    // 10
		{R: 199, G: 71, B: 7},    // 11
		{R: 223, G: 79, B: 7},    // 12
		{R: 223, G: 87, B: 7},    // 13
		{R: 223, G: 87, B: 7},    // 14
		{R: 215, G: 95, B: 7},    // 15
		{R: 215, G: 95, B: 7},    // 16
		{R: 215, G: 103, B: 15},  // 17
		{R: 207, G: 111, B: 15},  // 18
		{R: 207, G: 119, B: 15},  // 19
		{R: 207, G: 127, B: 15},  // 20
		{R: 207, G: 135, B: 23},  // 21
		{R: 199, G: 135, B: 23},  // 22
		{R: 199, G: 143, B: 23},  // 23
		{R: 199, G: 151, B: 31},  // 24
		{R: 191, G: 159, B: 31},  // 25
		{R: 191, G: 159, B: 31},  // 26
		{R: 191, G: 167, B: 39},  // 27
		{R: 191, G: 167, B: 39},  // 28
		{R: 191, G: 175, B: 47},  // 29
		{R: 183, G: 175, B: 47},  // 30
		{R: 183, G: 183, B: 47},  // 31
		{R: 183, G: 183, B: 55},  // 32
		{R: 207, G: 207, B: 111}, // 33
		{R: 223, G: 223, B: 159}, // 34
		{R: 239, G: 239, B: 199}, // 35
		{R: 255, G: 255, B: 255}, // 36
	}
)

func init() {
	for i := screenSize - screenWidth; i < screenSize; i++ {
		firePixels[i] = 36
	}
}

func updateFirePixels() {
	for i := 0; i < screenWidth; i++ {
		for j := 0; j < screenHeight; j++ {
			idx := i + (screenWidth * j)
			updateFireIntensityPerPixel(idx)
		}
	}
}

func updateFireIntensityPerPixel(currentPixelIndex int) {
	below := currentPixelIndex + screenWidth
	if below >= screenSize {
		return
	}

	d := rand.Intn(3)
	newI := int(firePixels[below]) - d
	if newI < 0 {
		newI = 0
	}

	if currentPixelIndex-d < 0 {
		return
	}
	firePixels[currentPixelIndex-d] = byte(newI)
}

func renderFire() {
	for i, v := range firePixels {
		p := firePalette[v]
		pixels[i*4] = p.R
		pixels[i*4+1] = p.G
		pixels[i*4+2] = p.B
		pixels[i*4+3] = 0xff
	}
}

func update(screen *ebiten.Image) error {
	updateFirePixels()
	renderFire()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	_ = screen.ReplacePixels(pixels)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := ebiten.Run(update, screenWidth, screenHeight, 6, "Doom Fire (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}

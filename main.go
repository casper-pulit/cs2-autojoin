package main

import (
	"fmt"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("Hello World!")

	n_display := screenshot.NumActiveDisplays()

	fmt.Println(n_display)

	bounds := screenshot.GetDisplayBounds(1)

	for {
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			panic(err)
		}

		mid_x := img.Rect.Bounds().Max.X / 2
		mid_y := img.Rect.Bounds().Max.Y/2 - 120

		fmt.Println("Mid X: ", mid_x, "Mid Y: ", mid_y)

		rgb_val := img.RGBAAt(mid_x, mid_y)

		fmt.Println(rgb_val)

		time.Sleep(1 * time.Second)
	}

	/* 	for i := 0; i < 1440; i++ {

		rgb_val := img.RGBAAt(mid_x, i)

		if rgb_val.G > 180 {
			fmt.Println("Pretty Green!")
			fmt.Println(rgb_val)
			fmt.Println(mid_x, i)
		}
	} */

}

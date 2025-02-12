package main

import (
	"fmt"
	sdeck "github.com/AKovalevich/streamdeck"
	"github.com/AKovalevich/streamdeck/label"
	"image"
	"image/color"
	"log"
	"strconv"
)

// This example will instanciate 15 labels on the streamdeck. Each Label
// is setup as a counter which will increment every 50ms. If a button is
// pressed it will be colored blue until it is released.

func main() {
	sd, err := sdeck.NewStreamDeck(nil)
	if err != nil {
		log.Panic(err)
	}
	defer sd.ClearAllBtns()

	labels := make(map[int]*label.Label)

	for i := 0; i < 15; i++ {
		label, err := label.NewLabel(sd, i, label.Text(strconv.Itoa(i)))
		if err != nil {
			fmt.Println(err)
		}
		label.Draw()
		labels[i] = label
	}

	handleBtnEvents := func(btnIndex int, state sdeck.BtnState) {
		fmt.Printf("Button: %d, %s\n", btnIndex, state)
		if state == sdeck.BtnPressed {
			col := color.RGBA{0, 0, 153, 0}
			labels[btnIndex].SetBgColor(image.NewUniform(col))
		} else { // must be BtnReleased
			col := color.RGBA{0, 0, 0, 255}
			labels[btnIndex].SetBgColor(image.NewUniform(col))
		}
	}

	sd.SetBtnEventCb(handleBtnEvents)

	stop := make(chan bool)
	sd.Serve(stop)
}

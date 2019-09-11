package main

import (
	"bytes"
	"fmt"
	sdeck "github.com/AKovalevich/streamdeck"
	"github.com/gobuffalo/packr/v2"
	"image"
	"log"
)

// This example loads icons and places them on buttons in the first row
// of the Stream Deck. The lightbulb icon on button 0 can be toggled.

func main() {
	sd, err := sdeck.NewStreamDeck(nil)
	if err != nil {
		log.Panic(err)
	}
	defer sd.ClearAllBtns()

	imgBox := packr.New("icons-images", "../assets/images")

	_user, err := imgBox.Find("user.png")
	if err != nil {
		log.Panic(err)
	}
	user, _, err := image.Decode(bytes.NewBuffer(_user))
	if err != nil {
		log.Panic(err)
	}

	_tux, err := imgBox.Find("tux.png")
	if err != nil {
		log.Panic(err)
	}
	tux, _, err := image.Decode(bytes.NewBuffer(_tux))
	if err != nil {
		log.Panic(err)
	}

	_warning, err := imgBox.Find("warning.png")
	if err != nil {
		log.Panic(err)
	}
	warning, _, err := image.Decode(bytes.NewBuffer(_warning))
	if err != nil {
		log.Panic(err)
	}

	_doctor, err := imgBox.Find("doctor.png")
	if err != nil {
		log.Panic(err)
	}
	doctor, _, err := image.Decode(bytes.NewBuffer(_doctor))
	if err != nil {
		log.Panic(err)
	}

	_lightbulbOn, err := imgBox.Find("lightbulb_on.png")
	if err != nil {
		log.Panic(err)
	}
	lightbulbOn, _, err := image.Decode(bytes.NewBuffer(_lightbulbOn))
	if err != nil {
		log.Panic(err)
	}

	_lightbulbOff, err := imgBox.Find("lightbulb_off.png")
	if err != nil {
		log.Panic(err)
	}
	lightbulbOff, _, err := image.Decode(bytes.NewBuffer(_lightbulbOff))
	if err != nil {
		log.Panic(err)
	}

	if err := sd.FillImage(4, warning); err != nil {
		log.Panic(err)
	}
	if err := sd.FillImage(3, doctor); err != nil {
		log.Panic(err)
	}
	if err := sd.FillImage(2, tux); err != nil {
		log.Panic(err)
	}
	if err := sd.FillImage(1, user); err != nil {
		log.Panic(err)
	}
	if err := sd.FillImage(0, lightbulbOff); err != nil {
		log.Panic(err)
	}

	lightbulb := false

	onPressedCb := func(btnIndex int, state sdeck.BtnState) {
		fmt.Printf("Button: %d, %s\n", btnIndex, state)
		if btnIndex == 0 && state == sdeck.BtnPressed {
			if lightbulb {
				if err := sd.FillImage(0, lightbulbOff); err != nil {
					log.Panic(err)
				}
				lightbulb = false
			} else {
				if err := sd.FillImage(0, lightbulbOn); err != nil {
					log.Panic(err)
				}
				lightbulb = true
			}
		}
	}

	sd.SetBtnEventCb(onPressedCb)

	stop := make(chan bool)
	sd.Serve(stop)
}

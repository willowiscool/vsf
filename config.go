package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LIST_LENGTH int `json:"list_length"` //the length of the list
	BLOCK_WIDTH int `json:"block_width"` //the width of each block
	BLOCK_HEIGHT_MULT float64 `json:"block_height_mult"` //the amout each block's height is multiplied by when displayed
	SHOWER string `json:"shower"` //the method to use when displaying the sort
	SLEEP float64 `json:"sleep"` //the amount of sleep time between showings
	BG [4]uint8 `json:"bg"` //the background color, in RGBA, with each number being between 0 and 0xff
	FG [4]uint8 `json:"fg"` //the foreground color, as above
	CHANGED [4]uint8 `json:"changed"` //the color of changed blocks, as above
	RAINBOW bool `json:"rainbow"` //whether to use rainbow or not
	VSYNC bool `json:"vsync"` //whether to use VSync or not
	FPSFILTER int `json:"fpsfilter"` //the number of recent frames to average for FPS. Default of 30 is fine
}
func parse(filename string) (*Config, error) {
	var file []byte
	var err error
	if filename != "" {
		file, err = ioutil.ReadFile(filename)
		if err != nil {
			return &Config{}, err
		}
	}
	config := &Config{
		LIST_LENGTH: 500,
		BLOCK_WIDTH: 2,
		BLOCK_HEIGHT_MULT: 1,
		SHOWER: "rect",
		SLEEP: 10,
		BG: [4]uint8{0, 0, 0, 0xff},
		FG: [4]uint8{0xff, 0xff, 0xff, 0xff},
		CHANGED: [4]uint8{0xff, 0, 0, 0xff},
		RAINBOW: false,
		VSYNC: false,
		FPSFILTER: 30,
	}
	if filename != "" {
		err := json.Unmarshal(file, config)
		if err != nil {
			return &Config{}, err
		}
	}
	return config, nil
}

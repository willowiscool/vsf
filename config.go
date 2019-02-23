package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LIST_LENGTH int `json:"list_length"` //the length of the list
	BLOCK_WIDTH int `json:"block_width"` //the width of each block
	BLOCK_HEIGHT_MULT int `json:"block_height_mult"` //the amout each block's height is multiplied by when displayed
	SLEEP int `json:"sleep"` //the amount of sleep time between showings
	BG [4]uint8 `json:"bg"` //the background color, in RGBA, with each number being between 0 and 0xff
	FG [4]uint8 `json:"fg"` //the foreground color, as above
	CHANGED [4]uint8 `json:"changed"` //the color of changed blocks, as above
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
		SLEEP: 10,
		BG: [4]uint8{0, 0, 0, 0xff},
		FG: [4]uint8{0xff, 0xff, 0xff, 0xff},
		CHANGED: [4]uint8{0xff, 0, 0, 0xff},
	}
	if filename != "" {
		err := json.Unmarshal(file, config)
		if err != nil {
			return &Config{}, err
		}
	}
	return config, nil
}
# VSF

A program to visualize sorting algorithms written in Lua

## Installation

First, make sure you have all of the dependencies listed [here](https://github.com/go-gl/glfw#installation). Next, just run `go get -u github.com/vityavv/vsf` to install it.

## Usage

### Basic Command Line Usage

Run vsf with the vsf command: `vsf <lua file> [config file]`

There are a few example lua files in the `sorts` directory and an example config file called `defaults.json`. The config file is optional.

### Config File

The config file is a JSON file. You can see an example of this in the `defaults.json` file. Here are the keys and their meanings:

| Key | Meaning | Default |
| --- | --- | --- |
| `list_length` | The length of the list to be sorted | `500` |
| `block_width` | The width of each displayed block | `2` |
| `block_height_mult` | The amount each block's height is multiplied by when displayed | `1` |
| `shower` | The shower to use. See below | `rect` |
| `sleep` | How long to sleep before showing something, in milliseconds | `10` |
| `bg` | An array of four values from 0-255, representing the RGBA color of the background | `[0, 0, 0, 255]` |
| `fg` | The foreground color (color of the blocks), in the same manner | `[255, 255, 255, 255]` |
| `changed` | The color that changed blocks are shown | `[255, 0, 0, 255]` |

### Writing Sorts

Sorts are written in Lua. You must have a function called `sort`, which takes an array (table) as its argument. That function will be called with the random list at the start. You must then call the function `show` whenever you make a change in that array. The `show` function takes an array of the same length as the original. Check the `sorts` folder for examples.

### Showers

There are currently three different showers: `"rect"`, `"point"`, and `"circle"`. For the point and rect showers, refer to the description above for what `block_width` and `block_height_mult` do. The `"circle"` shower uses `block_height_mult` and makes the width of each piece of the circle equal to that times its value, but ignores `block_width`.

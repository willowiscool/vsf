package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"github.com/yuin/gopher-lua"

	"math/rand"
	"image/color"
	"time"
	"fmt"
	"os"
	"sync"
)
const (
	LIST_LENGTH = 500 //the length of the list to be sorted
	BLOCK_WIDTH = 2 //the width of each block
	BLOCK_HEIGHT_MULT = 1 //the amount that the height of the block is multiplied by its position
	WIDTH = LIST_LENGTH * BLOCK_WIDTH //the width of the window
	HEIGHT = LIST_LENGTH * BLOCK_HEIGHT_MULT //the height of the window
	SLEEP = 10 //how many milliseconds to sleep between showings
)
var (
	list []int
	changed [LIST_LENGTH]bool
	stop = make(chan byte, 1)
	running = true
	wg = sync.WaitGroup{}
	FILENAME string
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	
	//background
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(0, 0))
	imd.Color = pixel.RGB(0, 1, 1)
	imd.Push(pixel.V(WIDTH, HEIGHT))
	imd.Rectangle(0)

	//filename
	atlas := text.NewAtlas(
		basicfont.Face7x13,
		[]rune(FILENAME),
		[]rune("VSF: "),
		text.ASCII)
	txt := text.New(pixel.V(0, HEIGHT - 26), atlas) //26 = 2 * height
	fmt.Fprintf(txt, "VSF: %s", FILENAME)

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeySpace) {
			if running {
				wg.Add(1)
			} else {
				wg.Done()
			}
			running = !running
		}
		win.Clear(color.RGBA{0, 0, 0, 0xff})
		//imd.Draw(win)
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		for i, val := range list {
			rect := imdraw.New(nil)
			if changed[i] {
				rect.Color = pixel.RGB(1, 0, 0)
			} else {
				rect.Color = pixel.RGB(1, 1, 1)
			}
			rect.Push(pixel.V(
				float64(i * BLOCK_WIDTH),
				float64(val * BLOCK_HEIGHT_MULT)))
			rect.Push(pixel.V(float64((i+1) * BLOCK_WIDTH), 0))
			rect.Rectangle(0)
			rect.Draw(win)
		}
		win.Update()
	}
	stop<-1
}

func show(L *lua.LState) int {
	wg.Wait()
	time.Sleep(SLEEP * time.Millisecond)
	newTable := L.ToTable(1)
	if newTable.Len() != LIST_LENGTH {
		//TODO: better error handling
		panic("List of improper length given")
	}
	newList := make([]int, LIST_LENGTH)
	exists := make(map[int]bool, LIST_LENGTH)
	newTable.ForEach(func(a, b lua.LValue) {
		if b.(lua.LNumber) <= 0 || b.(lua.LNumber) > LIST_LENGTH {
			panic("Invalid value found in list")
		}
		if exists[int(b.(lua.LNumber))] {
			panic("Duplicate value found in list")
		} else {
			exists[int(b.(lua.LNumber))] = true
		}
		newList[int(a.(lua.LNumber))-1] = int(b.(lua.LNumber))
		if int(b.(lua.LNumber)) != list[int(a.(lua.LNumber))-1] {
			changed[int(a.(lua.LNumber))-1] = true
		} else {
			changed[int(a.(lua.LNumber))-1] = false
		}
	})
	list = newList
	return 0
}

func main() {
	if len(os.Args) == 2 {
		FILENAME = os.Args[1]
	} else {
		fmt.Println("Usage: vsf <lua file>\nlua file must contain a function called sort that sorts an array and use the show function to display that array on the screen\npress space to pause")
		return
	}

	rand.Seed(time.Now().UnixNano())
	list = rand.Perm(LIST_LENGTH)
	for i, _ := range list {
		list[i]++
	}
	fmt.Println(list)

	go func() {
		L := lua.NewState()
		defer L.Close()
		L.SetGlobal("show", L.NewFunction(show))
		if err := L.DoFile(FILENAME); err != nil {
			panic(err)
		}
		tableList := lua.LTable{}
		for i, val := range list {
			tableList.Insert(i+1, lua.LNumber(val))
		}
		err := L.CallByParam(lua.P{
			Fn: L.GetGlobal("sort"),
			NRet: 1,
			Protect: true,
		}, &tableList)
		if err != nil {
			panic(err)
		}
	}()
	pixelgl.Run(run)
	<-stop
}


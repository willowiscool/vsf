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
var (
	CONFIG *Config
	list []int
	changed []bool
	stop = make(chan byte, 1)
	running = true
	wg = sync.WaitGroup{}
	FILENAME string
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0,
			float64(CONFIG.LIST_LENGTH * CONFIG.BLOCK_WIDTH),
			float64(CONFIG.LIST_LENGTH * CONFIG.BLOCK_HEIGHT_MULT)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	
	//filename
	atlas := text.NewAtlas(
		basicfont.Face7x13,
		[]rune(FILENAME),
		[]rune("VSF: "),
		text.ASCII)
	txt := text.New(pixel.V(0,
		float64((CONFIG.LIST_LENGTH * CONFIG.BLOCK_HEIGHT_MULT) - 26)),
		atlas) //26 = 2 * height
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
		win.Clear(color.RGBA{CONFIG.BG[0], CONFIG.BG[1], CONFIG.BG[2], CONFIG.BG[3]})
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		for i, val := range list {
			rect := imdraw.New(nil)
			if changed[i] {
				rect.Color = color.RGBA{CONFIG.CHANGED[0], CONFIG.CHANGED[1], CONFIG.CHANGED[2], CONFIG.CHANGED[3]}
			} else {
				rect.Color = color.RGBA{CONFIG.FG[0], CONFIG.FG[1], CONFIG.FG[2], CONFIG.FG[3]}
			}
			rect.Push(pixel.V(
				float64(i * CONFIG.BLOCK_WIDTH),
				float64(val * CONFIG.BLOCK_HEIGHT_MULT)))
			rect.Push(pixel.V(float64((i+1) * CONFIG.BLOCK_WIDTH), 0))
			rect.Rectangle(0)
			rect.Draw(win)
		}
		win.Update()
	}
	stop<-1
}

func show(L *lua.LState) int {
	wg.Wait()
	time.Sleep(time.Duration(CONFIG.SLEEP) * time.Millisecond)
	newTable := L.ToTable(1)
	if newTable.Len() != CONFIG.LIST_LENGTH {
		//TODO: better error handling
		panic("List of improper length given")
	}
	newList := make([]int, CONFIG.LIST_LENGTH)
	exists := make(map[int]bool, CONFIG.LIST_LENGTH)
	newTable.ForEach(func(a, b lua.LValue) {
		if b.(lua.LNumber) <= 0 || int(b.(lua.LNumber)) > CONFIG.LIST_LENGTH {
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
	if len(os.Args) >= 2 {
		FILENAME = os.Args[1]
	} else {
		fmt.Println("Usage: vsf <lua file> [settings file]\nlua file must contain a function called sort that sorts an array and use the show function to display that array on the screen\npress space to pause")
		return
	}
	configFile := ""
	if len(os.Args) >= 3 {
		configFile = os.Args[2]
	}
	var err error
	CONFIG, err = parse(configFile)
	if err != nil {
		panic(err)
	}
	changed = make([]bool, CONFIG.LIST_LENGTH)

	rand.Seed(time.Now().UnixNano())
	list = rand.Perm(CONFIG.LIST_LENGTH)
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


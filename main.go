package main

import (
	"io/ioutil"
	"time"

	"github.com/getlantern/systray"
)

const (
	// files
	fconfig = "config.csv"
	ficon   = "assets/logo.png"
	fwall   = "assets/desktop.png"
	// source
	sUnsplash = "https://source.unsplash.com/collection/"
	sKeywords = "76468395/"
)

var (
	err      error
	sz       string        // screen resolution
	de       string        // desktop environment
	fpath    string        // path of desktop.png
	delay    time.Duration = 1
	record   int           = 0
	nextFlag               = false
)

func main() {
	onExit := func() { println("exiting...") }
	// start systray process
	systray.Run(onReady, onExit)
}

func onReady() {
	// paint interface
	go func() {
		// set icon & tooltip
		icon, err := ioutil.ReadFile(ficon)
		checkErr("unable to open icon file : assets/logo.ico", err)
		systray.SetIcon(icon)
		systray.SetTooltip("Avens : Dynamic Wallpapers")

		// set menu options
		m := make(map[int]*systray.MenuItem)
		m[1] = systray.AddMenuItem("1 hr delay", "")
		m[4] = systray.AddMenuItem("4 hrs delay", "")
		m[6] = systray.AddMenuItem("6 hrs delay", "")
		m[10] = systray.AddMenuItem("10 hrs delay", "")
		systray.AddSeparator()
		nextImg := systray.AddMenuItem("Next Wallpaper", "")
		systray.AddSeparator()
		saveImg := systray.AddMenuItem("Save Wallpaper", "")
		systray.AddSeparator()
		quitTray := systray.AddMenuItem("Quit", "")

		// toggle delay options
		toggle := func(k int) {
			delay = time.Duration(k)
			updateConfig()
			for i := range m {
				if i == k {
					m[i].Disable()
				} else {
					m[i].Enable()
				}
			}
		}

		toggle(int(delay))

		// render loop
		for {
			// wait until a case satisfies
			select {
			case <-m[1].ClickedCh:
				toggle(1)
			case <-m[4].ClickedCh:
				toggle(4)
			case <-m[6].ClickedCh:
				toggle(6)
			case <-m[10].ClickedCh:
				toggle(10)
			case <-nextImg.ClickedCh:
				nextWall()
			case <-saveImg.ClickedCh:
				saveWall()
			case <-quitTray.ClickedCh:
				systray.Quit()
			}
		}
	}()

	go func() {
		for {
			time.Sleep(delay * time.Hour)
			nextWall()
		}
	}()
}

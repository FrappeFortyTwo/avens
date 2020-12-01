package main

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// log & panic errors
func checkErr(s string, e error) {
	if e != nil {
		print(s, e)
	}
}

// checks if a file exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// save current wallpaper
func saveWall() {
	if nextFlag {
		err = exec.Command("cp", "assets/desktop.png", "images/img"+strconv.Itoa(record)+".png").Run()
		checkErr("unable to perform copy operation : assets/desktop.png to images/", err)
		record++
		updateConfig()
		nextFlag = false
		println("wallpaper saved !")
	} else {
		println("already   saved !")
	}
}

// update config file for updated delay / record
func updateConfig() {

	// create file
	f, err := os.Create("config.csv")
	checkErr("unable to create config.csv file", err)
	// setup writer
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// write contents
	err = writer.Write([]string{
		"delay", strconv.Itoa(int(delay)),
	})
	checkErr("unable to update delay to config.csv : Writer Error", err)
	err = writer.Write([]string{
		"record", strconv.Itoa(record),
	})
	checkErr("unable to update record to config.csv : Writer Error", err)

	println("config  updated !")
}

// fetch, save and then set next wallpaper
func nextWall() {
	// fetch wallpaper from unsplash
	response, err := http.Get(sUnsplash + sz + sKeywords)
	checkErr("unable to fetch wallpaper : check internet connection", err)
	defer response.Body.Close()

	// save it to file i.e "desktop.png"
	f, err := os.Create("assets/desktop.png")
	checkErr("unable to create file : desktop.png", err)

	defer f.Close()
	// dump the response body to the file using io.copy
	_, err = io.Copy(f, response.Body)
	checkErr("unable to dump wallpaper into file : io.Copy error", err)

	// set the image as desktop for desired desktop environment
	setWall()

	nextFlag = true
}

// set desktop wallpaper
func setWall() {
	switch {
	// XFCE
	case strings.Contains(de, "xfce"):
		// fetch properties
		out, err := exec.Command("bash", "-c", "xfconf-query -c xfce4-desktop -l | grep \"last-image$\"").Output()
		checkErr("unable to check xfconf properties", err)
		attrs := strings.Split(string(out), "\n")
		// apply command
		for i := 0; i < len(attrs); i++ {
			exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", attrs[i], "--set", fpath).Run()
		}

	// CINNAMON
	case strings.Contains(de, "cinnamon"):
		err = exec.Command("bash", "-c", "gsettings set org.cinnamon.desktop.background picture-uri "+strconv.Quote("file://"+fpath)).Run()
		checkErr("unable to set wallpaper", err)

	// MATE
	case strings.Contains(de, "mate"):
		_ = exec.Command("bash", "-c", "gsettings set org.mate.background picture-filename "+fpath).Run()
		// disabled error : false positive

	// GNOME
	case strings.Contains(de, "gnome"):
		_ = exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri "+strconv.Quote("file://"+fpath)).Run()
		// disabled error : false positive
	}
	println("desktop updated !")
}

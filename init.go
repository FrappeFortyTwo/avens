package main

import (
	"encoding/csv"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {

	println("\nstarting avens...\n")

	// detech path for desktop wallpaper
	out, err := exec.Command("bash", "-c", "echo $PWD").Output()
	checkErr("unable to fetch pwd", err)
	fpath = strings.Replace(string(out), "\n", "", -1)
	fpath = fpath + "/" + fwall
	println("desktop path :", fpath)

	// detect desktop environment
	de = strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP"))
	println("\ndesktop environment :", de)

	// detect resolution
	out, err = exec.Command("bash", "-c", "xdpyinfo | grep -oP 'dimensions:\\s+\\K\\S+'").Output()
	if err != nil {
		println("xdpyinfo failed. trying xrandr command")
		out, err = exec.Command("bash", "-c", "xrandr |grep \\* |awk '{print $1}'").Output()
		checkErr("unable to detect resolution", err)
	}
	sz = string(out)
	sz = strings.Replace(sz, "\n", "", -1)
	println("desktop resolution  :", sz, "\n")

	// check if config.csv exists
	integrityCheck := exists(fconfig)
	println("config exists ?", integrityCheck)

	if !integrityCheck { // if the file doesn't exist
		// create file
		f, err := os.Create(fconfig)
		checkErr("unable to create config.csv", err)
		defer f.Close()
		// setup writer
		writer := csv.NewWriter(f)
		defer writer.Flush()
		// write defaults
		err = writer.Write([]string{
			"delay", "1",
		})
		checkErr("unable to write default config, writer error", err)
		err = writer.Write([]string{
			"record", "0",
		})
		checkErr("unable to write default config, writer error", err)
	} else { // If the file exists fetch data

		// open config file
		f, err := os.Open(fconfig)
		checkErr("unable to open config.csv", err)
		defer f.Close()

		// read config data
		rows, err := csv.NewReader(f).ReadAll()
		checkErr("unable to create csv.NewReader to read files", err)

		// set data for program
		tmp, err := strconv.Atoi(rows[0][1])
		checkErr("Unable to fetch delay from config file", err)
		delay = time.Duration(tmp)

		record, err = strconv.Atoi(rows[1][1])
		checkErr("Unable to fetch record from config file", err)
	}

	// check if imgs & assets folder exists
	s := make([]string, 2)
	s[0], s[1] = "images", "assets"

	for i := range s {
		integrityCheck = exists(s[i])
		// if not then create them
		if !integrityCheck {
			println(s[i], " exists ? false")
			cmd := "mkdir " + s[i]
			err = exec.Command("bash", "-c", cmd).Run()
			checkErr("unable to create missing directories", err)
		} else {
			println(s[i], "exists ? true")
		}
	}
	println()

	// check if logo.ico exists
	integrityCheck = exists(ficon)
	println(ficon, "    exists ?", integrityCheck)
	if !integrityCheck {
		logoGen()
	}
	// check if desktop.png exists
	integrityCheck = exists(fwall)
	println(fwall, " exists ?", integrityCheck, "\n")
	if !integrityCheck {
		nextWall()
	}

}

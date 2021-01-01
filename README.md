# Avens
Avens is a desktop application for dynamic wallpapers. It's written in Go and targeted for linux.

Here is a screencap of avens in action :

![Avens screencap](/README/avens.png)

## Features

- Forever 100% free
- Easy-to-use UI : runs from the taskbar / panel.
- Light weight : low on system resources.
- Option to set delay for wallpaper change.
- Option to save current wallpaper.
- Verification & restoration : restores missing files ( if any ).

Supported desktop environments : 
* XFCE
* CINNAMON
* MATE
* GNOME

## License

Aven's source code is licensed under GNU GPLv3.

## Installation

The application is madeup of a single binary so to get it running :
1. Download binary
2. Grant access : `chmod +x avens`
3. Execute it : `./avens`

That's all !

You can add it to startup for a more seamless experience.

If something doesn't work, delete all files except the binary and run it. This should recreate all necessary files with defaults. If the issue persists, you can report it here or reach out to me [here](https://twitter.com/frappefortytwo).

See [releases](github.com/FrappeFortyTwo/avens/releases) for available packages.

## Building the project

1. Install Go.
2. Install systray package : `go get github.com/getlantern/systray`
3. Clone or fork this repository.
4. Open the directory with this repo's contents.
5. Run the command : `go build` ( this should create a new file within the present working directory )
6. Finally, you can run the executable : `./avens`.

The executable creates all necessary files ( icons, config, saved images ). 

## Credits

* Wallpaper Source : [unsplash](https://unsplash.com/)
* Tray Implementation : [systray](https://github.com/getlantern/systray)

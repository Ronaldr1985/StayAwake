package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ronaldr1985/stayawake/icons/disabledicon"
	"github.com/ronaldr1985/stayawake/icons/enabledicon"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
)

func onReady() {
	var seconds int = 20
	var enabled bool = true
	systray.SetIcon(enabledicon.Data)
	systray.SetTitle("Stay Awake")
	systray.SetTooltip("Stay Awake")
	mChangeGUI := systray.AddMenuItem(
		"Change interval", "Change how often a key is pressed",
	)
	mEnabled := systray.AddMenuItemCheckbox(
		"Enabled", "Whether we should keep the screen on", true,
	)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		nextTimeToPressKeys := time.Now().Add(time.Duration(-seconds) * time.Second)
		for {
			if enabled && time.Since(nextTimeToPressKeys) > time.Duration(seconds)*time.Second {
				fmt.Println("Pressing key")

				PressAndReleaseF24Key()

				nextTimeToPressKeys = time.Now()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					mEnabled.Uncheck()
					systray.SetIcon(disabledicon.Data)
					enabled = false
				} else {
					systray.SetIcon(enabledicon.Data)
					mEnabled.Check()
					enabled = true
				}
			case <-mChangeGUI.ClickedCh:
				for {
					entered_seconds, err := zenity.Entry(
						"Enter a number of seconds",
						zenity.Title("StayAwake, enter a number of seconds"),
					)

					if err != nil {
						if strings.Contains(err.Error(), "not found") {
							err := beeep.Alert(
								"Error",
								"Error: Zenity is not installed, please install Zenity if you wish to use the change interval GUI",
								"assets/warning.png",
							)
							if err != nil {
								fmt.Fprintln(os.Stderr, "Failed to send alert failed with:", err)
							}

							fmt.Fprintln(os.Stderr, "Zenity not installed")
						}

						break
					}

					if seconds, err = strconv.Atoi(entered_seconds); err != nil {
						fmt.Fprintln(os.Stderr, "Failed to convert string to integer...")
						fmt.Fprintln(os.Stderr, "Got the following error:", err)
					} else {
						break
					}
				}
			}
		}
	}()

}

func main() {
	onExit := func() {
		return
	}

	systray.Run(onReady, onExit)
}

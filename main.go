package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"stayawake/icons/disabledicon"
	"stayawake/icons/enabledicon"
	"stayawake/utils/checkrunning"
	"stayawake/utils/presskey"

	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
)

func changeIntervalGUI() (enteredseconds int) {
	var entered_seconds int

	for true {
		entry, _, err := dlgs.Entry("StayAwake", "Enter seconds between keypresses:", "20")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error when creating window for seconds: ", err)
		}
		entered_seconds, err = strconv.Atoi(entry)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Atoi failed with error: ", err)
			entered_seconds = 20
		}
		if entered_seconds < 1 {
			dlgs.Error("StayAwake", "Must enter a number greater than 0")
		} else {
			break
		}
	}
	return entered_seconds
}

func on_ready() {
	var seconds int = 20
	var enabled bool = true

	systray.SetIcon(enabledicon.Data)
	systray.SetTitle("Stay Awake")
	systray.SetTooltip("Stay Awake")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		for {
			if enabled == true {
				presskey.PressKey(0x7C)
			}
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}()

	go func() {
		systray.AddSeparator()
		mChecked := systray.AddMenuItemCheckbox("Enabled", "Check me", true)
		mChangeInterval := systray.AddMenuItem("Change Interval", "Change interval")

		for {
			select {
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					systray.SetIcon(disabledicon.Data)
					enabled = false
				} else {
					systray.SetIcon(enabledicon.Data)
					mChecked.Check()
					enabled = true
				}
			case <-mChangeInterval.ClickedCh:
				enabled = false
				seconds = changeIntervalGUI()
				enabled = true
			}
		}
	}()
}

func on_exit() {
	fmt.Println("Cleaning up")
}

func main() {
	running, err := checkrunning.IsRunning("stayawake")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to determine if app already running")
		os.Exit(1)
	}
	if running {
		dlgs.Error("StayAwake", "StayAwake is already running")
	} else {
		systray.Run(on_ready, on_exit)
	}
}

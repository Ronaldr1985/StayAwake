package main

import (
	"fmt"
	"github.com/ronaldr1985/stayawake/icons/disabledicon"
	"github.com/ronaldr1985/stayawake/icons/enabledicon"
	"time"

	"github.com/getlantern/systray"
)

func onReady() {
	var seconds int64 = 20
	var enabled bool = true
	systray.SetIcon(enabledicon.Data)
	systray.SetTitle("Stay Awake")
	systray.SetTooltip("Stay Awake")
	systray.AddSeparator()
	mEnabled := systray.AddMenuItemCheckbox(
		"Enabled", "Whether we should keep the screen on", true,
	)
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		nextTimeToPressKeys := time.Now().Add(time.Duration(-seconds) * time.Second)
		for {
			if enabled && time.Since(nextTimeToPressKeys) > time.Duration(seconds)*time.Second {
				fmt.Println("Pressing key")

				PressKey("F24")

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

package main

/*
#include <windows.h>
#include <stdio.h>
void pressKeys(void) {
    INPUT inp;

    inp.type = INPUT_KEYBOARD;
    inp.ki.wScan = 0;
    inp.ki.time = 0;
    inp.ki.dwExtraInfo = 0;

    inp.ki.wVk = 0x7D;
    inp.ki.dwFlags = 0;
    SendInput(1, &inp, sizeof(INPUT));

    inp.ki.dwFlags = KEYEVENTF_KEYUP;
    SendInput(1, &inp, sizeof(INPUT));
}
*/
import "C"

import (
	"fmt"
	"log"
	"stayawake/icons/disabledicon"
	"stayawake/icons/enabledicon"
	"syscall"
	"time"
	"unsafe"

	"github.com/getlantern/systray"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
)

func CreateMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func main() {
	_, err := CreateMutex("StayAwake")
	if err != nil {
		log.Fatal("Application already running, quitting.")
		return
	}
	onExit := func() {
		return
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	var enabled bool = true
	var seconds int = 20
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
		systray.AddSeparator()
		mChecked := systray.AddMenuItemCheckbox("Enabled", "Check Me", true)

		for {
			select {
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					enabled = false
					systray.SetIcon(disabledicon.Data)
				} else {
					systray.SetIcon(enabledicon.Data)
					mChecked.Check()
					enabled = true
				}
			}
		}

	}()

	go func() {
		for {
			if enabled == true {
				C.pressKeys()
			}
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}()
}

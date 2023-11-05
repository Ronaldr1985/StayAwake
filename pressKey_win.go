//go:build windows
// +build windows

package main

import (
	"syscall"
	"time"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	procKeyBdEvent = user32.NewProc("keybd_event")
)

func PressAndReleaseF24Key() bool {
	flag := 0

	key := 0x87
	vkey := key + 0x80

	_, _, err := procKeyBdEvent.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0) // Press key
	if err != nil {
		return false
	}

	time.Sleep(10 * time.Millisecond)

	flag = 0x0002
	_, _, err = procKeyBdEvent.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0) // Press key
	if err != nil {
		return false
	}

	return true
}

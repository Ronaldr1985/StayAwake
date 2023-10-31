//go:build unix
// +build unix

package main

/*
#cgo LDFLAGS: -lX11 -lXtst
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/extensions/XTest.h>

int pressKey(const char *key) {
	Display *display = XOpenDisplay(NULL);
	KeyCode keyCode = 0;

	keyCode = XKeysymToKeycode(display, XStringToKeysym(key));

	XTestFakeKeyEvent(display, keyCode, False, 0);
	XFlush(display);

	XTestFakeKeyEvent(display, keyCode, True, 0);
	XFlush(display);

	XTestFakeKeyEvent(display, keyCode, False, 0);
	XFlush(display);

	XCloseDisplay(display);

	return 0;
}
*/
import "C"

import (
	"strings"
	"unsafe"
)

func PressKey(key string) bool {
	var keyCode string
	if strings.Contains(key, "F") {
		keyCode = "XK_" + key
	} else {
		keyCode = key
	}

	key_cstr := C.CString(keyCode)

	C.pressKey(key_cstr)

	C.free(unsafe.Pointer(key_cstr))

	return true
}

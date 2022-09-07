//go:build linux

package presskey

/*
#cgo LDFLAGS: -lX11 -lXtst

#include <stdio.h>
#include <unistd.h>
#include <X11/Xlib.h>
#include <X11/keysym.h>
#include <X11/extensions/XTest.h>

typedef short keycode;

void pressKey(int key)
{
	Display *display = XOpenDisplay(NULL);

	XTestFakeKeyEvent(display, key, True, 0);
	XTestFakeKeyEvent(display, key, False, 0);
	XFlush(display);

	XCloseDisplay(display);
}
*/
import "C"

const (
	KEY_0   C.int = 1
	KEY_1         = 2
	KEY_2         = 3
	KEY_3         = 4
	KEY_4         = 5
	KEY_5         = 6
	KEY_6         = 7
	KEY_7         = 8
	KEY_8         = 9
	KEY_9         = 10
	KEY_A         = 30
	KEY_B         = 48
	KEY_C         = 46
	KEY_D         = 32
	KEY_E         = 18
	KEY_F         = 33
	KEY_G         = 34
	KEY_H         = 35
	KEY_I         = 23
	KEY_J         = 36
	KEY_K         = 37
	KEY_L         = 38
	KEY_M         = 50
	KEY_N         = 49
	KEY_O         = 24
	KEY_P         = 25
	KEY_Q         = 16
	KEY_R         = 19
	KEY_S         = 31
	KEY_T         = 20
	KEY_U         = 22
	KEY_V         = 47
	KEY_W         = 17
	KEY_X         = 45
	KEY_Y         = 21
	KEY_Z         = 44
	KEY_F1        = 59
	KEY_F2        = 60
	KEY_F3        = 61
	KEY_F4        = 62
	KEY_F5        = 63
	KEY_F6        = 64
	KEY_F7        = 65
	KEY_F8        = 66
	KEY_F9        = 67
	KEY_F10       = 68
	KEY_F11       = 87
	KEY_F12       = 88
	KEY_F13       = 183
	KEY_F14       = 184
	KEY_F15       = 185
	KEY_F16       = 186
	KEY_F17       = 187
	KEY_F18       = 188
	KEY_F19       = 189
	KEY_F20       = 190
	KEY_F21       = 191
	KEY_F22       = 192
	KEY_F23       = 193
	KEY_F24       = 194
)

func PressKey(key C.int) {
	C.pressKey(key)
}

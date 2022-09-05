//go:build windows

/* Keycodes
    	0x30 	0 key
	    0x31 	1 key
	    0x32 	2 key
	    0x33 	3 key
	    0x34 	4 key
	    0x35 	5 key
	    0x36 	6 key
	    0x37 	7 key
	    0x38 	8 key
	    0x39 	9 key
	    0x41 	A key
	    0x42 	B key
	    0x43 	C key
	    0x44 	D key
	    0x45 	E key
	    0x46 	F key
	    0x47 	G key
	    0x48 	H key
	    0x49 	I key
	    0x4A 	J key
	    0x4B 	K key
	    0x4C 	L key
	    0x4D 	M key
	    0x4E 	N key
	    0x4F 	O key
	    0x50 	P key
	    0x51 	Q key
	    0x52 	R key
	    0x53 	S key
	    0x54 	T key
	    0x55 	U key
	    0x56 	V key
	    0x57 	W key
	    0x58 	X key
	    0x59 	Y key
	    0x5A 	Z key
VK_F1 	0x70 	F1 key
VK_F2 	0x71 	F2 key
VK_F3 	0x72 	F3 key
VK_F4 	0x73 	F4 key
VK_F5 	0x74 	F5 key
VK_F6 	0x75 	F6 key
VK_F7 	0x76 	F7 key
VK_F8 	0x77 	F8 key
VK_F9 	0x78 	F9 key
VK_F10 	0x79 	F10 key
VK_F11 	0x7A 	F11 key
VK_F12 	0x7B 	F12 key
VK_F13 	0x7C 	F13 key
VK_F14 	0x7D 	F14 key
VK_F15 	0x7E 	F15 key
VK_F16 	0x7F 	F16 key
VK_F17 	0x80 	F17 key
VK_F18 	0x81 	F18 key
VK_F19 	0x82 	F19 key
VK_F20 	0x83 	F20 key
VK_F21 	0x84 	F21 key
VK_F22 	0x85 	F22 key
VK_F23 	0x86 	F23 key
VK_F24 	0x87 	F24 key
*/

package presskey

/*
#include <windows.h>
#include <stdio.h>
void pressKey(short keycode) {
    INPUT inp;
    inp.type = INPUT_KEYBOARD;
    inp.ki.wScan = 0;
    inp.ki.time = 0;
    inp.ki.dwExtraInfo = 0;
    inp.ki.wVk = keycode;
    inp.ki.dwFlags = 0;
    SendInput(1, &inp, sizeof(INPUT));
    inp.ki.dwFlags = KEYEVENTF_KEYUP;
    SendInput(1, &inp, sizeof(INPUT));
}
*/
import "C"

func PressKey(keycode C.short) {
    C.pressKey(keycode)
}
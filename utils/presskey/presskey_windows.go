//go:build windows

package presskey

type keycode c.Short

/* Keycodes */
const (
	KEY_0	keycode =	0x30
	KEY_1	=	0x31
	KEY_2	=	0x32
	KEY_3	=	0x33
	KEY_4	=	0x34
	KEY_5	=	0x35
	KEY_6	=	0x36
	KEY_7	=	0x37
	KEY_8	=	0x38
	KEY_9	=	0x39
	KEY_A	=	0x41
	KEY_B	=	0x42
	KEY_C	=	0x43
	KEY_D	=	0x44
	KEY_E	=	0x45
	KEY_F	=	0x46
	KEY_G	=	0x47
	KEY_H	=	0x48
	KEY_I	=	0x49
	KEY_J	=	0x4A
	KEY_K	=	0x4B
	KEY_L	=	0x4C
	KEY_M	=	0x4D
	KEY_N	=	0x4E
	KEY_O	=	0x4F
	KEY_P	=	0x50
	KEY_Q	=	0x51
	KEY_R	=	0x52
	KEY_S	=	0x53
	KEY_T	=	0x54
	KEY_U	=	0x55
	KEY_V	=	0x56
	KEY_W	=	0x57
	KEY_X	=	0x58
	KEY_Y	=	0x59
	KEY_Z	=	0x5A
	KEY_F1	=	0x70
	KEY_F2	=	0x71
	KEY_F3	=	0x72
	KEY_F4	=	0x73
	KEY_F5	=	0x74
	KEY_F6	=	0x75
	KEY_F7	=	0x76
	KEY_F8	=	0x77
	KEY_F9	=	0x78
	KEY_F10	=	0x79
	KEY_F11	=	0x7A
	KEY_F12	=	0x7B
	KEY_F13	=	0x7C
	KEY_F14	=	0x7D
	KEY_F15	=	0x7E
	KEY_F16	=	0x7F
	KEY_F17	=	0x80
	KEY_F18	=	0x81
	KEY_F19	=	0x82
	KEY_F20	=	0x83
	KEY_F21	=	0x84
	KEY_F22	=	0x85
	KEY_F23	=	0x86
	KEY_F24	=	0x87
)


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

func PressKey(key keycode) {
	C.pressKey(keycode)
}

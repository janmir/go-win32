package win32

import (
	"testing"
	"time"
	"unsafe"

	util "github.com/janmir/go-util"
)

func TestEnumWindows(t *testing.T) {
	win := New()
	cb := func(hwnd HWND, p uintptr) uintptr {
		text := win.GetWindowText(hwnd)
		class := win.GetClassName(hwnd)

		util.Logger("Title: %s, Class: %s", text, class)

		return 1
	}
	win.EnumWindows(cb, 0)
}

func TestGetForegroundWindow(t *testing.T) {
	time.Sleep(time.Second * 2)
	win := New()
	hwnd := win.GetForegroundWindow()
	// bytes := make([]uint16, 200)
	text := win.GetWindowText(hwnd)
	class := win.GetClassName(hwnd)

	util.Logger("Title: %s, Class: %s", text, class)
}

func TestGetWindowRect(t *testing.T) {
	win := New()
	hwnd := win.GetForegroundWindow()
	rect := win.GetWindowRect(hwnd)

	util.Logger("Rect: %+v", rect)
}

func TestGetDesktopWindow(t *testing.T) {
	win := New()
	hwnd := win.GetDesktopWindow()
	rect := win.GetWindowRect(hwnd)

	util.Logger("Desktop Rect: %+v", rect)
}

func TestGetDesktopWindowConstant(t *testing.T) {
	return

	win := New()
	rect := win.GetWindowRect(HWND_DESKTOP)

	util.Logger("Desktop Constant Rect: %+v", rect)
}

func TestMapWindowPoints(t *testing.T) {
	win := New()
	desk := win.GetDesktopWindow()
	fore := win.GetForegroundWindow()
	points, ok := win.MapWindowPoints(desk, fore)

	if !ok {
		t.Fail()
	}

	util.Logger("Desktop to Foreground Points: %+v", points)
}

func TestClientToScreen(t *testing.T) {
	win := New()
	fore := win.GetForegroundWindow()

	newPoints, ok := win.ClientToScreen(fore)
	if !ok {
		t.Fail()
	}

	util.Logger("Client to Screen Points: %+v", newPoints)
}

func TestSendInput(t *testing.T) {
	win := New()
	input := make([]KEYBD_INPUT, 0)
	down := KEYBDINPUT{
		WVk:         160, //keycode
		WScan:       0,
		DwFlags:     KEYEVENTF_KEYDOWN,
		Time:        0,
		DwExtraInfo: 0,
	}
	up := KEYBDINPUT{
		WVk:         160, //keycode
		WScan:       0,
		DwFlags:     KEYEVENTF_KEYUP,
		Time:        0,
		DwExtraInfo: 0,
	}
	down2 := KEYBDINPUT{
		WVk:         66, //keycode
		WScan:       0,
		DwFlags:     KEYEVENTF_KEYDOWN,
		Time:        0,
		DwExtraInfo: 0,
	}
	up2 := KEYBDINPUT{
		WVk:         66, //keycode
		WScan:       0,
		DwFlags:     KEYEVENTF_KEYUP,
		Time:        0,
		DwExtraInfo: 0,
	}
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   down,
	})
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   down2,
	})
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   up2,
	})
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   up,
	})
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   down2,
	})
	input = append(input, KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   up2,
	})
	util.Logger("Input: %+v, Size: %d", input, unsafe.Sizeof(KEYBD_INPUT{}))
	count := win.SendInput(len(input), unsafe.Pointer(&input[0]), unsafe.Sizeof(KEYBD_INPUT{}))

	util.Logger("KeyCount: ", count)
}

func TestSetWinEventHook(t *testing.T) {
	return
	util.Logger("starting")
	win := New()
	hinst := win.GetModuleHandle("")

	util.Logger("Module Handle: %+v", hinst)

	winEvHook := win.SetWinEventHook(
		EVENT_MIN,
		EVENT_MAX, 0,
		func(hook HHOOK, event uint32, hwnd HWND, idObject, idChild, idEventThread, dwmsEventTime uint32) uintptr {
			switch event {
			case EVENT_SYSTEM_FOREGROUND:
				fallthrough
			case EVENT_SYSTEM_MOVESIZESTART:
				util.Logger("Event: %x\n", event)

			}
			return 0
		}, 0, 0,
		WINEVENT_OUTOFCONTEXT|WINEVENT_SKIPOWNPROCESS)
	util.Logger("Windows Event Hook: %+v", winEvHook)

	for {

		if msg, m := win.GetMessage(0, 0, 0); m {
			msg, _ = win.TranslateMessage()
			msg, _ = win.DispatchMessage()
			util.Logger("Windows Message: %+v", msg)
		}
	}
	win.UnhookWinEvent(winEvHook)
}

func TestGetGUIThreadInfo(t *testing.T) {
	win := New()
	info, ok := win.GetGUIThreadInfo(0)
	if !ok {
		t.Fail()
	}
	newPoints, ok := win.ClientToScreen(info.HWNDCaret)

	util.Logger("GUI Info: %+v, %+v", info, newPoints)
}

func TestKeyboard(t *testing.T) {
	win := New()
	hkl := win.GetKeyboardLayout(0)
	util.Logger("Keyboard: ", hkl)
	hi, lo := win.VkKeyScanEx('7', hkl)
	util.Logger("Hi/Low: %d, %d", hi, lo)
}

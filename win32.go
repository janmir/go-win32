package win32

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/janmir/go-util"
)

var (
	user32   uintptr
	kernel32 uintptr
)

//Win32 collection of win32 api
//functions for windows
type Win32 struct {
	//User 32
	getDesktopWindow            uintptr
	getWindowRect               uintptr
	enumWindows                 uintptr
	getWindowText               uintptr
	getClassName                uintptr
	getForegroundWindow         uintptr
	mapWindowPoints             uintptr
	clientToScreen              uintptr
	sendInput                   uintptr
	setWinEventHook             uintptr
	unhookWinEvent              uintptr
	getMessage                  uintptr
	translateMessage            uintptr
	dispatchMessage             uintptr
	getGUIThreadInfo            uintptr
	vkKeyScanEx                 uintptr
	loadKeyboardLayout          uintptr
	getKeyboardLayout           uintptr
	postMessage                 uintptr
	postQuitMessage             uintptr
	changeWindowMessageFilterEx uintptr
	getWindowThreadProcessID    uintptr
	isGUIThread                 uintptr
	getCaretPos                 uintptr
	bringWindowToTop            uintptr
	setActiveWindow             uintptr
	setForegroundWindow         uintptr
	findWindowA                 uintptr
	findWindowExW               uintptr
	setWindowLongW              uintptr

	//Kernel 32
	getModuleHandle  uintptr
	setPriorityClass uintptr
	getCurrentThread uintptr
	getLastError     uintptr
}

func init() {
	var err error

	//Load user32 dll
	user32, err = MustLoadLibrary("user32.dll")
	util.Catch(err, "Unable to load user32.dll")

	kernel32, err = MustLoadLibrary("kernel32.dll")
	util.Catch(err, "Unable to load kernel32.dll")
}

//New creates a new instance of Win32 Object
func New() Win32 {
	return Win32{
		//User 32
		getDesktopWindow:            MustGetProcAddress(user32, "GetDesktopWindow"),
		getWindowRect:               MustGetProcAddress(user32, "GetWindowRect"),
		enumWindows:                 MustGetProcAddress(user32, "EnumWindows"),
		getWindowText:               MustGetProcAddress(user32, "GetWindowTextW"),
		getClassName:                MustGetProcAddress(user32, "GetClassNameW"),
		getForegroundWindow:         MustGetProcAddress(user32, "GetForegroundWindow"),
		mapWindowPoints:             MustGetProcAddress(user32, "MapWindowPoints"),
		clientToScreen:              MustGetProcAddress(user32, "ClientToScreen"),
		sendInput:                   MustGetProcAddress(user32, "SendInput"),
		setWinEventHook:             MustGetProcAddress(user32, "SetWinEventHook"),
		unhookWinEvent:              MustGetProcAddress(user32, "UnhookWinEvent"),
		getMessage:                  MustGetProcAddress(user32, "GetMessageW"),
		translateMessage:            MustGetProcAddress(user32, "TranslateMessage"),
		dispatchMessage:             MustGetProcAddress(user32, "DispatchMessageW"),
		getGUIThreadInfo:            MustGetProcAddress(user32, "GetGUIThreadInfo"),
		vkKeyScanEx:                 MustGetProcAddress(user32, "VkKeyScanExW"),
		loadKeyboardLayout:          MustGetProcAddress(user32, "LoadKeyboardLayoutW"),
		getKeyboardLayout:           MustGetProcAddress(user32, "GetKeyboardLayout"),
		postMessage:                 MustGetProcAddress(user32, "PostMessageW"),
		changeWindowMessageFilterEx: MustGetProcAddress(user32, "ChangeWindowMessageFilterEx"),
		getWindowThreadProcessID:    MustGetProcAddress(user32, "GetWindowThreadProcessId"),
		isGUIThread:                 MustGetProcAddress(user32, "IsGUIThread"),
		getCaretPos:                 MustGetProcAddress(user32, "GetCaretPos"),
		bringWindowToTop:            MustGetProcAddress(user32, "BringWindowToTop"),
		setActiveWindow:             MustGetProcAddress(user32, "SetActiveWindow"),
		setForegroundWindow:         MustGetProcAddress(user32, "SetForegroundWindow"),
		findWindowA:                 MustGetProcAddress(user32, "FindWindowA"),
		findWindowExW:               MustGetProcAddress(user32, "FindWindowExW"),
		setWindowLongW:              MustGetProcAddress(user32, "SetWindowLongW"),

		//Kernel 32
		getModuleHandle:  MustGetProcAddress(kernel32, "GetModuleHandleW"),
		setPriorityClass: MustGetProcAddress(kernel32, "SetPriorityClass"),
		getCurrentThread: MustGetProcAddress(kernel32, "GetCurrentThread"),
		getLastError:     MustGetProcAddress(kernel32, "GetLastError"),
	}
}

//GetWindowRect Retrieves the dimensions of the
//bounding rectangle of the specified window.
//The dimensions are given in screen coordinates
//that are relative to the upper-left corner of the screen.
func (win Win32) GetWindowRect(hWnd HWND) RECT {
	rect := new(RECT)
	ret, _, err := syscall.Syscall(win.getWindowRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	if err != 0 || ret == 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return *rect
}

//EnumWindows Enumerates all top-level windows on the screen by passing
//the handle to each window, in turn, to an application-defined
//callback function. EnumWindows continues until the last
//top-level window is enumerated or the callback function returns FALSE.
func (win Win32) EnumWindows(callback WNDENUMPROC, lParam uintptr) bool {
	ret, _, _ := syscall.Syscall(win.enumWindows, 2,
		uintptr(syscall.NewCallback(callback)),
		lParam,
		0)

	return ret != 0
}

//GetWindowText Copies the text of the specified window's title
//bar (if it has one) into a buffer. If the specified window is
//a control, the text of the control is copied. However,
//GetWindowText cannot retrieve the text of a control in another application.
func (win Win32) GetWindowText(hWnd HWND) string {
	ll := 100
	text := make([]uint16, ll)
	len, _, err := syscall.Syscall(win.getWindowText, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&text[0])),
		uintptr(int32(ll)))

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
	_ = len
	return UintToString(text)
}

//GetClassName Retrieves the name of the
//class to which the specified window belongs.
func (win Win32) GetClassName(hwnd HWND) string {
	className := make([]uint16, 255)
	len, _, err := syscall.Syscall(win.getClassName, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&className[0])),
		uintptr(int32(len(className))))

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
	_ = len
	return UintToString(className)
}

//GetForegroundWindow Retrieves a handle to the foreground
// window (the window with which the user is currently working).
//The system assigns a slightly higher priority to the thread
//that creates the foreground window than it does to other threads.
func (win Win32) GetForegroundWindow() HWND {
	hwnd, _, err := syscall.Syscall(win.getForegroundWindow, 0,
		0,
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return HWND(hwnd)
}

//GetDesktopWindow Retrieves a handle to the desktop window.
//The desktop window covers the entire screen. The desktop
//window is the area on top of which other windows are painted.
func (win Win32) GetDesktopWindow() HWND {
	hwnd, _, err := syscall.Syscall(win.getDesktopWindow, 0,
		0,
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return HWND(hwnd)
}

//MapWindowPoints function converts (maps) a set of points from a
//coordinate space relative to one window to a coordinate space
//relative to another window.
func (win Win32) MapWindowPoints(from, to HWND) (POINT, bool) {
	point := new(POINT)
	ret, _, _ := syscall.Syscall6(win.mapWindowPoints, 4,
		uintptr(from),
		uintptr(to),
		uintptr(unsafe.Pointer(point)),
		2, 0, 0)

	return *point, ret != 0
}

//ClientToScreen function converts the client-area coordinates
//of a specified point to screen coordinates.
func (win Win32) ClientToScreen(hwnd HWND) (POINT, bool) {
	point := new(POINT)
	ret, _, _ := syscall.Syscall(win.clientToScreen, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(point)),
		0)

	return *point, ret != 0
}

//SendInput Synthesizes keystrokes, mouse motions, and button clicks.
//nInputs: The number of structures in the pInputs array.
//pInputs: An array of INPUT structures. Each structure represents an event
//		   to be inserted into the keyboard or mouse input stream.
//		   expects a unsafe.Pointer to a slice of MOUSE_INPUT or KEYBD_INPUT
//		   or HARDWARE_INPUT structs.
//cbSize: The size, in bytes, of an INPUT structure. If cbSize is not the size of
//		  an INPUT structure, the function fails.
//Returns the number of events sent
func (win Win32) SendInput(nInputs int, pInputs unsafe.Pointer, cbSize uintptr) uint32 {
	ret, _, err := syscall.Syscall(win.sendInput, 3,
		uintptr(nInputs),
		uintptr(pInputs),
		cbSize)

	if err > 5 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return uint32(ret)
}

//SetWinEventHook Sets an event hook function for a range of events.
func (win Win32) SetWinEventHook(eventMin, eventMax uint32,
	hmodWinEventProc HMODULE,
	winEventCallback HOOKPROC,
	idProcess, idThread, dwFlags uint32) HHOOK {
	ret, _, _ := syscall.Syscall9(win.setWinEventHook, 7,
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		uintptr(syscall.NewCallback(winEventCallback)),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags), 0, 0)
	return HHOOK(ret)
}

//UnhookWinEvent Removes an event hook function created by a
//previous call to SetWinEventHook.
func (win Win32) UnhookWinEvent(hWinEventHook HHOOK) bool {
	ret, _, _ := syscall.Syscall(win.unhookWinEvent, 1,
		uintptr(hWinEventHook),
		0, 0)
	return ret != 0
}

//GetMessage Retrieves a message from the calling thread's message queue.
//The function dispatches incoming sent messages until a posted message
//is available for retrieval.
func (win Win32) GetMessage(hWnd HWND, msgFilterMin, msgFilterMax uint32) (MSG, bool) {
	msg := new(MSG)
	ret, _, err := syscall.Syscall6(win.getMessage, 4,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
	util.Logger(".")
	return *msg, ret > 0
}

//TranslateMessage Translates virtual-key messages into character
//messages. The character messages are posted to the calling
//thread's message queue, to be read the next time the thread
//calls the GetMessage or PeekMessage function.
func (win Win32) TranslateMessage() (MSG, bool) {
	msg := new(MSG)
	ret, _, err := syscall.Syscall(win.translateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return *msg, ret != 0
}

//DispatchMessage Dispatches a message to a window procedure.
//It is typically used to dispatch a message retrieved by the
//GetMessage function.
func (win Win32) DispatchMessage() (MSG, uintptr) {
	msg := new(MSG)
	ret, _, err := syscall.Syscall(win.dispatchMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return *msg, ret
}

//PostMessage Places (posts) a message in the message queue associated
//with the thread that created the specified window and returns without
//waiting for the thread to process the message.
func (win Win32) PostMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(win.postMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

//PostQuitMessage Indicates to the system that a thread has made a
//request to terminate (quit). It is typically used in
//response to a WM_DESTROY message.
func (win Win32) PostQuitMessage(exitCode int32) {
	syscall.Syscall(win.postQuitMessage, 1,
		uintptr(exitCode),
		0,
		0)
}

//GetGUIThreadInfo Retrieves information about the active window or a specified GUI thread.
//idThread: The identifier for the thread for which information is to be retrieved.
//			To retrieve this value, use the GetWindowThreadProcessId function.
//			If this parameter is NULL, the function returns information for
//			the foreground thread.
func (win Win32) GetGUIThreadInfo(idThread uint32) (GUITHREADINFO, bool) {
	gui := new(GUITHREADINFO)
	size := uint32(unsafe.Sizeof(GUITHREADINFO{}))
	gui.cbSize = size

	ret, _, err := syscall.Syscall(win.getGUIThreadInfo, 2,
		uintptr(idThread),
		uintptr(unsafe.Pointer(gui)),
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return *gui, ret != 0
}

//VkKeyScanEx Translates a character to the corresponding
//virtual-key code and shift state. The function translates
//the character using the input language and physical keyboard
//layout identified by the input locale identifier.
//ch: The character to be translated into a virtual-key code.
//dwhkl: Input locale identifier used to translate the character.
//		 This parameter can be any input locale identifier
//		 previously returned by the LoadKeyboardLayout function.
//Return: low-order byte of the return value contains the virtual-key
//		  code and the high-order byte contains the shift state,
//		  which can be a combination of the following flag bits.
//		  1 - SHIFT , 2 - CTRL, 4 - ALT,  8 - Hankaku, 16 - Reserved(defined by the keyboard layout driver).
func (win Win32) VkKeyScanEx(ch rune, dwhkl HKL) (int, int) {
	ret, _, err := syscall.Syscall(win.vkKeyScanEx, 2,
		uintptr(ch),
		uintptr(dwhkl),
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
	short := int16(ret)
	//1,01000001
	hi := short >> 8  //modifier key
	lo := byte(short) // & 15 //virtual keycode

	// util.Logger("Inner Hi/Low: %b,  %d, %d", short, hi, lo)

	return int(hi), int(lo)
}

//LoadKeyboardLayout Loads a new input locale identifier
//(formerly called the keyboard layout) into the system.
func (win Win32) LoadKeyboardLayout(pwszKLID string, flags uint32) HKL {
	ret, _, err := syscall.Syscall(win.loadKeyboardLayout, 2,
		uintptr(unsafe.Pointer(StringToUint(pwszKLID))), //LPCWSTR
		uintptr(flags),
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return HKL(ret)
}

//GetKeyboardLayout Retrieves the active input locale identifier (formerly called the keyboard layout).
//idThread: The identifier of the thread to query, or 0 for the current thread.
func (win Win32) GetKeyboardLayout(idThread uint32) HKL {
	ret, _, err := syscall.Syscall(win.getKeyboardLayout, 1,
		uintptr(idThread),
		0,
		0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
	return HKL(ret)
}

//ChangeWindowMessageFilterEx Modifies the User Interface
//Privilege Isolation (UIPI) message filter for a specified window.
func (win Win32) ChangeWindowMessageFilterEx(hwnd HWND, message, action uint32) bool {
	ret, _, err := syscall.Syscall6(win.setPriorityClass, 4,
		uintptr(hwnd),
		uintptr(message),
		uintptr(action),
		0, // PCHANGEFILTERSTRUCT pChangeFilterStruct
		0, 0)

	if err != 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return ret != 0
}

//GetWindowThreadProcessID Retrieves the identifier of the
//thread that created the specified window and, optionally,
//the identifier of the process that created the window.
func (win Win32) GetWindowThreadProcessID(hwnd HWND) (HANDLE, uint32) {
	var processID uint32
	ret, _, _ := syscall.Syscall(win.getWindowThreadProcessID, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processID)),
		0)

	return HANDLE(ret), processID
}

//IsGUIThread Determines whether the calling thread is already a
//GUI thread. It can also optionally convert the thread to a GUI thread.
func (win Win32) IsGUIThread(bConvert uint32) bool {
	ret, _, _ := syscall.Syscall(win.isGUIThread, 1,
		uintptr(bConvert),
		0, 0)
	return ret != 0
}

//GetCaretPos Copies the caret's position to the specified POINT
//structure.
func (win Win32) GetCaretPos() POINT {
	var p POINT
	ret, _, err := syscall.Syscall(win.getCaretPos, 1,
		uintptr(unsafe.Pointer(&p)),
		0,
		0)

	if err != 0 || ret == 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}

	return p

}

//GetModuleHandle Retrieves a module handle for the specified
//module. The module must have been loaded by the calling process.
func (win Win32) GetModuleHandle(name string) HINSTANCE {
	var mod uintptr
	if name == "" {
		mod = 0
	} else {
		mod = uintptr(unsafe.Pointer(StringToUint(name)))
	}
	ret, _, _ := syscall.Syscall(win.getModuleHandle, 1,
		uintptr(mod),
		0, 0)
	return HINSTANCE(ret)
}

//SetPriorityClass Sets the priority class for the specified
//process. This value together with the priority value of
//each thread of the process determines each thread's base priority level.
func (win Win32) SetPriorityClass(hProcess HANDLE, dwPriorityClass uint32) bool {
	ret, _, _ := syscall.Syscall(win.setPriorityClass, 2,
		uintptr(hProcess),
		uintptr(dwPriorityClass),
		0)
	return ret != 0
}

//GetCurrentThread Retrieves a pseudo handle for the calling thread.
func (win Win32) GetCurrentThread() HANDLE {
	ret, _, _ := syscall.Syscall(win.getCurrentThread, 0,
		0,
		0,
		0)

	return HANDLE(ret)
}

//GetLastError Retrieves the calling thread's last-error code
//value. The last-error code is maintained on a per-thread basis.
//Multiple threads do not overwrite each other's last-error code.
func (win Win32) GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(win.getLastError, 0,
		0,
		0,
		0)
	return uint32(ret)
}

//BringWindowToTop Brings the specified window to the top of
//the Z order. If the window is a top-level window, it is
//activated. If the window is a child window, the top-level
//parent window associated with the child window is activated.
func (win Win32) BringWindowToTop(hWnd HWND) {
	ret, _, err := syscall.Syscall(win.bringWindowToTop, 1,
		uintptr(hWnd),
		0,
		0)

	if err != 0 || ret == 0 {
		util.Catch(fmt.Errorf("error: %d", err))
	}
}

//SetActiveWindow Activates a window. The window must
//be attached to the calling thread's message queue.
func (win Win32) SetActiveWindow(hWnd HWND) {
	ret, _, err := syscall.Syscall(win.setActiveWindow, 1,
		uintptr(hWnd),
		0,
		0)

	if err != 0 || ret == 0 {
		util.Logger(fmt.Errorf("error: %d, return: %x", err, ret))
	}
}

//SetForegroundWindow Brings the thread that created the
//specified window into the foreground and activates the window.
func (win Win32) SetForegroundWindow(hWnd HWND) bool {
	ret, _, err := syscall.Syscall(win.setForegroundWindow, 1,
		uintptr(hWnd),
		0,
		0)

	if err > 5 || ret == 0 {
		util.Logger(fmt.Errorf("error: %d, return: %x", err, ret))
	}

	return ret != 0
}

//FindWindowA Retrieves a handle to the top-level window whose class name
//and window name match the specified strings. This function does not search
//child windows. This function does not perform a case-sensitive search.
func (win Win32) FindWindowA(className, windowName string) HWND {
	var lpClassName, lpWindowName *uint16
	if len(className) > 0 {
		lpClassName = StringToUint(className)
	} else {
		lpClassName = nil
	}
	if len(windowName) > 0 {
		lpWindowName = StringToUint(windowName)
	} else {
		lpWindowName = nil
	}
	ret, _, err := syscall.Syscall(win.findWindowA, 2,
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)

	if err != 0 {
		util.Logger(fmt.Errorf("error: %d", err))
	}

	return HWND(ret)
}

//FindWindowExW Retrieves a handle to a window whose class name and window name match the
//specified strings. The function searches child windows, beginning with the one following
//the specified child window. This function does not perform a case-sensitive search.
func (win Win32) FindWindowExW(hwndParent, hwndChildAfter HWND, className, windowName string) HWND {
	var lpClassName, lpWindowName *uint16
	if len(className) > 0 {
		lpClassName = StringToUint(className)
	} else {
		lpClassName = nil
	}
	if len(windowName) > 0 {
		lpWindowName = StringToUint(windowName)
	} else {
		lpWindowName = nil
	}
	ret, _, err := syscall.Syscall6(win.findWindowExW, 4,
		uintptr(hwndParent),
		uintptr(hwndChildAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0, 0)

	if err != 0 {
		util.Logger(fmt.Errorf("error: %d", err))
	}

	return HWND(ret)
}

//SetWindowLong ...
func (win Win32) SetWindowLong(hwnd HWND, index int, value uint32) uint32 {
	ret, _, _ := syscall.Syscall(win.setWindowLongW, 3,
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return uint32(ret)
}

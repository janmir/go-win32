package win32

import "syscall"

// Predefined window handles
const (
	HWND_BROADCAST = HWND(0xFFFF)
	HWND_BOTTOM    = HWND(1)
	HWND_NOTOPMOST = ^HWND(1) // -2
	HWND_TOP       = HWND(0)
	HWND_TOPMOST   = ^HWND(0) // -1
	HWND_DESKTOP   = HWND(0)
	HWND_MESSAGE   = ^HWND(2) // -3
)

//Window Priorities
const (
	ABOVE_NORMAL_PRIORITY_CLASS   = 0x00008000
	BELOW_NORMAL_PRIORITY_CLASS   = 0x00004000
	HIGH_PRIORITY_CLASS           = 0x00000080
	IDLE_PRIORITY_CLASS           = 0x00000040
	NORMAL_PRIORITY_CLASS         = 0x00000020
	PROCESS_MODE_BACKGROUND_BEGIN = 0x00100000
	PROCESS_MODE_BACKGROUND_END   = 0x00200000
	REALTIME_PRIORITY_CLASS       = 0x00000100
)

// INPUT Type
const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

// GetWindowLong and GetWindowLongPtr constants
const (
	GWL_EXSTYLE     = -20
	GWL_STYLE       = -16
	GWL_WNDPROC     = -4
	GWLP_WNDPROC    = -4
	GWL_HINSTANCE   = -6
	GWLP_HINSTANCE  = -6
	GWL_HWNDPARENT  = -8
	GWLP_HWNDPARENT = -8
	GWL_ID          = -12
	GWLP_ID         = -12
	GWL_USERDATA    = -21
	GWLP_USERDATA   = -21
)

// Extended window style constants
const (
	WS_EX_DLGMODALFRAME    = 0X00000001
	WS_EX_NOPARENTNOTIFY   = 0X00000004
	WS_EX_TOPMOST          = 0X00000008
	WS_EX_ACCEPTFILES      = 0X00000010
	WS_EX_TRANSPARENT      = 0X00000020
	WS_EX_MDICHILD         = 0X00000040
	WS_EX_TOOLWINDOW       = 0X00000080
	WS_EX_WINDOWEDGE       = 0X00000100
	WS_EX_CLIENTEDGE       = 0X00000200
	WS_EX_CONTEXTHELP      = 0X00000400
	WS_EX_RIGHT            = 0X00001000
	WS_EX_LEFT             = 0X00000000
	WS_EX_RTLREADING       = 0X00002000
	WS_EX_LTRREADING       = 0X00000000
	WS_EX_LEFTSCROLLBAR    = 0X00004000
	WS_EX_RIGHTSCROLLBAR   = 0X00000000
	WS_EX_CONTROLPARENT    = 0X00010000
	WS_EX_STATICEDGE       = 0X00020000
	WS_EX_APPWINDOW        = 0X00040000
	WS_EX_OVERLAPPEDWINDOW = 0X00000100 | 0X00000200
	WS_EX_PALETTEWINDOW    = 0X00000100 | 0X00000080 | 0X00000008
	WS_EX_LAYERED          = 0X00080000
	WS_EX_NOINHERITLAYOUT  = 0X00100000
	WS_EX_LAYOUTRTL        = 0X00400000
	WS_EX_NOACTIVATE       = 0X08000000
)

// Window style constants
const (
	WS_OVERLAPPED       = 0X00000000
	WS_POPUP            = 0X80000000
	WS_CHILD            = 0X40000000
	WS_MINIMIZE         = 0X20000000
	WS_VISIBLE          = 0X10000000
	WS_DISABLED         = 0X08000000
	WS_CLIPSIBLINGS     = 0X04000000
	WS_CLIPCHILDREN     = 0X02000000
	WS_MAXIMIZE         = 0X01000000
	WS_CAPTION          = 0X00C00000
	WS_BORDER           = 0X00800000
	WS_DLGFRAME         = 0X00400000
	WS_VSCROLL          = 0X00200000
	WS_HSCROLL          = 0X00100000
	WS_SYSMENU          = 0X00080000
	WS_THICKFRAME       = 0X00040000
	WS_GROUP            = 0X00020000
	WS_TABSTOP          = 0X00010000
	WS_MINIMIZEBOX      = 0X00020000
	WS_MAXIMIZEBOX      = 0X00010000
	WS_TILED            = 0X00000000
	WS_ICONIC           = 0X20000000
	WS_SIZEBOX          = 0X00040000
	WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
	WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
	WS_CHILDWINDOW      = 0X40000000
)

// GetSystemMetrics constants
const (
	SM_CXSCREEN             = 0
	SM_CYSCREEN             = 1
	SM_CXVSCROLL            = 2
	SM_CYHSCROLL            = 3
	SM_CYCAPTION            = 4
	SM_CXBORDER             = 5
	SM_CYBORDER             = 6
	SM_CXDLGFRAME           = 7
	SM_CYDLGFRAME           = 8
	SM_CYVTHUMB             = 9
	SM_CXHTHUMB             = 10
	SM_CXICON               = 11
	SM_CYICON               = 12
	SM_CXCURSOR             = 13
	SM_CYCURSOR             = 14
	SM_CYMENU               = 15
	SM_CXFULLSCREEN         = 16
	SM_CYFULLSCREEN         = 17
	SM_CYKANJIWINDOW        = 18
	SM_MOUSEPRESENT         = 19
	SM_CYVSCROLL            = 20
	SM_CXHSCROLL            = 21
	SM_DEBUG                = 22
	SM_SWAPBUTTON           = 23
	SM_RESERVED1            = 24
	SM_RESERVED2            = 25
	SM_RESERVED3            = 26
	SM_RESERVED4            = 27
	SM_CXMIN                = 28
	SM_CYMIN                = 29
	SM_CXSIZE               = 30
	SM_CYSIZE               = 31
	SM_CXFRAME              = 32
	SM_CYFRAME              = 33
	SM_CXMINTRACK           = 34
	SM_CYMINTRACK           = 35
	SM_CXDOUBLECLK          = 36
	SM_CYDOUBLECLK          = 37
	SM_CXICONSPACING        = 38
	SM_CYICONSPACING        = 39
	SM_MENUDROPALIGNMENT    = 40
	SM_PENWINDOWS           = 41
	SM_DBCSENABLED          = 42
	SM_CMOUSEBUTTONS        = 43
	SM_CXFIXEDFRAME         = SM_CXDLGFRAME
	SM_CYFIXEDFRAME         = SM_CYDLGFRAME
	SM_CXSIZEFRAME          = SM_CXFRAME
	SM_CYSIZEFRAME          = SM_CYFRAME
	SM_SECURE               = 44
	SM_CXEDGE               = 45
	SM_CYEDGE               = 46
	SM_CXMINSPACING         = 47
	SM_CYMINSPACING         = 48
	SM_CXSMICON             = 49
	SM_CYSMICON             = 50
	SM_CYSMCAPTION          = 51
	SM_CXSMSIZE             = 52
	SM_CYSMSIZE             = 53
	SM_CXMENUSIZE           = 54
	SM_CYMENUSIZE           = 55
	SM_ARRANGE              = 56
	SM_CXMINIMIZED          = 57
	SM_CYMINIMIZED          = 58
	SM_CXMAXTRACK           = 59
	SM_CYMAXTRACK           = 60
	SM_CXMAXIMIZED          = 61
	SM_CYMAXIMIZED          = 62
	SM_NETWORK              = 63
	SM_CLEANBOOT            = 67
	SM_CXDRAG               = 68
	SM_CYDRAG               = 69
	SM_SHOWSOUNDS           = 70
	SM_CXMENUCHECK          = 71
	SM_CYMENUCHECK          = 72
	SM_SLOWMACHINE          = 73
	SM_MIDEASTENABLED       = 74
	SM_MOUSEWHEELPRESENT    = 75
	SM_XVIRTUALSCREEN       = 76
	SM_YVIRTUALSCREEN       = 77
	SM_CXVIRTUALSCREEN      = 78
	SM_CYVIRTUALSCREEN      = 79
	SM_CMONITORS            = 80
	SM_SAMEDISPLAYFORMAT    = 81
	SM_IMMENABLED           = 82
	SM_CXFOCUSBORDER        = 83
	SM_CYFOCUSBORDER        = 84
	SM_TABLETPC             = 86
	SM_MEDIACENTER          = 87
	SM_STARTER              = 88
	SM_SERVERR2             = 89
	SM_CMETRICS             = 91
	SM_REMOTESESSION        = 0x1000
	SM_SHUTTINGDOWN         = 0x2000
	SM_REMOTECONTROL        = 0x2001
	SM_CARETBLINKINGENABLED = 0x2002
)

//MSG ...
type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

const (
	GUI_CARETBLINKING  = 0x00000001
	GUI_INMENUMODE     = 0x00000004
	GUI_INMOVESIZE     = 0x00000002
	GUI_POPUPMENUMODE  = 0x00000010
	GUI_SYSTEMMENUMODE = 0x00000008
)

const (
	KBD_JAPANESE = "00000411"
)

//GUITHREADINFO
//cbSize: The size of this structure, in bytes. The caller must set this member to sizeof(GUITHREADINFO).
//flags: The thread state. This member can be one or more of the following values. e.g GUI_CARETBLINKING
//hwndActive: A handle to the active window within the thread.
//hwndFocus: A handle to the window that has the keyboard focus.
//hwndCapture: A handle to the window that has captured the mouse.
//hwndMenuOwner: A handle to the window that owns any active menus.
//hwndMoveSize: A handle to the window in a move or size loop.
//hwndCaret: A handle to the window that is displaying the caret.
//rect: The caret's bounding rectangle, in client coordinates, relative to the window specified by the hwndCaret member.
type GUITHREADINFO struct {
	cbSize        uint32 //DWORD
	flags         uint32 //DWORD
	HWNDActive    HWND
	HWNDFocus     HWND
	HWNDCapture   HWND
	HWNDMenuOwner HWND
	HWNDMoveSize  HWND
	HWNDCaret     HWND
	CRect         RECT
}

//MOUSE_INPUT ...
type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

//MOUSEINPUT ...
type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

//KEYBD_INPUT ...
type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

//KEYBDINPUT ...
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo uintptr
}

// KEYBDINPUT DwFlags
const (
	KEYEVENTF_KEYDOWN     = 0
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002
	KEYEVENTF_SCANCODE    = 0x0008
	KEYEVENTF_UNICODE     = 0x0004
)

//Virtual Keys
const (
	VK_SHIFT   = 0x10
	VK_CONTROL = 0x11
	VK_ALT     = 0x12
)

const (
	MSGFLT_RESET    = 0
	MSGFLT_ALLOW    = 1
	MSGFLT_DISALLOW = 2
)

const (
	WH_CALLWNDPROC     = 4
	WH_CALLWNDPROCRET  = 12
	WH_CBT             = 5
	WH_DEBUG           = 9
	WH_FOREGROUNDIDLE  = 11
	WH_GETMESSAGE      = 3
	WH_JOURNALPLAYBACK = 1
	WH_JOURNALRECORD   = 0
	WH_KEYBOARD        = 2
	WH_KEYBOARD_LL     = 13
	WH_MOUSE           = 7
	WH_MOUSE_LL        = 14
	WH_MSGFILTER       = -1
	WH_SHELL           = 10
	WH_SYSMSGFILTER    = 6
)

//HARDWARE_INPUT ...
type HARDWARE_INPUT struct {
	Type uint32
	Hi   HARDWAREINPUT
}

//HARDWAREINPUT ...
type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
	Unused  [16]byte
}

//POINT xy points in coordinate space
type POINT struct {
	X, Y int32
}

//RECT defines a rectangle
type RECT struct {
	Left, Top, Right, Bottom int32
}

type (
	//HANDLE ...
	HANDLE uintptr

	//HWND ...
	HWND      HANDLE
	HMODULE   HANDLE
	HHOOK     HANDLE
	HINSTANCE HANDLE
	HKL       HANDLE

	DWORD   uint32
	UINT    uint32
	BOOL    int32
	CHAR    rune
	LRESULT uintptr
	LPARAM  uintptr
	WPARAM  uintptr
)

//HOOKPROC
// type HOOKPROC func(int, uintptr, uintptr) uintptr

//HOOKPROC An application-defined callback (or hook) function that
//the system calls in response to events generated by an accessible object.
//hook: Handle to an event hook function. This value is returned
//			 by SetWinEventHook when the hook function is installed and
//			 is specific to each instance of the hook function.
//event: Specifies the event that occurred. This value is one of the event constants.
//hwnd: Handle to the window that generates the event, or NULL if no window is
//		associated with the event. For example, the mouse pointer is not associated with a window.
//idObject: Identifies the object associated with the event. This is one of the object
//			identifiers or a custom object ID.
//idChild: Identifies whether the event was triggered by an object or a child element of the object.
//idEventThread: ?????
//dwmsEventTime: Specifies the time, in milliseconds, that the event was generated.
type HOOKPROC func(hook HHOOK, event uint32, hwnd HWND, idObject, idChild, idEventThread, dwmsEventTime uint32) uintptr
type HOOKPROC3 func(code int, wparam uintptr, lparam uintptr) uintptr

//WNDENUMPROC ...
type WNDENUMPROC func(hwnd HWND, p uintptr) uintptr

//Events
const (
	EVENT_SYSTEM_FOREGROUND     = 0x0003
	EVENT_SYSTEM_MOVESIZESTART  = 0x000A
	EVENT_SYSTEM_MOVESIZEEND    = 0x000B
	EVENT_SYSTEM_SWITCHSTART    = 0x0014
	EVENT_SYSTEM_SWITCHEND      = 0x0015
	EVENT_OBJECT_FOCUS          = 0x8005
	EVENT_OBJECT_HIDE           = 0x8003
	EVENT_OBJECT_SELECTION      = 0x8006
	EVENT_OBJECT_LOCATIONCHANGE = 0x800B
	EVENT_MIN                   = 0x00000001
	EVENT_MAX                   = 0x7FFFFFFF
	EVENT_SYSTEM_DESKTOPSWITCH  = 0x0020

	OBJID_WINDOW   = 0x00000000
	OBJID_SYSMENU  = 0xFFFFFFFF
	OBJID_TITLEBAR = 0xFFFFFFFE
	OBJID_MENU     = 0xFFFFFFFD
	OBJID_CLIENT   = 0xFFFFFFFC
	OBJID_VSCROLL  = 0xFFFFFFFB
	OBJID_HSCROLL  = 0xFFFFFFFA
	OBJID_SIZEGRIP = 0xFFFFFFF9
	OBJID_CARET    = 0xFFFFFFF8
	OBJID_CURSOR   = 0xFFFFFFF7
	OBJID_ALERT    = 0xFFFFFFF6
	OBJID_SOUND    = 0xFFFFFFF5
)

const (
	WINEVENT_OUTOFCONTEXT   = 0x0
	WINEVENT_SKIPOWNTHREAD  = 0x1
	WINEVENT_SKIPOWNPROCESS = 0x2
	WINEVENT_INCONTEXT      = 0x4
)

const (
	WM_USER                   = 0x0400
	WM_APP                    = 0x8000
	WM_ACTIVATE               = 6
	WM_ACTIVATEAPP            = 28
	WM_AFXFIRST               = 864
	WM_AFXLAST                = 895
	WM_ASKCBFORMATNAME        = 780
	WM_CANCELJOURNAL          = 75
	WM_CANCELMODE             = 31
	WM_CAPTURECHANGED         = 533
	WM_CHANGECBCHAIN          = 781
	WM_CHAR                   = 258
	WM_CHARTOITEM             = 47
	WM_CHILDACTIVATE          = 34
	WM_CLEAR                  = 771
	WM_CLOSE                  = 16
	WM_COMMAND                = 273
	WM_COMMNOTIFY             = 68 /* OBSOLETE */
	WM_COMPACTING             = 65
	WM_COMPAREITEM            = 57
	WM_CONTEXTMENU            = 123
	WM_COPY                   = 769
	WM_COPYDATA               = 74
	WM_CREATE                 = 1
	WM_CTLCOLORBTN            = 309
	WM_CTLCOLORDLG            = 310
	WM_CTLCOLOREDIT           = 307
	WM_CTLCOLORLISTBOX        = 308
	WM_CTLCOLORMSGBOX         = 306
	WM_CTLCOLORSCROLLBAR      = 311
	WM_CTLCOLORSTATIC         = 312
	WM_CUT                    = 768
	WM_DEADCHAR               = 259
	WM_DELETEITEM             = 45
	WM_DESTROY                = 2
	WM_DESTROYCLIPBOARD       = 775
	WM_DEVICECHANGE           = 537
	WM_DEVMODECHANGE          = 27
	WM_DISPLAYCHANGE          = 126
	WM_DRAWCLIPBOARD          = 776
	WM_DRAWITEM               = 43
	WM_DROPFILES              = 563
	WM_ENABLE                 = 10
	WM_ENDSESSION             = 22
	WM_ENTERIDLE              = 289
	WM_ENTERMENULOOP          = 529
	WM_ENTERSIZEMOVE          = 561
	WM_ERASEBKGND             = 20
	WM_EXITMENULOOP           = 530
	WM_EXITSIZEMOVE           = 562
	WM_FONTCHANGE             = 29
	WM_GETDLGCODE             = 135
	WM_GETFONT                = 49
	WM_GETHOTKEY              = 51
	WM_GETICON                = 127
	WM_GETMINMAXINFO          = 36
	WM_GETTEXT                = 13
	WM_GETTEXTLENGTH          = 14
	WM_HANDHELDFIRST          = 856
	WM_HANDHELDLAST           = 863
	WM_HELP                   = 83
	WM_HOTKEY                 = 786
	WM_HSCROLL                = 276
	WM_HSCROLLCLIPBOARD       = 782
	WM_ICONERASEBKGND         = 39
	WM_INITDIALOG             = 272
	WM_INITMENU               = 278
	WM_INITMENUPOPUP          = 279
	WM_INPUT                  = 0X00FF
	WM_INPUTLANGCHANGE        = 81
	WM_INPUTLANGCHANGEREQUEST = 80
	WM_KEYDOWN                = 256
	WM_KEYUP                  = 257
	WM_KILLFOCUS              = 8
	WM_MDIACTIVATE            = 546
	WM_MDICASCADE             = 551
	WM_MDICREATE              = 544
	WM_MDIDESTROY             = 545
	WM_MDIGETACTIVE           = 553
	WM_MDIICONARRANGE         = 552
	WM_MDIMAXIMIZE            = 549
	WM_MDINEXT                = 548
	WM_MDIREFRESHMENU         = 564
	WM_MDIRESTORE             = 547
	WM_MDISETMENU             = 560
	WM_MDITILE                = 550
	WM_MEASUREITEM            = 44
	WM_GETOBJECT              = 0X003D
	WM_CHANGEUISTATE          = 0X0127
	WM_UPDATEUISTATE          = 0X0128
	WM_QUERYUISTATE           = 0X0129
	WM_UNINITMENUPOPUP        = 0X0125
	WM_MENURBUTTONUP          = 290
	WM_MENUCOMMAND            = 0X0126
	WM_MENUGETOBJECT          = 0X0124
	WM_MENUDRAG               = 0X0123
	WM_APPCOMMAND             = 0X0319
	WM_MENUCHAR               = 288
	WM_MENUSELECT             = 287
	WM_MOVE                   = 3
	WM_MOVING                 = 534
	WM_NCACTIVATE             = 134
	WM_NCCALCSIZE             = 131
	WM_NCCREATE               = 129
	WM_NCDESTROY              = 130
	WM_NCHITTEST              = 132
	WM_NCLBUTTONDBLCLK        = 163
	WM_NCLBUTTONDOWN          = 161
	WM_NCLBUTTONUP            = 162
	WM_NCMBUTTONDBLCLK        = 169
	WM_NCMBUTTONDOWN          = 167
	WM_NCMBUTTONUP            = 168
	WM_NCXBUTTONDOWN          = 171
	WM_NCXBUTTONUP            = 172
	WM_NCXBUTTONDBLCLK        = 173
	WM_NCMOUSEHOVER           = 0X02A0
	WM_NCMOUSELEAVE           = 0X02A2
	WM_NCMOUSEMOVE            = 160
	WM_NCPAINT                = 133
	WM_NCRBUTTONDBLCLK        = 166
	WM_NCRBUTTONDOWN          = 164
	WM_NCRBUTTONUP            = 165
	WM_NEXTDLGCTL             = 40
	WM_NEXTMENU               = 531
	WM_NOTIFY                 = 78
	WM_NOTIFYFORMAT           = 85
	WM_NULL                   = 0
	WM_PAINT                  = 15
	WM_PAINTCLIPBOARD         = 777
	WM_PAINTICON              = 38
	WM_PALETTECHANGED         = 785
	WM_PALETTEISCHANGING      = 784
	WM_PARENTNOTIFY           = 528
	WM_PASTE                  = 770
	WM_PENWINFIRST            = 896
	WM_PENWINLAST             = 911
	WM_POWER                  = 72
	WM_POWERBROADCAST         = 536
	WM_PRINT                  = 791
	WM_PRINTCLIENT            = 792
	WM_QUERYDRAGICON          = 55
	WM_QUERYENDSESSION        = 17
	WM_QUERYNEWPALETTE        = 783
	WM_QUERYOPEN              = 19
	WM_QUEUESYNC              = 35
	WM_QUIT                   = 18
	WM_RENDERALLFORMATS       = 774
	WM_RENDERFORMAT           = 773
	WM_SETCURSOR              = 32
	WM_SETFOCUS               = 7
	WM_SETFONT                = 48
	WM_SETHOTKEY              = 50
	WM_SETICON                = 128
	WM_SETREDRAW              = 11
	WM_SETTEXT                = 12
	WM_SETTINGCHANGE          = 26
	WM_SHOWWINDOW             = 24
	WM_SIZE                   = 5
	WM_SIZECLIPBOARD          = 779
	WM_SIZING                 = 532
	WM_SPOOLERSTATUS          = 42
	WM_STYLECHANGED           = 125
	WM_STYLECHANGING          = 124
	WM_SYSCHAR                = 262
	WM_SYSCOLORCHANGE         = 21
	WM_SYSCOMMAND             = 274
	WM_SYSDEADCHAR            = 263
	WM_SYSKEYDOWN             = 260
	WM_SYSKEYUP               = 261
	WM_TCARD                  = 82
	WM_THEMECHANGED           = 794
	WM_TIMECHANGE             = 30
	WM_TIMER                  = 275
	WM_UNDO                   = 772
	WM_USERCHANGED            = 84
	WM_VKEYTOITEM             = 46
	WM_VSCROLL                = 277
	WM_VSCROLLCLIPBOARD       = 778
	WM_WINDOWPOSCHANGED       = 71
	WM_WINDOWPOSCHANGING      = 70
	WM_WININICHANGE           = 26
	WM_KEYFIRST               = 256
	WM_KEYLAST                = 264
	WM_SYNCPAINT              = 136
	WM_MOUSEACTIVATE          = 33
	WM_MOUSEMOVE              = 512
	WM_LBUTTONDOWN            = 513
	WM_LBUTTONUP              = 514
	WM_LBUTTONDBLCLK          = 515
	WM_RBUTTONDOWN            = 516
	WM_RBUTTONUP              = 517
	WM_RBUTTONDBLCLK          = 518
	WM_MBUTTONDOWN            = 519
	WM_MBUTTONUP              = 520
	WM_MBUTTONDBLCLK          = 521
	WM_MOUSEWHEEL             = 522
	WM_XBUTTONDOWN            = 523
	WM_XBUTTONUP              = 524
	WM_XBUTTONDBLCLK          = 525
	WM_MOUSEHWHEEL            = 526
	WM_MOUSEFIRST             = 512
	WM_MOUSELAST              = 526
	WM_MOUSEHOVER             = 0X2A1
	WM_MOUSELEAVE             = 0X2A3
	WM_CLIPBOARDUPDATE        = 0x031D
)

//MustLoadLibrary loads a dll and then returns error if
//not able to
func MustLoadLibrary(name string) (uintptr, error) {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		return 0, err
	}

	return uintptr(lib), nil
}

//MustGetProcAddress ...
func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}

//UintToString converts []uint16 to a string
func UintToString(in []uint16) string {
	return syscall.UTF16ToString(in)
}

//StringToUint converts []uint16 to a string
func StringToUint(in string) *uint16 {
	return syscall.StringToUTF16Ptr(in)
}

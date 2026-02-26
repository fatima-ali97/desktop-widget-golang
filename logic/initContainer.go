package logic

import (
	"log"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyWidget struct {
	walk.WidgetBase
}

var myWidgetWindowClass = "MyWidget Class"

func NewMyWidget(parent walk.Container) (*MyWidget, error) {
	w := new(MyWidget)

	if err := walk.InitWidget(
		w,
		parent,
		myWidgetWindowClass,
		win.WS_VISIBLE,
		0); err != nil {

		return nil, err
	}

	bg, err := walk.NewSolidColorBrush(walk.RGB(67,120,245))
	if err != nil {
		return nil, err
	}
	w.SetBackground(bg)

	return w, nil
}

func (*MyWidget) CreateLayoutItem(ctx *walk.LayoutContext) walk.LayoutItem {
	return &myWidgetLayoutItem{idealSize: walk.SizeFrom96DPI(walk.Size{50, 50}, ctx.DPI())}
}

type myWidgetLayoutItem struct {
	walk.LayoutItemBase
	idealSize walk.Size // in native pixels
}

func (li *myWidgetLayoutItem) LayoutFlags() walk.LayoutFlags {
	return 0
}

func (li *myWidgetLayoutItem) IdealSize() walk.Size {
	return li.idealSize
}

func (w *MyWidget) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", w.Name())
	}

	return w.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

type MyPushButton struct {
	*walk.PushButton
}

func NewMyPushButton(parent walk.Container) (*MyPushButton, error) {
	pb, err := walk.NewPushButton(parent)
	if err != nil {
		return nil, err
	}

	mpb := &MyPushButton{pb}

	if err := walk.InitWrapperWindow(mpb); err != nil {
		return nil, err
	}

	return mpb, nil
}

func (mpb *MyPushButton) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", mpb.Text())
	}

	return mpb.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

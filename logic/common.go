package logic

import "github.com/lxn/walk"

var Mw *walk.MainWindow

func SetParent(mw *walk.MainWindow) {
	Mw = mw
}

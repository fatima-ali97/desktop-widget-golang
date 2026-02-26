package main

import (
	"log"
	"main/logic"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

const myWidgetWindowClass = "MyWidget Class"
var mw *walk.MainWindow
func init() {
	walk.AppendToWalkInit(func() {
		walk.MustRegisterWindowClass(myWidgetWindowClass)
	})
	logic.SetParent(mw)
}

var buttons = []string{"Pomodoro", "Short Break", "Long Break"}

func main() {
	

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "Timer",
		Size:     Size{400, 300},
		Layout:   HBox{},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	for _, name := range buttons {
		if w, err := logic.NewMyWidget(mw); err != nil {
			log.Fatal(err)
		} else {
			w.SetName(name)
		}
	}

	mpb, err := logic.NewMyPushButton(mw)
	if err != nil {
		log.Fatal(err)
	}
	mpb.SetText("MyPushButton")

	mw.Run()
}

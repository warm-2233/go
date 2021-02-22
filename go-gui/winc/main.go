package main

import (
	"github.com/tadvi/winc"
)

func main() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(750, 750) // (width, height)
	mainWindow.SetText("Hello World Demo")

	// mainWindow.SetIcon(1)

	edt := winc.NewEdit(mainWindow)
	edt.SetPos(10, 20)
	// Most Controls have default size unless SetSize is called.
	edt.SetText("edit text")

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Show or Hide")
	btn.SetPos(40, 50)   // (x, y)
	btn.SetSize(100, 40) // (width, height)
	btn.OnClick().Bind(func(e *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})
	mainWindow.Center()
	mainWindow.Show()
	// mainWindow.OnClose().Bind(wndOnClose)
	mainWindow.OnClose().Bind(func(arg *winc.Event) {
		winc.Exit()
	})

	img := winc.NewImageView(mainWindow)
	img.SetPos(0, 0)
	img.DrawImageFile("./i.jpg")
	// img.SetSize(mainWindow.Width(), mainWindow.Height())

	winc.RunMainLoop() // Must call to start event loop.
}

// func wndOnClose(arg *winc.Event) {
// 	winc.Exit()
// }

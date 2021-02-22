// package main

// import (
//     "github.com/andlabs/ui"
//     _ "github.com/andlabs/ui/winmanifest"
// )

// func main() {
//     err := ui.Main(func() {
//         // 生成：文本框
//         name := ui.NewEntry()
//         // 生成：标签
//         greeting := ui.NewLabel(``)
//         // 生成：按钮
//         button := ui.NewButton(`欢迎`)
//         // 设置：按钮点击事件
//         button.OnClicked(func(*ui.Button) {
//             greeting.SetText(`你好，` + name.Text() + `！`)
//         })
//         // 生成：垂直容器
//         box := ui.NewVerticalBox()

//         // 往 垂直容器 中添加 控件
//         box.Append(ui.NewLabel(`请输入你的名字：`), false)
//         box.Append(name, false)
//         box.Append(button, false)
//         box.Append(greeting, false)

//         // 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
//         window := ui.NewWindow(`你好`, 200, 100, false)

//         // 窗口容器绑定
//         window.SetChild(box)

//         // 设置：窗口关闭时
//         window.OnClosing(func(*ui.Window) bool {
//             // 窗体关闭
//             ui.Quit()
//             return true
//         })

//         // 窗体显示
//         window.Show()
//     })
//     if err != nil {
//         panic(err)
//     }
// }

package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appId = "com.nayoso.example"
	//每个gtk3程序都需要一步
	app, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	//为activate事件绑定函数, activate会在程序启动时触发，也就是app.Run()时
	app.Connect("activate", func() {
		onActivate(app)
	})

	app.Run(os.Args) //运行gtkApplication
}

func onActivate(application *gtk.Application) {
	appWindow, err := gtk.ApplicationWindowNew(application) //创建window控件

	if err != nil {
		log.Fatal("Could not create application window.", err)
	}
	//设置窗口属性
	appWindow.SetTitle("Basic Application.")
	appWindow.SetDefaultSize(400, 400)
	//显示窗口
	appWindow.Show()
}

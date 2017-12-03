package main

import (
	"os"

	"github.com/zurek87/go-gtk3/example/builder/callback"
	"github.com/zurek87/go-gtk3/gtk"
)

//"github.com/zurek87/go-gtk3/example/builder/callback"
func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()

	builder.AddFromFile("hello.ui")
	obj := builder.GetObject("window1")

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk.MainQuit)

	callback.Init(builder)

	gtk.Main()
}

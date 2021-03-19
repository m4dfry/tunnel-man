package main

import (
	"github.com/webview/webview"
)

func main() {
	go func() {
		r := NewRoutes()
		r.Run("localhost:8090")
	}()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Go Webview Test")
	w.SetSize(800, 600, webview.HintNone)

	// Create a GoLang function callable from JS
	w.Bind("hello", func() string { return "Welcome to my World!" })

	w.Navigate("http://localhost:8090/")

	w.Run()
}

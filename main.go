package main

import (
	"encoding/json"
	"os"
	"syscall/js"

	"github.com/vqvw/runic"
)

func main() {
	parser := runic.New()

	var runicCallback js.Func
	runicCallback = js.FuncOf(func(this js.Value, args []js.Value) any {
		editorData, err := parser.EditorData(args[0].String())
		if err != nil {
			return ""
		}

		editorDataJson, err := json.Marshal(editorData)
		if err != nil {
			return ""
		}

		return string(editorDataJson)
	})

	var unmountCallback js.Func
	unmountCallback = js.FuncOf(func(this js.Value, args []js.Value) any {
		runicCallback.Release()
		unmountCallback.Release()
		os.Exit(0)
		return nil
	})

	js.Global().Set("runicParse", runicCallback)
	js.Global().Set("runicUnmount", unmountCallback)

	// Keep program running until exited by `unmountCallback`
	select {}
}

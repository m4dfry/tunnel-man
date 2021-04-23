package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/webview/webview"
)

func main() {
	tunnelDef := readConf()

	go func() {
		r := NewRoutes(tunnelDef)
		r.Run("localhost:8090")
	}()

	startInterface()
	/*for {
		fmt.Println("Infinite loop for BE test purpose")
		time.Sleep(time.Second)
	}*/

}

func readConf() TunnelsMap {
	// open json file
	jsonFile, err := os.Open("default.config.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of file
	defer jsonFile.Close()

	// read file as byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize & unmarshal config
	var conf ConfigurationFile
	json.Unmarshal(byteValue, &conf)

	// type TunnelsMap map[string]Tunnel
	var ret = make(TunnelsMap)
	var certs = make(map[string][]string)

	for _, c := range conf.Certificates {
		certs[c.Name] = c.Files
	}

	for _, t := range conf.Tunnels {
		ret[t.Name] = &Tunnel{
			t.Bastion,
			t.Address,
			t.Localport,
			certs[t.Certificate],
		}
	}

	return ret
}

func startInterface() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Tunnel Man")
	w.SetSize(800, 600, webview.HintNone)

	// Create a GoLang function callable from JS
	w.Bind("hello", func() string { return "Welcome to my World!" })

	w.Navigate("http://localhost:8090/")

	w.Run()
}

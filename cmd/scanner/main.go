package main

import (
	"fmt"
	"github.com/zaouldyeck/plugin-core/scanner"
	"log"
	"os"
	"plugin"
)

const PluginsDir = "../../plugins/"

func main() {
	var (
		files []os.DirEntry
		err   error
		p     *plugin.Plugin
		n     plugin.Symbol
		check scanner.Checker
		res   *scanner.Result
	)
	if files, err = os.ReadDir(PluginsDir); err != nil {
		log.Fatalln(err)
	}

	for idx := range files {
		fmt.Println("Found plugin: " + files[idx].Name())
		if p, err = plugin.Open(PluginsDir + "/" + files[idx].Name()); err != nil {
			log.Fatalln(err)
		}

		if n, err = p.Lookup("New"); err != nil {
			log.Fatalln(err)
		}

		newFunc, ok := n.(func() scanner.Checker)
		if !ok {
			log.Fatalln("Plugin entry point is no good. Expecting: func New() scanner.Checker{ ... }")
		}
		check = newFunc()
		res = check.Check("192.168.100.100", 8080)
		if res.Vulnerable {
			log.Println("Host is vulnerable: " + res.Details)
		} else {
			log.Println("Host is NOT vulnerable")
		}

	}
}

package helper

import "fmt"

func PrintHelp() {
	fmt.Println(`Marketflow
	
Usage:
  marketflow [--port <N>]
  marketflow --help

Options:
  --port N     Port number`)
}

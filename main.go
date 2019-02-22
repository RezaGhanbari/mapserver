package main

import (
	"log"
	"mapserver/app"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
	application := app.App{}
	application.Initialize()
	application.Run("3000")
}

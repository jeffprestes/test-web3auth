package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(macaron.Static("static"))
	m.Run(5000)
}

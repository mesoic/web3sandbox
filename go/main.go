package main

import (
	fn "./func"
	"./msg"
)

func main() {
	m := msg.Msg{}
	m.Message("-- web3sandbox --")

	fn.Run()
}

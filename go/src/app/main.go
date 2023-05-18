package main

import (
	"github.com/mesoic/web3sandbox/go/src/msg"
	"github.com/mesoic/web3sandbox/go/src/run"
)

func main() {
	m := msg.Msg{}
	m.Message("-- web3sandbox --")

	run.Run()
}

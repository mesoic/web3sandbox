package msg

import "fmt"

type Msg struct{}

func (*Msg) Warning(msg string) {

	fmt.Printf("%s", msg)
}

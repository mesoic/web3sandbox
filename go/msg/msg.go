package msg

import "fmt"

type Msg struct{}

func (*Msg) Verified(msg string) {
	fmt.Printf("  verified: %s\n", msg)
}

func (*Msg) Warning(msg string) {
	fmt.Printf("  warning: %s\n", msg)
}

func (*Msg) Error(msg string) {
	fmt.Printf("  error: %s\n", msg)
}

func (*Msg) Critical(msg string) {
	fmt.Printf("  critical: %s\n", msg)
}

func (*Msg) Info(msg string) {
	fmt.Printf("  info: %s\n", msg)
}

func (*Msg) Message(msg string) {
	fmt.Printf("  %s\n", msg)
}

func (*Msg) Log(msg string) {
	fmt.Printf("  > %s\n", msg)
}

func (*Msg) Newline(msg string) {
	fmt.Printf("\n")
}

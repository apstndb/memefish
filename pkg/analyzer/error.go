package analyzer

import (
	"bytes"
	"fmt"

	"github.com/cloudspannerecosystem/memefish/pkg/token"
)

type Error struct {
	Message  string
	Position *token.Position // optional
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) Error() string {
	if e.Position == nil {
		return fmt.Sprintf("syntax error: %s", e.Message)
	}

	var message bytes.Buffer
	fmt.Fprintf(&message, "analyze error:%s: %s\n", e.Position, e.Message)
	if e.Position.Source != "" {
		fmt.Fprintln(&message)
		fmt.Fprintln(&message, e.Position.Source)
	}
	return message.String()
}

package human

import "fmt"

// Head of human
type Head struct {
	Name string
}

// Greet ing
func (h Head) Greet() string {
	var message string
	if h.Name != "" {
		message = fmt.Sprintf("hello, my name is %s", h.Name)
	} else {
		message = "hello, nice to meet you"
	}
	return message
}

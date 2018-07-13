package kubeless

import (
	"github.com/kubeless/kubeless/pkg/functions"
	"encoding/json"
)

type Message struct {
	Name string
}

// Foo sample function
func DoStuff(event functions.Event, context functions.Context) (string, error) {
	var input Message
	if event.Data == "" {
		return "You didn't provide a name!", nil
	}
	err := json.Unmarshal([]byte(event.Data), &input)
	if err != nil {
		return "You did not provide valid input.  Please provide JSON that has \"Name\" field only", nil
	}
	return "Hello " + input.Name + "! This is a sample function, Mike!", nil
}

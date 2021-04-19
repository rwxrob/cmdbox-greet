package greet

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("greet")
	//x.Default = "eng"
	x.Summary = "greet others in different tongues"

	x.Method = func() error {
		fmt.Println("would greet")
		fmt.Println(x.Summary)
		return nil
	}

	//x.Add(english)
	//x.Add(french)
}

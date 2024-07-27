package gite

import "fmt"

func reportRegisteredHandlers(path string, hanlders int) {
	fmt.Printf("%v => hanlders: %v\n", path, hanlders)
}

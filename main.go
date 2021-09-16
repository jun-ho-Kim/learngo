package main

import (
	"fmt"

	"github.com/junhoKim/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)

	if err != nil {
		fmt.Println(err)
	}
	errUpdate := dictionary.Update(word, "hihi")
	dictionary.Delete(word)
	hello, _ := dictionary.Search(word)
	fmt.Println("errUpdate", errUpdate)
	fmt.Println("found", word, "definition", hello)
}

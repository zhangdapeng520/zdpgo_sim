package quick_test

import (
	"log"
	"os"

	"github.com/zhangdapeng520/zdpgo_sim/chroma/quick"
)

func Example() {
	code := `package main

func main() { }
`
	err := quick.Highlight(os.Stdout, code, "go", "html", "monokai")
	if err != nil {
		log.Fatal(err)
	}
}

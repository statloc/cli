package main

import (
	"fmt"
	"os"

	"cli/internal"
)

func main () {
    if len(os.Args) > 2 {
        fmt.Println("Error parsing argument: path specified incorrectly")
    } else {
        path := "."
        if len(os.Args) == 2 {
            path = os.Args[1]
        }
        fmt.Print(internal.Respond(path))
    }
}

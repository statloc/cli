package main

import (
	"fmt"
	"os"

	core "github.com/statloc/core"

	"cli/internal"
)

func main () {
    if len(os.Args) != 2 {
        fmt.Println("Error parsing argument: path specified incorrectly")
    } else {
        path := os.Args[1]
        response, err := core.GetStatistics(path)

        if err != nil {
            fmt.Printf("Path %s is not found\n", path)
        }

        fmt.Print(internal.GetTable(response.Items))
    }
}

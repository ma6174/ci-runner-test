package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	c := color.New(color.FgCyan, color.Bold)
	c.Println("Hello from ci-runner-test!")
	fmt.Println("Go module cache V2 smoke test.")
}

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"dummy",
}

func main() {
	fmt.Println("dummy")
}

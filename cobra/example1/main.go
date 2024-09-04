package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
	}

	subCmd := &cobra.Command{
		Use:   "print-command",
		Short: "print full command line string",
		Run: func(cmd *cobra.Command, args []string) {
			fullCommand := cmd.Use + " " + strings.Join(args, " ")
			fmt.Println("Full command line string: ", fullCommand)
		},
	}

	rootCmd.AddCommand(subCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

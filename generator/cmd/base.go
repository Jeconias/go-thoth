package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var baseCMD = &cobra.Command{
	Use:   "thoth",
	Short: "Thoth is simple lib for validate structs",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := baseCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

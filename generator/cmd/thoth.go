package cmd

import (
	"fmt"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/spf13/cobra"
)

var (
	directory string
)

var thoth = &cobra.Command{
	Use:   "gen",
	Short: "This command parse your package and generate the files necessary",
	Run: func(cmd *cobra.Command, args []string) {
		flag := cmd.Flag("d")
		fmt.Println(flag.Value)

		_, err := myasthurts.NewEnvironment()
		if err != nil {
			panic("error") //TODO(Jeconias): add pattern
		}

	},
}

func init() {
	baseCMD.AddCommand(thoth)

	//thoth.PersistentFlags().String("d", "adf", "Directory to parse")
	thoth.Flags().String("d", "", "Use this flag to set directory")
	thoth.MarkFlagRequired("d")
}

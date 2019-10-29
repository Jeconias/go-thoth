package cmd

import (
	"fmt"

	template "github.com/jeconias/go-thoth/generator/template"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/spf13/cobra"
)

var (
	directory string
)

var thoth = &cobra.Command{
	Use:   "gen",
	Short: "This command parse your package and generate the files necessary",
	RunE: func(cmd *cobra.Command, args []string) error {
		flag := cmd.Flag("d")

		env, err := myasthurts.NewEnvironment()
		if err != nil {
			return err
		}

		pkg, err := env.ParseDir(flag.Value.String())
		if err != nil {
			return err
		}

		fmt.Println(template.Struct(pkg))
		return nil
	},
}

func init() {
	baseCMD.AddCommand(thoth)

	thoth.Flags().String("d", "", "Use this flag to set directory")
	thoth.MarkFlagRequired("d")
}

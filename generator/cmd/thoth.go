package cmd

import (
	"github.com/spf13/cobra"
)

var (
	directory string
)

var genCommand = &cobra.Command{
	Use:   "gen",
	Short: "This command parse your package and generate the files necessary",
	//TODO(Jeconias): Add long description
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := gen(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

var genTestCommand = &cobra.Command{
	Use:   "genTest",
	Short: "This command parse your package and generate the files necessary for test",
	//TODO(Jeconias): Add long description
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := gen(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

var genTemplate = &cobra.Command{
	Use:   "tpl",
	Short: "This command build all templates",
	//TODO(Jeconias): Add long description
	RunE: func(cmd *cobra.Command, args []string) error {
		//TODO(Jeconias): Add command to build template
		return nil
	},
}

func init() {
	baseCMD.AddCommand(genCommand)
	baseCMD.AddCommand(genTestCommand)

	genCommand.Flags().String("d", ".", "Use this flag to set directory for parse")
	genCommand.Flags().String("n", "thothGen", "Use this flag to set name of file")
	genCommand.Flags().String("s", ".", "Use this flag to set directory for save")
	// This is temp for tests
	genTestCommand.Flags().String("d", "tests/parse_dir", "Temp")
	genTestCommand.Flags().String("n", "thothGenForTest", "Temp")
	genTestCommand.Flags().String("s", "./tests/gen", "Temp")
}

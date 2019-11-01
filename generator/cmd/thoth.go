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

		if cmd.Flag("t").Value.String() == "generate" {
			cmd.Flag("d").Value.Set("test/parse_dir")
			cmd.Flag("n").Value.Set("generated")
			cmd.Flag("s").Value.Set("./test/parse_dir")
		}

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

	genCommand.Flags().String("d", ".", "Use this flag to set directory for parse")
	genCommand.Flags().String("n", defaultNameFileGenerated, "Use this flag to set name of file")
	genCommand.Flags().String("s", ".", "Use this flag to set directory for save")
	genCommand.Flags().String("t", "", "Use this flag to generate files to test")
}

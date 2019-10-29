package cmd

import (
	"fmt"
	"os"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	template "github.com/lab259/go-thoth/generator/template"
	"github.com/spf13/cobra"
)

func gen(cmd *cobra.Command, args []string) error {
	structsThoth := make([]*myasthurts.Struct, 0)

	flagDirectory := cmd.Flag("d").Value.String()
	flagNameFile := cmd.Flag("n").Value.String()
	flagDirectoryToSave := cmd.Flag("s").Value.String()

	if flagDirectory == "" {
		flagDirectory = "."
	}
	if flagNameFile == "" {
		flagNameFile = "thothGen"
	}

	if flagDirectoryToSave == "." {
		flagDirectoryToSave = "./generator/gen"
	}

	env, err := myasthurts.NewEnvironment()
	if err != nil {
		return err
	}

	pkg, err := env.ParseDir(flagDirectory)
	if err != nil {
		return err
	}

	for _, s := range pkg.Structs {
		for _, f := range s.Fields {
			if isThoth := f.Tag.TagParamByName("thoth"); isThoth != nil {
				structsThoth = append(structsThoth, s)
				break
			}
		}
	}

	if err := os.Chdir(flagDirectoryToSave); err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s.go", flagNameFile), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	template.RenderStruct(file, pkg, structsThoth)
	return nil
}

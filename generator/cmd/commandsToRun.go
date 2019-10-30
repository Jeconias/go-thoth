package cmd

import (
	"fmt"
	"os"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	templates "github.com/lab259/go-thoth/generator/templates"
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
		flagNameFile = defaultNameFileGenerated
	}

	if flagDirectoryToSave == "." {
		flagDirectoryToSave = flagDirectory
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
			if isThoth := f.Tag.TagParamByName(defaultNameTag); isThoth != nil {
				structsThoth = append(structsThoth, s)
				break
			}
		}
	}

	if err := os.Chdir(flagDirectoryToSave); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.go", flagNameFile)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	templates.RenderStruct(file, fileName, pkg, structsThoth)
	return nil
}

package internal

import (
	"fmt"

	"cuelang.org/go/cue"
	"github.com/spf13/cobra"
)

type docOptions struct {
	schemaPath   string
	outputPath   string
	templatePath string
}

var defaultOutputPath = "/docs/output.md"

func (o *docOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.outputPath, "out", "o", defaultOutputPath, "path to write the output file, defaults to /docs/output.md",
	)

	cmd.PersistentFlags().StringVarP(
		&o.templatePath, "template", "t", "template.md",
		"path to the checklist template file",
	)

	cmd.PersistentFlags().StringVarP(
		&o.schemaPath, "schema", "s", "../schema.cue",
		"path to the cue schema",
	)
}

func AddDocs(parentCmd *cobra.Command, cueContext *cue.Context) {
	opts := docOptions{}
	compileCmd := &cobra.Command{
		Use:           "docs",
		Short:         "Generate the documentation for Insights",
		SilenceUsage:  false,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			generator := NewDocsGenerator(opts, cueContext)
			contents := generator.Generate()
			fmt.Print(contents)
			return nil
		},
	}
	opts.AddFlags(compileCmd)
	parentCmd.AddCommand(compileCmd)
}

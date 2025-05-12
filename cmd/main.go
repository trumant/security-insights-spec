package main

import (
	"cuelang.org/go/cue/cuecontext"
	"github.com/ossf/security-insights-spec/internal"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "si",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cueContext := cuecontext.New()
	// Add the subcommands
	internal.AddDocs(rootCmd, cueContext)

	err := rootCmd.Execute()
	if err != nil {
		// Handle the error
		panic(err)
	}
}

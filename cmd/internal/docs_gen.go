package internal

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

type DocsGenerator struct {
	options    docOptions
	cueContext *cue.Context
}

func NewDocsGenerator(options docOptions, cueContext *cue.Context) *DocsGenerator {
	return &DocsGenerator{
		options:    options,
		cueContext: cueContext,
	}
}

func (d *DocsGenerator) Generate() error {
	args := []string{d.options.schemaPath}
	// Load the schema
	insts := load.Instances(args, &load.Config{})
	for _, inst := range insts {
		instVal := d.cueContext.BuildInstance(inst)
		if err := instVal.Validate(); err != nil {
			return err
		}
		iter, err := instVal.Fields(cue.Definitions(true))
		if err != nil {
			return err
		}

		for iter.Next() {
			sel := iter.Selector()
			if sel.IsDefinition() {
				doc, _ := d.generateMarkdown(sel, instVal)
				fmt.Println("Doc line: ", doc)
			}
		}
	}
	return nil
}

func (d *DocsGenerator) generateMarkdown(definition cue.Selector, value cue.Value) (string, error) {
	path := cue.MakePath(definition)
	definitionName := strings.Split(definition.String(), "#")[1]

	val := value.LookupPath(path)
	if val.Err() != nil {
		return "", val.Err()
	}
	definitionComment := ""
	doc := val.Doc()
	for _, d := range doc {
		for _, line := range d.List {
			definitionComment += fmt.Sprintf("  - %s\n", strings.Split(line.Text, "// ")[1])
		}
	}
	// definitionFields, err := value.Fields()
	// if err != nil {
	// 	return "", err
	// }
	return fmt.Sprintf("%s\n\n%s", definitionName, definitionComment), nil
}

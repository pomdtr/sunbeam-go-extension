package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pomdtr/sunbeam/pkg/types"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if len(os.Args) == 1 {
		encoder.Encode(types.Manifest{
			Title: "Sunbeam",
			Commands: []types.CommandSpec{
				{
					Name:  "hello",
					Title: "Hello",
					Mode:  types.CommandModeView,
					Params: []types.Param{
						{
							Name: "name",
							Type: types.ParamTypeString,
						},
					},
				},
			},
		})
		os.Exit(0)
	}

	switch os.Args[1] {
	case "hello":
		var input struct {
			Params struct {
				Name string
			}
		}

		if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		name := input.Params.Name
		if name == "" {
			name = "World"
		}

		page := types.Detail{
			Title:    "Hello",
			Markdown: fmt.Sprintf("> Hello, %s!", name),
			Actions: []types.Action{
				{
					Title: "Copy Text",
					OnAction: types.Command{
						Type: types.CommandTypeCopy,
						Text: fmt.Sprintf("Hello, %s!", name),
						Exit: true,
					},
				},
			},
		}

		encoder.Encode(page)
	}
}

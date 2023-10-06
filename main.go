package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
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
							Name:     "name",
							Type:     types.ParamTypeString,
							Optional: true,
						},
					},
				},
			},
		})
		os.Exit(0)
	}

	var input types.CommandInput
	if err := json.NewDecoder(strings.NewReader(os.Args[1])).Decode(&input); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch input.Command {
	case "hello":
		var params struct {
			Name string
		}

		if err := mapstructure.Decode(input.Params, &params); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if params.Name == "" {
			params.Name = "World"
		}

		page := types.Detail{
			Title: "Hello",
			Text:  fmt.Sprintf("Hello, %s!", params.Name),
			Actions: []types.Action{
				{
					Title: "Copy Text",
					OnAction: types.Command{
						Type: types.CommandTypeCopy,
						Text: fmt.Sprintf("Hello, %s!", params.Name),
						Exit: true,
					},
				},
			},
		}

		encoder.Encode(page)
	}
}

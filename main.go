package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type ValueAnnotator struct {
	Variables map[string]interface{} `yaml:"variables" json:"variables"`
	Templates map[string]string      `yaml:"templates" json:"templates"`
}

func main() {
	config := new(ValueAnnotator)

	fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		for i := range items {
			str, err := items[i].String()
			if err != nil {
				return nil, err
			}

			for key, value := range config.Variables {
				var valueStr = fmt.Sprint(value)
				if templateStr, ok := config.Templates[key]; ok {
					tmpl, err := template.New(key).Parse(templateStr)
					if err != nil {
						return nil, err
					}

					var tmplResult bytes.Buffer
					if err := tmpl.Execute(&tmplResult, value); err != nil {
						return nil, err
					}

					valueStr = tmplResult.String()
				}

				str = strings.ReplaceAll(str, "$"+key, valueStr)
			}

			items[i], err = yaml.Parse(str)
			if err != nil {
				return nil, err
			}
		}
		return items, nil
	}
	p := framework.SimpleProcessor{Config: config, Filter: kio.FilterFunc(fn)}
	if err := command.Build(p, command.StandaloneDisabled, false).Execute(); err != nil {
		os.Exit(1)
	}
}

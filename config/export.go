package config

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func ExportToFile(cfg Config, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer f.Close()

	err = ExportToWriter(cfg, f)
	if err != nil {
		return err
	}

	return nil
}

func ExportToWriter(cfg Config, w io.Writer) error {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)

	err := enc.Encode(cfg)
	if err != nil {
		return fmt.Errorf("encode yaml: %w", err)
	}

	err = envconfig.Usaget("APP", &cfg, w, template.Must(template.New("cfg").Parse(envTemplate)))
	if err != nil {
		return fmt.Errorf("encode env usage: %w", err)
	}

	return nil
}

const envTemplate = `
{{ range . }}# {{ .Key }}={{ .Field }}
{{ end }}`

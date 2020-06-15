package cli

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"
)

type Renderer struct {
	TemplatePath string
}

func (r *Renderer) render(data interface{}) (string, error) {
	t := template.Must(template.New(path.Base(r.TemplatePath)).Funcs(template.FuncMap{"split": strings.Split}).ParseFiles(r.TemplatePath))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("couldn't parse template for %+v: %v", data, err)
	}

	return buf.String(), nil
}

func NewRenderer(templatePath string) Renderer {
	return Renderer{TemplatePath: templatePath}
}

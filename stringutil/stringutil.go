package stringutil

import (
	"bytes"
	"html/template"
)

// parse string template
func ParseString(str string, vars interface{}) (string, error) {
	tmpl, err := template.New("tmpl").Parse(str)
	if err != nil {
		return "", err
	}
	return parse(tmpl, vars)
}

func parse(t *template.Template, vars interface{}) (string, error) {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

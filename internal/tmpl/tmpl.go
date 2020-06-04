package tmpl

import (
	"bytes"
	"text/template"
)

// Template holds data that can be applied to a template string.
type Template struct {
	fields Fields
}

// Fields that will be available to the template engine.
type Fields map[string]interface{}

// New creates a new template
func New() *Template {
	return &Template{
		fields: Fields{},
	}
}

// WithFields creates a new template with fields
func (t *Template) WithFields(f Fields) *Template {
	for k, v := range f {
		t.fields[k] = v
	}
	return t
}

// Apply template
func (t *Template) Apply(s string) (string, error) {
	var out bytes.Buffer
	tmpl, err := template.New("tmpl").Option("missingkey=error").Parse(s)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&out, t.fields)
	return out.String(), err
}

package fmt

import (
	"bytes"
	"html/template"
	"testing"
)

type fmtTest struct {
	args        []interface{}
	expected    interface{}
	expectedErr error
}

func TestModuleFunc(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"fmt": ModuleFunc,
	}).Parse(`{{ "Hello\nWorld" | fmt.Indent "  " }}`)
	if err != nil {
		t.Error(err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		t.Error(err)
	}
	if got := buf.String(); got != "  Hello\n  World" {
		t.Errorf("unexpected executed template content: %s", got)
	}
}

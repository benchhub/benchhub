package generator

import (
	"bytes"
	"io"
	"text/template"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/genutil"
)

// util.go contains wrapper around text/template

func renderTo(name string, dst io.Writer, tmpl string, data interface{}) error {
	tpl, err := template.New(name).Parse(tmpl)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return errors.Wrap(err, "error render template")
	}
	formatted, err := genutil.Format(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "error format generated code")
	}
	if _, err := dst.Write(formatted); err != nil {
		return errors.Wrap(err, "error write generated code to destination")
	}
	return nil
}

package generator

import (
	"bytes"
	"io"
	"strings"
	"text/template"

	"github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/genutil"
)

const ddlMainTemplate = `
// Code generated by tqbuilder DO NOT EDIT.

package main

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/errors"
	"github.com/benchhub/benchhub/lib/tqbuilder/generator"

{{ range .DDLImports }}
	{{ .Name }} "{{ .Path }}"
{{- end }}
)

var (
	logReg = dlog.NewRegistry()
	log = logReg.NewLogger()
)

func main() {
	ddls := []generator.DDLTables{
{{ range .DDLImports }}
		{
			Path: "{{ .Path }}",
			Tables: {{ .Name }}.Tables(),
		},
{{- end }}
	}
	merr := errors.NewMultiErr()
	for _, ddl := range ddls {
		merr.Append(generator.GenDDL(ddl))
	}
	if merr.HasError() {
		log.Fatal(merr.ErrorOrNil())
	}
}
`

type DDLImport struct {
	Name string
	Path string
}

type DDLTables struct {
	Path   string
	Tables []ddl.TableDef
}

// GenDDLMain generates a main.go file that can generate go binding and SQL
// based on DDL ast written in go.
func GenDDLMain(dst io.Writer, importPrefix string, ddls []string) error {
	// Generate unique import name based on path
	// TODO: it is no longer unique if there are packages with same name ...\
	var ddlImports []DDLImport
	for _, ddlPath := range ddls {
		// user/schema/ddl -> userddl
		segs := strings.Split(ddlPath, "/")
		if len(segs) < 3 {
			return errors.Errorf("expect at least 3 segments in ddl path, got %d from %s", len(segs), ddlPath)
		}
		last := len(segs) - 1
		name := segs[last-2] + segs[last]
		ddlImports = append(ddlImports, DDLImport{
			Name: name,
			// core/services/user/schema/ddl -> github.com/benchhub/benchhub/core/services/user/schema/ddl
			Path: importPrefix + "/" + ddlPath,
		})
	}

	// Render template
	tpl, err := template.New("ddlmain").Parse(ddlMainTemplate)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, map[string]interface{}{
		"DDLImports": ddlImports,
	}); err != nil {
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

// GenDDL generates table(s) for a single package (service).
func GenDDL(d DDLTables) error {
	log.Infof("GenDDL TODO: %s %d", d.Path, len(d.Tables))
	return nil
}

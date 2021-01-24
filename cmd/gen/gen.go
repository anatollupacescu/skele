package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Spec struct {
	Pkg  string
	Funs []Fun
}

type Fun struct {
	Name   string
	Prepos []Prepos
}

type Prepos struct {
	Given, Assert string
}

func extractGiven(in string) string {
	if !strings.Contains(in, "|") {
		return in
	}

	for _, part := range strings.Split(in, "|") {
		return strings.Trim(part, " ")
	}

	return ""
}

func extractAssert(in string) string {
	if !strings.Contains(in, "|") {
		return in
	}

	for i, part := range strings.Split(in, "|") {
		if i == 0 {
			continue
		}
		return strings.Trim(part, " ")
	}

	return ""
}

func extractTestName(in string) (out string) {
	split := strings.Split(in, " ")
	for _, word := range split {
		out += strings.Title(word)
	}

	return
}

func (m *machine) write() {
	t := template.Must(template.New("testfile").Parse(testfile))

	for _, ms := range m.specs {
		var spec Spec
		spec.Pkg = ms.pkg

		for _, fun := range ms.fun {
			var pp []Prepos

			for _, fp := range fun.pre {
				pp = append(pp, Prepos{
					Given:  extractGiven(fp.comm),
					Assert: "error",
				})
			}

			for _, fp := range fun.pos {
				pp = append(pp, Prepos{
					Given:  extractGiven(fp.comm),
					Assert: extractAssert(fp.comm),
				})
			}

			spec.Funs = append(spec.Funs, Fun{
				Name:   extractTestName(fun.comm),
				Prepos: pp,
			})
		}

		var filename = ms.filename

		if ms.pkg != "" {
			_ = os.Mkdir(ms.pkg, 0755)
			filename = ms.pkg + "/" + filename
		}

		outFile, err := os.Create(filename)
		if err != nil {
			log.Fatalf("create %s: %v", ms.filename, err)
		}

		if err := t.Execute(outFile, spec); err != nil {
			log.Fatalf("execute: %v", err)
		}
	}
}

var testfile = `package {{.Pkg}}

import "testing"

{{ range .Funs }}func Test{{.Name}}(t *testing.T) { {{ range .Prepos }}
	t.Run("given {{.Given}}", func(t *testing.T) {

		t.Run("assert {{.Assert}}", func(t *testing.T) {

		})
	}){{ end }}
}

{{ end }}`

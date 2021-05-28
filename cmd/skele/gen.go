package main

import (
	"io"
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

var templt = template.Must(template.New("testfile").Parse(testfile))

func (m *machine) write() {
	for _, ms := range m.specs {
		spec := Spec{
			Pkg: ms.pkg,
		}
		folder := ms.fol
		if folder == "" {
			folder = ms.pkg
		}
		for _, file := range ms.file {
			for _, fun := range file.fun {
				funName := toCamelCase(fun.name)
				if contains(spec.Funs, funName) {
					log.Fatalf("duplicate fun declaration: %s", fun.name)
				}
				var pp []Prepos
				for _, fp := range fun.pre {
					pp = append(pp, Prepos{
						Given:  fp.domain,
						Assert: "error",
					})
					if fp.impl == "" {
						continue
					}
					pp = append(pp, Prepos{
						Given:  fp.impl,
						Assert: "error",
					})
				}
				for _, fp := range fun.pos {
					pp = append(pp, Prepos{
						Given:  fp.domain,
						Assert: "success",
					})
					if fp.impl == "" {
						continue
					}
					pp = append(pp, Prepos{
						Given:  fp.impl,
						Assert: "error",
					})
				}
				spec.Funs = append(spec.Funs, Fun{
					Name:   funName,
					Prepos: pp,
				})
			}
			filename := file.name
			if folder != "" {
				_ = os.Mkdir(folder, 0755)
				filename = folder + "/" + filename
			}
			outFile, err := os.Create(filename)
			if err != nil {
				log.Fatalf("create %s: %v", filename, err)
			}

			if err := templt.Execute(outFile, spec); err != nil {
				log.Fatalf("execute: %v", err)
			}
		}
		if len(ms.doc) == 0 {
			continue
		}
		docPath := folder + "/doc.go"
		docFile, err := os.Create(docPath)
		if err != nil {
			log.Fatalf("create file %s: %v", docPath, err)
		}
		contents := "// " + strings.Join(ms.doc, "\n// ") + "\n\n" + "package " + ms.pkg + "\n"
		if _, err := io.WriteString(docFile, contents); err != nil {
			log.Fatalf("write doc.go for %s", ms.pkg)
		}
	}
}

func toCamelCase(in string) (out string) {
	split := strings.Split(in, " ")
	for _, word := range split {
		out += strings.Title(word)
	}

	return
}

func contains(funs []Fun, funName string) bool {
	for _, f := range funs {
		if f.Name == funName {
			return true
		}
	}
	return false
}

var testfile = `package {{.Pkg}}

import "testing"
{{ range .Funs }}
func Test{{.Name}}(t *testing.T) { {{ range .Prepos }}
	t.Run("given {{.Given}}", func(t *testing.T) {
		t.Run("assert {{.Assert}}", func(t *testing.T) {
		})
	}){{ end }}
}
{{ end }}`

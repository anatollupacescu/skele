package main

import (
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

type Pkg struct {
	Name string
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
	for _, ms := range m.pkgs {
		pkg := Pkg{
			Name: ms.name,
		}
		// write each file
		for _, file := range ms.file {
			for _, fun := range file.fun {
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
				funName := toCamelCase(fun.name)
				pkg.Funs = append(pkg.Funs, Fun{
					Name:   funName,
					Prepos: pp,
				})
			}
			filename := file.name
			folder := ms.fol
			if ms.fol != "" {
				_ = os.Mkdir(folder, 0755)
				filename = folder + "/" + filename
			}
			outFile, err := os.Create(filename)
			if err != nil {
				log.Fatalf("create %s: %v", filename, err)
			}
			if err := templt.Execute(outFile, pkg); err != nil {
				log.Fatalf("execute: %v", err)
			}
			pkg.Funs = nil // start over
		}
		// write doc
		if len(ms.doc) > 0 {
			folder := "."
			if ms.fol != "" {
				folder = ms.fol
			}
			docPath := folder + "/doc.go"
			docFile, err := os.Create(docPath)
			if err != nil {
				log.Fatalf("create file %s: %v", docPath, err)
			}
			contents := "// " + strings.Join(ms.doc, "\n// ") + "\n\n" + "package " + ms.name + "\n"
			if _, err := io.WriteString(docFile, contents); err != nil {
				log.Fatalf("write doc.go for %s", ms.name)
			}
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

var testfile = `package {{.Name}}

import "testing"
{{ range .Funs }}
func Test{{.Name}}(t *testing.T) { {{ range .Prepos }}
	t.Run("given {{.Given}}", func(t *testing.T) {
		t.Run("assert {{.Assert}}", func(t *testing.T) {
		})
	}){{ end }}
}
{{ end }}`

package main

import (
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	Given      string
	TestCases  []TestCase
	Assertions []string
}

type TestCase struct {
	Case       string
	Assertions []string
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
					tcs := testCases(fp.tcS, "error")
					pp = append(pp, Prepos{
						Given:      fp.succ,
						TestCases:  tcs,
						Assertions: assertions(tcs, fp.arnS, "error"),
					})
					if fp.fail == "" {
						continue
					}
					tcs = testCases(fp.tcF, "error")
					pp = append(pp, Prepos{
						Given:      fp.fail,
						TestCases:  tcs,
						Assertions: assertions(tcs, fp.arnF, "error"),
					})
				}
				for _, fp := range fun.pos {
					tcs := testCases(fp.tcS, "success")
					pp = append(pp, Prepos{
						Given:      fp.succ,
						TestCases:  tcs,
						Assertions: assertions(tcs, fp.arnS, "success"),
					})
					if fp.fail == "" {
						continue
					}
					tcs = testCases(fp.tcF, "error")
					pp = append(pp, Prepos{
						Given:      fp.fail,
						TestCases:  tcs,
						Assertions: assertions(tcs, fp.arnF, "error"),
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
				_ = os.Mkdir(folder, 0o755)
				filename = folder + "/" + filename
			}
			outFile, err := os.Create(filename)
			if err != nil {
				log.Fatalf("create %s: %s", filename, err)
			}
			if err := templt.Execute(outFile, pkg); err != nil {
				log.Fatalf("execute: %s", err)
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
				log.Fatalf("create file %s: %s", docPath, err)
			}
			contents := "// " + strings.Join(ms.doc, "\n// ") + "\n\n" + "package " + ms.name + "\n"
			if _, err := io.WriteString(docFile, contents); err != nil {
				log.Fatalf("write doc.go for package %s: %s", pkg.Name, err)
			}
		}
	}
}

func toCamelCase(in string) (out string) {
	split := strings.Split(in, " ")
	for _, word := range split {
		out += cases.Title(language.Und).String(word)
	}

	return
}

func testCases(in []string, at string) (cc []TestCase) {
	if len(in) == 0 {
		return
	}

	cc = make([]TestCase, 0, len(in))

	for _, v := range in {
		cs, aa := withAssertions(v)
		if cs == "" {
			cs = v
		}
		if len(aa) == 0 {
			aa = append(aa, at)
		}
		c := TestCase{Case: cs, Assertions: aa}
		cc = append(cc, c)
	}

	return
}

func assertions(tcs []TestCase, arn []string, in string) []string {
	if len(tcs) > 0 && len(arn) == 0 {
		return nil
	}

	if len(arn) == 0 {
		return []string{in}
	}

	return arn
}

func withAssertions(v string) (cs string, aa []string) {
	ss := strings.Split(v, ", assert ")
	for i, s := range ss {
		if i == 0 {
			continue
		}
		aa = append(aa, s)
	}

	if len(ss) > 0 {
		cs = ss[0]
	}

	return
}

var testfile = `package {{.Name}}

import "testing"
{{ range .Funs }}
func Test{{.Name}}(t *testing.T) { {{- range .Prepos }}
	t.Run("given {{.Given}}", func(t *testing.T) { {{- range .TestCases}}
		t.Run("{{.Case}}", func(t *testing.T) { {{- range .Assertions}}
			t.Run("assert {{.}}", func(t *testing.T) {
			}){{end}}
		}){{ end }}{{ range .Assertions}}
		t.Run("assert {{.}}", func(t *testing.T) {
		}){{end}}
	}){{end}}
}
{{ end }}`

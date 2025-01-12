//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package {{ .Package }}

import (
{{- range $day, $funs := .Days }}
	"github.com/kgaughan/aoc/2015/solutions/day{{ $day }}"
{{- end }}
)

func init() {
	days = map[int][]func(string){
	{{- range $day, $fns := .Days }}
		{{ $day }}: {
		{{- range $fns }}
			{{ . }},
		{{- end }}
		},
	{{- end }}
	}
}
`))

func main() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	params := struct {
		Package string
		Days    map[int64][]string
	}{
		Package: os.Getenv("GOPACKAGE"),
		Days:    make(map[int64][]string, 25),
	}

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		panic(err)
	}

	for _, entry := range files {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "day") {
			name := entry.Name()
			suffix := name[3:]
			if d, err := strconv.ParseInt(suffix, 10, 64); err != nil {
				log.Printf("bad day %q: %v", name, err)
			} else {
				dayFiles, err := ioutil.ReadDir(path.Join(cwd, name))
				if err != nil {
					panic(err)
				}
				part := 1
				for _, partEntry := range dayFiles {
					if matched, err := path.Match("part*.go", partEntry.Name()); err != nil {
						panic(err)
					} else if matched {
						params.Days[d] = append(params.Days[d], fmt.Sprintf("day%v.Part%v", d, part))
						part++
					}
				}
			}
		}
	}

	var w bytes.Buffer
	if err := tmpl.Execute(&w, params); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(path.Join(cwd, "days.go"), w.Bytes(), 0o644); err != nil {
		panic(err)
	}
}

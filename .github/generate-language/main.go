package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"text/template"
)

func main() {
	bs, err := ioutil.ReadFile(".github/generate-language/data.txt")
	if err != nil {
		panic(err)
	}
	langs := []Lang{}
	for _, v := range strings.Split(string(bs), "\n") {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		vv := strings.Split(v, ",")
		if len(vv) != 2 {
			panic(fmt.Sprintf(v))
		}
		desc := vv[0]
		str := strings.ToLower(vv[1])
		x := []byte(str)
		name := string(append([]byte{x[0] + 'A' - 'a'}, x[1:]...))

		langs = append(langs, Lang{
			Str:  str,
			Name: name,
			Desc: desc,
		})
	}

	sort.Slice(langs, func(i, j int) bool {
		return langs[i].Str < langs[j].Str
	})

	readme := generate(langs)
	err = ioutil.WriteFile("language.go", []byte(readme), 0o666)
	if err != nil {
		panic(err)
	}
}

type Lang struct {
	Name string
	Str  string
	Desc string
}

func generate(langs []Lang) string {
	t, err := template.New("").Parse(tem)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, langs)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

var tem = `package baidufanyi

type Language string

const (
{{ range . }}Language{{.Name}} Language = "{{.Str}}" // {{.Desc}}
{{ end }}
)
`

package main

//go:generate statik -src=templates/

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/ryota0624/typed-json-struct/statik"
)

type Param struct {
	Package    string
	DetailName string
	TypeDef    string
}

var (
	defaultDestName = "___empty___"
	detailName      = flag.String("detail_name", "", ``)
	packageName     = flag.String("package_name", "", ``)
	typeDefName     = flag.String("type_def_name", "", ``)
	destName        = flag.String("dest_name", defaultDestName, ``)
)

func main() {
	flag.Parse()
	param := Param{
		DetailName: *detailName,
		Package:    *packageName,
		TypeDef:    *typeDefName,
	}
	statikFS, err := fs.New()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	tmplF, err := statikFS.Open("/struct_def.go.tpl")
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	tmplBytes, err := ioutil.ReadAll(tmplF)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	tmpl, err := template.New("struct_def").Parse(string(tmplBytes))
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	var dest io.Writer
	if *destName == defaultDestName {
		dest = os.Stdout
	} else {
		writeFile, err := os.OpenFile(*destName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		defer writeFile.Close()
		dest = writeFile
	}

	err = tmpl.Execute(dest, param)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

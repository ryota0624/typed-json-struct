package main

//go:generate statik -src=templates/

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/ryota0624/typed-json-struct/statik"
)

type Param struct {
	Package         string
	Interface       string
	EnumConstructor string
}

var (
	emptyString         = "___empty___"
	interfaceName       = flag.String("interface_name", "", ``)
	packageName         = flag.String("package_name", "", ``)
	typeConstructorName = flag.String("type_constructor_name", emptyString, ``)
	destName            = flag.String("dest_name", emptyString, ``)
)

func main() {
	flag.Parse()

	enumConstructor := *typeConstructorName
	if *typeConstructorName == emptyString {
		enumConstructor = fmt.Sprintf("%sTypeConstructor", *interfaceName)
	}
	param := Param{
		Interface:       *interfaceName,
		Package:         *packageName,
		EnumConstructor: enumConstructor,
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
	if *destName == emptyString {
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

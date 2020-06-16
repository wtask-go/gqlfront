//go:generate go-bindata -ignore \.go$ -nometadata -pkg assets .
//go:generate go fmt .

// Package assets contains required templates to build GraphQL frontend.
//
// Sources of the HTML templates for GraphQL Playground can be found here:
//	* https://github.com/prisma-labs/graphql-playground/tree/master/packages/graphql-playground-html
//
// Required JS components for GraphQL Playground are available here:
// 	* https://www.jsdelivr.com/package/npm/graphql-playground-react
//
package assets

import (
	"bytes"
	"html/template"
)

func compileAsset(name string) (*template.Template, error) {
	data, err := Asset(name)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(bytes.NewBuffer(data).String())
}

// Playground builds template from asset data.
func Playground() (*template.Template, error) {
	return compileAsset("playground.html")
}

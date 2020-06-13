// Package playground contains http handler to display
package playground

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/wtask-go/gqlfront/internal/assets"
)

// Option is used to define optional Playground settings
type Option func(*settings)

// WithTitle replaces default Playground title with specified one.
func WithTitle(title string) Option {
	return func(s *settings) {
		s.Title = title
	}
}

// WithReactBuild replaces default version of `graphql-playground-react` components with given one.
// Check https://www.jsdelivr.com/package/npm/graphql-playground-react for available versions.
func WithReactBuild(version string) Option {
	if version != "" && !strings.HasPrefix(version, "@") {
		version = "@" + version
	}
	return func(s *settings) {
		s.ReactBuild = version
	}
}

type settings struct {
	Title      string
	ReactBuild string
	Endpoint   string
}

const (
	// DefaultReactBuild represents verified version of `graphql-playground-react` components,
	// see https://www.jsdelivr.com/package/npm/graphql-playground-react
	DefaultReactBuild = "1.7.22"
	// LatestReactBuild provides using of latest available build of `graphql-playground-react` components,
	// and may break Playground in that case.
	LatestReactBuild = ""
)

var (
	playground = template.Must(assets.Playground())
)

// Handler builds http handler to display GraphQL Playground page,
// where endpoint is the target of GraphQL API handler.
// If error will occur on building stage the panic is throwed.
func Handler(endpoint string, options ...Option) http.HandlerFunc {
	data := settings{
		Title:      "GraphQL Playground",
		ReactBuild: "@" + DefaultReactBuild,
		Endpoint:   endpoint,
	}
	for _, o := range options {
		o(&data)
	}
	buf := &bytes.Buffer{}
	err := playground.Execute(buf, map[string]string{
		"title":    data.Title,
		"version":  data.ReactBuild,
		"endpoint": data.Endpoint,
	})
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		_, _ = buf.WriteTo(w)
	}
}

# gqlfront
GraphQL frontend handlers.

Inspired by [gqlgen/graphql/playground](https://github.com/99designs/gqlgen/tree/v0.11.3/graphql/playground).

## Automation

* Install required tools locally before working with module:

```cmd
go generate tools.go
```

* Regenerate required `assets` if you made changes into html-sources as follows:

```cmd
go generate assets.go
```

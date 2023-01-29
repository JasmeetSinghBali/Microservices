> # Golang+ Graphql + MongoDB

> initial-setup

```bash

mkdir graphql_mongodb_go
cd graphql_mongodb_go

go get github.com/99designs/gqlgen

# add gqlgen to tools.go, run in git bash
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go


# graphql-gen init, inside graphql_mongodb_go
go run github.com/99designs/gqlgen init

# models_gen.go and generated.go can be cleared out

```

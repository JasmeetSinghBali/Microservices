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

# models_gen.go , generated.go, schema.graphqls can be cleared out

```

> dev

```bash
# the schema.graphqls is the dev reff origin
# each time you start dev graphql mutations/queries the schema must be decided upon first
schema.graphqls is that file

# after the schema.graphqls is sorted, use below command to generate all the files based of the schema defined by you
go run github.com/99designs/gqlgen generate

```

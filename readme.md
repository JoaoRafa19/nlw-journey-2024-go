# NLW Journey Go api

Nessa api é utilizado o `goapi-gen` para gerar código boilerplate para facilitar o desenvolvimento da API.
`sqlc` para gerar as interfaces das entidades das tabelas dos bancos de dados (não é um ORM)


## Deps

#### Install all deps:
```shell
go mod tidy
```

- **goapi-gen**
```shell
go install github.com/discord-gophers/goapi-gen@latest
```

- **tern**
```shell
 go install github.com/jackc/tern/v2@latest
 ```

- **sqlc**


## Generate

```shell
goapi-gen --out [output-file] [input json spec file]
```
Ex.:
```shell
goapi-gen --out .\internal\api\spec\journey.gen.spec.go .\internal\api\spec\journey.spec.json
```
# NLW Journey Go api

Nessa api é utilizado o `goapi-gen` para gerar código boilerplate para facilitar o desenvolvimento da API.
`sqlc` para gerar as interfaces das entidades das tabelas dos bancos de dados (não é um ORM)

## Generate

```shell
go generate ./...
```

### Generate interfaces implementation
Using goimpl
```shell
go install github.com/josharian/impl@latest
```

ex.:
```shell
impl 'api API' github.com/JoaoRafa19/nlw-journey-2024-go/internal/api/spec.ServerInterface
```

## Deps


#### Install all deps:
```shell
go mod tidy
```

- **gomail** `github.com/wneessen/go-mail`
Usado para escrever emails em go 


- **goapi-gen**
Para gerar boilerplate da api com base na especificação do json
```shell
go install github.com/discord-gophers/goapi-gen@latest
```

- **tern**
```shell
 go install github.com/jackc/tern/v2@latest
 ```

- **sqlc**
```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```



## Generate

```shell
goapi-gen --out [output-file] [input json spec file]
```
Ex.:
```shell
goapi-gen --out .\internal\api\spec\journey.gen.spec.go .\internal\api\spec\journey.spec.json
```
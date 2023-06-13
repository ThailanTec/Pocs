## Pocs criação banco de dados unico endpoint

## Tecnologias utilizadas

- Golang
    * Gin
- Airbyte


## Estrutura da aplicação

```shell
.
├── Collections
│   ├── airbyte_Collections.json
│   ├── Application_collection.json
│   └── Json_ExamplesReq.json
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── dto.go
│   └── intefaces.go
├── README.md
├── service
│   ├── contextS.go
│   ├── dto_map.go
│   └── factory.go
└── strategy
    ├── bigQuery.go
    ├── http.go
    ├── mysql.go
    ├── oracle.go
    ├── postgress.go
    ├── redshift.go
    └── sqlserver.go

4 directories, 19 files

```
## Sobre o código

A aplicação contem 2 padrões de *Design Patterns*:

* Factory 
* Strategy 

Para auxiliar no desenvolvimento e na solução do problema de criação do banco de dados de forma mais rapida. 

## Iniciar a aplicação

*<b>Itens a serem instalados: </b>*

* [Golang](https://go.dev/doc/install)
* [Airbyte](https://airbyte.com/) (Caso deseje rodar a aplicação no seu local)


## Comandos & Rotas

<b>Comandos para iniciar a aplicação:</b>

```shell

go run .

```

A aplicação então vai iniciar a aplicação. 

### Estamos utilizando o *gin*, como biblioteca *http*.

*A rota padrão para acesso é:*

* http://localhost:8080/pocs/connection

Para a criação deve-se ser alterado alguns dados do Json, esses são:

- workspaceId - Deve-se consultar no seu airbyte, qual o seu workspaceID
- <b>*format*</b> 

Os dados acima são respectivos ao airbyte(tirando o format que é referente a aplicação) e devem ser consultado no airbyte seus respectivos dados.
*Obs*: Os payload *json* respectivos a produção já estão com seus dados baseados no ambiente de produção. Então, quando forem utilizados vão criar com perfeição seus respectivos bancos.

O <b>*format*</b> é o principal na utilização da aplicação. Pois ele vai chamar a respectiva função para criar o banco de dados desejado. 

## Sobre o Postman

Na pasta "collections", contida aqui na estrutura, temos a collection para ser utilizada nas requisição. Incluindo as do propio airbyte e as da aplicação. E uma extra, contendo os Json que são utilizados para enviar cada requisição. 

### Configurações possiveis:

- {URL_PRD} = Contem a configuração da rota produção
- {URL_LOCAL}= Contem a configuração da rota local

- {PWD_PRD} = Contem a senha para acessar a produção.
- {USER_PRD} = Contem o usuario para acessar a produção.

- {PWD_LOCAL} = Contem a senha para acessar a produção.
- {USER_LOCAL} = Contem o usuario para acessar a produção.

## Informações importantes

* Para editar o *endpoint* e a *Authorization* que deseja fazer os teste da aplicação. Fica em:

interfaces > const > [Linha 17 até 21]

* Para editar a porta que o servidor vai rodar na maquina:

main > router [Linha 13] Alterar a primeira opção dentro das strings

```go 

	http.ListenAndServe(":8080", router)
    
```

Na aplicação contem const onde a mesma tem o objetivo de setar o padrão do SourceDefinitionID. 


## Rota de GET

*/pocs/getdb*

A rota disponibiliza um exemplo de como vai ser cada body para criação de um modelo de banco de dados. 
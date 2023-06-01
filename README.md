## Pocs criação banco de dados unico endpoint

## Tecnologias utilizadas

- Golang
- Airbyte


## Estrutura da aplicação

```shell
.
├── airbyte_codigo.go
├── contextS.go
├── env.go
├── factory.go
├── go.mod
├── go.sum
├── intefaces.go
├── main.go
├── oracle.go
├── postgress.go
├── README.md
└── sql.go

0 directories, 12 files

```
## Sobre o código

A aplicação contem 2 padrões de *Design Patterns*:

* Factory 
* Strategy 

Para auxiliar no desenvolvimento e na solução do problema de criação do banco de dados de forma mais rapida. 

## Iniciar a aplicação

*<b>Itens a serem instalados: </b>*

* [Docker]( https://docs.docker.com/engine/install/ubuntu/)
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

Abaixo segue os *body's* que são necessarios no postman, para criação dos bancos de dados.

### Para criação em produção:

```json
Postgres PRD 
{
    "sourceDefinitionId": "decd338e-5647-4c0b-adf4-da0e75f5a750",
    "sourceId": "79b7782b-a3d4-4814-bb41-743146d493cb",
    "workspaceId": "c2e9e9ae-1289-4016-8220-a6ddc3f8d225",
    "connectionConfiguration": {
        "dataset_name": "",
        "host": "35.227.17.122",
        "port": 5432,
        "database": "data-dev",
        "username": "dhuodata",
        "sslMode": "disable",
        "schemas": [
            "landing_zone"
        ],
        "password": "LBP5Xt0QpR5&43!N6z3pc",
        "replicationMethod": "true"
    },
    "sourceType": "postgres",
    "name": "Thing",
    "sourceName": "Serve",
    "icon": "icon",
    "format": "postgres"
} 

MysqlServer PRD

{
    "sourceDefinitionId":"b5ea17b1-f170-46dc-bc31-cc744ca984c1",
   "sourceId":"33e2f50d-fdfd-4a5b-9bb4-32350c0f9cc6",
   "workspaceId":"c2e9e9ae-1289-4016-8220-a6ddc3f8d225",
    "connectionConfiguration": {
        "dataset_name": "",
        "host": "enge-sql-server.database.windows.net",
        "schemas": [
            "dto"
        ],
         "sourceType":"mssql",
        "port": 1433,
        "database": "mySampleDatabase",
        "username": "azureuser",
        "sslMode": "disable",
        "password": "Pwd@1234",
        "replicationMethod": "true"
    },
   
    "name": "SQLEMO",
    "sourceName": "SQLEMO",
    "icon": "icon",
    "format":"mssql"
}

```

### Para criação local:

```json


Criação SQLServer Local
{
   "sourceDefinitionId":"b5ea17b1-f170-46dc-bc31-cc744ca984c1",
   "sourceId":"3b714e15-9929-4886-92b5-e267684026b4",
   "workspaceId":"b2cceae1-84f2-4188-b160-ee393190c88d",
    "connectionConfiguration": {
        "dataset_name": "",
        "host": "enge-sql-server.database.windows.net",
        "schemas": [
            "dto"
        ],
         "sourceType":"mssql",
        "port": 1433,
        "database": "mySampleDatabase",
        "username": "azureuser",
        "sslMode": "disable",
        "password": "Pwd@1234",
        "replicationMethod": "true"
    },
   
    "name": "HOJE",
    "sourceName": "HOJE",
    "icon": "icon",
    "format":"mssql"
}

Criação Postgres Local

{
    "sourceDefinitionId": "decd338e-5647-4c0b-adf4-da0e75f5a750",
    "sourceId": "952fde34-051d-4465-a59a-25d46f6341a2",
    "workspaceId": "b2cceae1-84f2-4188-b160-ee393190c88d",
    "connectionConfiguration": {
        "dataset_name": "",
        "host": "35.227.17.122",
        "port": 5432,
        "database": "data-dev",
        "username": "dhuodata",
        "sslMode": "disable",
        "schemas": [
            "landing_zone"
        ],
        "password": "LBP5Xt0QpR5&43!N6z3pc",
        "replicationMethod": "true"
    },
    "sourceType": "postgres",
    "name": "ThailanLINDO",
    "sourceName": "postgres",
    "icon": "icon",
    "format": "postgres"
}

```

Para a criação deve-se ser alterado alguns dados do Json, esses são:

- sourceDefinitionId
- sourceId
- workspaceId
- <b>*format*</b> 

Os dados acima são respectivos ao airbyte(tirando o format que é referente a aplicação) e devem ser consultado no airbyte seus respectivos dados.
*Obs*: Os payload *json* respectivos a produção já estão com seus dados baseados no ambiente de produção. Então, quando forem utilizados vão criar com perfeição seus respectivos bancos.

O <b>*format*</b> é o principal na utilização da aplicação. Pois ele vai chamar a respectiva função para criar o banco de dados desejado. 

## Sobre o Postman

Na pasta "collections", contida aqui na estrutura, temos a collection para ser utilizada nas requisição. Incluindo as do propio airbyte e as da aplicação. 

### Configurações possiveis:

- {URL_PRD} = Contem a configuração da rota produção
- {URL_LOCAL}= Contem a configuração da rota local

- {PWD_PRD} = Contem a senha para acessar a produção.
- {USER_PRD} = Contem o usuario para acessar a produção.

- {PWD_LOCAL} = Contem a senha para acessar a produção.
- {USER_LOCAL} = Contem o usuario para acessar a produção.

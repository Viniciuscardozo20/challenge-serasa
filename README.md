# Desafio Serasa

Candidato: Vinicius da Silva Cardozo

LinkedIn: https://www.linkedin.com/in/vinicius-cardozo-669a15136/

Recrutador: Roberto Krieger

## Requesitos

* [Docker-compose](https://docs.docker.com/compose/install/)

## Setup

Após clonar o repositório, é necessário ir no caminho raiz e rodar o seguinte comando 

    docker-compose up

Isso subirá a API na porta `8082`, o banco na porta `27017` e o serviço legado na `3000`

## Rotas

### Realizar integração com o sistema legado

METÓDO `POST`v1/

    /v1/integration
    
#### Exemplo

No terminal execute

    curl -XPOST 'http://localhost:8082/v1/integration'

Ou em alguma outra interface preferência, por exemplo o Postman ou Insomnia

Saída esperada:

_Status http_: `200 OK`
```json
{
  "status": "success"
}
```

### Login

METÓDO `POST`v1/

    /v1/login
    
#### Exemplo

No terminal execute

    curl -XPOST -H "Content-type: application/json" -d '{"customerDocument":"62824334010"}' 'http://localhost:8082/v1/login''

Ou em alguma outra interface preferência, tipo Postman ou Insomnia por exemplo

**Body** (application/json):

    {
	    "customerDocument": "62824334010"
    }

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjdXN0b21lckRvY3VtZW50IjoiNjI4MjQzMzQwMTAiLCJleHAiOjE2MTExODQ0NjR9.tWO_reQrV3jtgP5fC-1F3oPd1JcR00Lq5RCroWEXkxw"
}
```


### Pegar negativações por cliente

METÓDO `GET`

    /v1/negativations/:customerDocument

**Parametros**: 

* CustomerDocument: Id do documento do cliente 

#### Exemplo

No terminal execute

    curl -XGET -H 'Token: token-recebido' 'http://localhost:8082/v1/negativations/62824334010'

Ou em alguma outra interface preferência, tipo Postman ou Insomnia por exemplo

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": [
    {
      "companyDocument": "23993551000107",
      "companyName": "XPTO S.A.",
      "customerDocument": "62824334010",
      "value": 230.5,
      "contract": "8b441dbb-3bb4-4fc9-9b46-bdaad00a7a98",
      "debtDate": "2015-08-10T23:32:51Z",
      "inclusionDate": "2020-08-10T23:32:51Z"
    }
  ]
}
```


## Executar os testes
### Requesitos

* Golang 1:15

Para executar os testes feitos é necessario entrar na pasta ./api do projeto e rodar o comando:

  `make run-tests` 

A saída esperada será a seguinte: 

    
      go test -race -coverpkg= ./... -coverprofile=./test/cover/cover.out
      ok      challenge-serasa/api/app        0.152s  coverage: 53.8% of statements
      ok      challenge-serasa/api/controller 1.256s  coverage: 67.6% of statements
      ok      challenge-serasa/api/controller/auth    0.030s  coverage: 52.9% of statements
      ok      challenge-serasa/api/controller/cryptoModule    0.029s  coverage: 71.4% of statements
      ok      challenge-serasa/api/database   1.240s  coverage: 81.6% of statements
      ?       challenge-serasa/api/handlers/integration       [no test files]
      ?       challenge-serasa/api/handlers/login     [no test files]
      ?       challenge-serasa/api/handlers/negativations     [no test files]
      ?       challenge-serasa/api/helper_tests/h_database    [no test files]
      ?       challenge-serasa/api/helper_tests/h_mainframe   [no test files]
      ok      challenge-serasa/api/mainframe  0.032s  coverage: 83.3% of statements
## How to build

Since this is a single binary project, no Makefile needed and a simple
`go build` at the project's root should do the trick.

## How to test

Given the assignmet time frame, no unit tests were implemented.
However, the API can be tested directly with cURL or a client of your
preference. Here are some examples with cURL:
```
curl -X POST -H 'Content-Type: application/json' \
     http://someIP:8080/quotation -d '[{"name": "MÓDULO POLI 330W", "quant": 12}]'

curl -X GET -H 'Content-Type: application/json' http://someIP:8080/products
```

## How to run

Edit the .env file as needed, build/run the database container and
execute the binary, i.e.:
```
docker-compose -f build/conexao.compose up -d

go build

./price_api
```

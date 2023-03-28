# go-rest-api
rest api using golang


## REST API Calls

```
cd <project_folder>

go run .

curl http://localhost:8080/AIs

curl http://localhost:8080/AIs \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"AIProduct": "Dolly","Creator": "Databricks"}'

curl http://localhost:8080/AIs/Alphabet
```
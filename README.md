## Запуск
```
    docker compose up
```

## swagger
Доступ по адресу
http://localhost:8080

## curl
```
curl -X 'GET' \
  'http://localhost:8080/access?userID=0f2aa1a7-5b74-43a3-8343-65450e2e1835' \
  -H 'accept: application/json'
```

```
curl -X 'POST' \
  'http://localhost:8080/refresh' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ0NDE2NTcsImlkIjoiNjYzMDRkOWI3NTgxYWQzN2YwZTJlMGQ0Iiwic3ViIjoiMGYyYWExYTctNWI3NC00M2EzLTgzNDMtNjU0NTBlMmUxODM1In0.UV_mphVqFuuictY3c2QXcN8LOAgdUr7SbV1-neHTJEM",
  "refreshToken": "MjE1NTBkYjMtZWIxNC00MTRjLTgxOTQtNmE5YTY2YjhhOWZi"
}'
```

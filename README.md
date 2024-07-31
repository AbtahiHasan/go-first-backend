# Go rss first backend

## project information

- language uses go lang
- database uses postgresql
- for migrations uses goose
- for handle row db queries uses sqlc
- for routing uses chi router
- pq uses db query

### for api authentication header here token auto generated and it's auto saved in user table

```bash
    key authentication
    value ApiKey 2022a6b664f81a07097d8d0a1a5ac4b85c4706de1e8405aafb96ce8f29ba2dba
```

#### how to run this project locally

1. create .env file

2. copy the blow code

```bash

    PORT=8000
    DB_URL=postgres://username:password@localhost:5432/go-first-backend?sslmode=disable

```

3. change db url

4. before run this project make sure go have already installed

5. run blow command into your terminal

```bash
    go build && ./go-first-backend
```

#### postman documentation [postman](https://documenter.getpostman.com/view/30464992/2sA3kbhyTu)

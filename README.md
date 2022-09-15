# GO API Todo List

## Routes

- Get a single task
  - route: `host:8000/api/v1/get/{_id}/`
  - description: replace {_id} with mongodb document id
  - method: GET

- Get all tasks
  - route: `host:8000/api/v1/get/`
  - description: return a list of tasks
  - method: GET

- Insert a new task
  - route: `host:8000/api/v1/craete/`
  - desription: send a json request with these fields:
  ```json
    {
      "name": "name of task",
      "description": "info about task",
      "status": true or false
    }
  ```
  - method: POST

- Delete a task
  - route: `host:8000/api/v1/delete/{_id}/`
  - description: replace {_id} with desired task id returned in create response
  - method: DELETE

- Update/Edit a task
  - route: `host:8000/api/v1/update/{_id}`
  - description: replace {_id} with desired task id and send a json in body same as create route.
  - method: PUT

## How to run

You can run this app using 2 methods:

1. Manual
    1. clone this repo, rename `env.env` to `.env`, fill fields in `.env` file as desired
    2. run `go build .`
    3. run binary named `./go-api-todolist`
 
2. Docker compose 
    1. clone this repo.
    2. run `docker-compose up -d` and wait for app to run on localhost:8000

## Tools

- ENV: https://github.com/joho/godotenv
- MUX: https://github.com/gorilla/mux
- Mongo: https://go.mongodb.org/mongo-driver

## Authors

- Amir
- Max Base

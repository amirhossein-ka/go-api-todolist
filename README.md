# GO API Todo List

## Routes

- Get a single task
  - Route: `host:8000/api/v1/get/{_id}/`
  - Description: replace {_id} with mongodb document id
  - Method: GET

- Get all tasks
  - Route: `host:8000/api/v1/get/`
  - Description: return a list of tasks
  - Method: GET

- Insert a new task
  - Route: `host:8000/api/v1/craete/`
  - Description: send a JSON request with these fields:
  ```json
    {
      "name": "name of task",
      "description": "info about task",
      "status": true
    }
  ```
  - method: POST

- Delete a task
  - Route: `host:8000/api/v1/delete/{_id}/`
  - Description: replace {_id} with desired task id returned in create response
  - Method: DELETE

- Update/Edit a task
  - Route: `host:8000/api/v1/update/{_id}`
  - Description: replace {_id} with desired task id and send a JSON in body same as create route.
  - Method: PUT

## How to run

You can run this app using 2 methods:

1. Manual
    1. Clone this repo, rename `env.env` to `.env`, fill fields in `.env` file as desired
    2. Run `go build .`
    3. Run binary named `./go-api-todolist`
 
2. Docker compose 
    1. Clone this repo.
    2. Run `docker-compose up -d` and wait for app to run on `localhost:8000`

## Tools

- ENV: https://github.com/joho/godotenv
- MUX: https://github.com/gorilla/mux
- Mongo: https://go.mongodb.org/mongo-driver

## Authors

- Amir
- Max Base

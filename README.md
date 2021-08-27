# Game-Project (Lucas Augusto Sales)

Game-Project is a simple backend CRUD for a game application.

## Tech

Game-Project uses te following technologies:

- [Go] - The language for the backend
- [PostgreSQL] - Relational database
- [Gorilla-Mux] - fast and lightweight router framework

## Routes
- [GET - "/user"]
- [POST - "/user"]
- [PUT - "/user/{userId}/state"]
- [GET - "/user/{userId}/state"]
- [PUT - "/user/{userId}/friends"]
- [GET - "/user/{userId}/friends"]

## Collections
You may export the collections at the path ./data/collections.json to your Postman/Insomnia.

## Starting the application
Chose one of the options below:

### 1 - Docker-Compose
Simply run: ```docker-compose up -d ``` and call the endpoints.

If you wish to test the application using docker-compose. Use base url as ```http://localhost:8082```

### 2 - Starting the plain application
Remove the application container from docker-compose.
Run:
```
docker-compose up -d
make clean
make run
```

If you wish to test starting the plain application. Use base url as ```http://localhost:8080```

## Executing unitary tests
Simply run:
```
make test
```

Sadly I couldn't make the coverage 100% and add integration tests :(


   

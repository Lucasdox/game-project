# Game-Project (Lucas Augusto Sales)



Game-Project is a simple backend CRUD for a game application.


## Features

- Import a HTML file and watch it magically convert to Markdown
- Drag and drop images (requires your Dropbox account be linked)
- Import and save files from GitHub, Dropbox, Google Drive and One Drive
- Drag and drop markdown and HTML files into Dillinger
- Export documents as Markdown, HTML and PDF

## Tech

Game-Project uses te following technologies:

- [Go] - The language for the backend
- [PostgreSQL] - Relational database
- [Gorilla-Mux] - fast and lightweight router framework

## Routes
[GET - "/user"]
[POST - "/user"]
[PUT - "/user/{userId}/state"]
[GET - "/user/{userId}/state"]
[PUT - "/user/{userId}/friends"]
[GET - "/user/{userId}/friends"]

## Starting the application
Chose one of the options below:

### 1 - Docker-Compose
Simply run: ```docker-compose up -d ``` and call the endpoints.

### 2 - Starting the plain application
Remove the application container from docker-compose.
Run:
```
docker-compose up -d
make clean
make run
```
   

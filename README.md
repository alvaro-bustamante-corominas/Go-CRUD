# Go CRUD Tasks

This is a simple CRUD (Create, Read, Update, Delete) application written in Go. It utilizes MySQL to store tasks.

## Methods Available

- **POST /task**: Create a new task.
- **DELETE /task/:id**: Delete a task by its ID.
- **GET /task**: Retrieve all tasks.
- **PUT /task/:id**: Update a task.
- **PATCH /task/:id**: Update the status of a task.

## Installation and Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/alvaro-bustamante-corominas/Go-CRUD.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-crud-tasks
    ```

3. Build the Docker image and start the containers:

    ```bash
    docker-compose up --build
    ```

4. Access the application at [http://localhost:8080](http://localhost:8080).

## Requirements

- Docker
- Docker Compose

## Technologies Used

- Go
- MySQL
- Docker


# CRUD-Redis

## Introduction
This repository contains a CRUD application that utilizes PostgresSQL as the database with Redis for caching. It provides basic functionality for Create, Read, Update, and Delete operations.

## Migration
To set up the necessary database schema, utilize Goose migration tool. Run the following command:

```bash
goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:5432/golangpg?sslmode=disable" up
```
## Setup
1. Ensure you have Redis installed and running locally.
2. Set up your environment variables by creating a .env file. You may refer to .env.example for guidance.
3. Run the server to start the application.
    ``` bash
    go run main.go
    ```
## Usage
Once the server is up and running, you can interact with the CRUD operations through HTTP endpoints. Refer to the API documentation for details on available endpoints and their usage.

# gin-api: REST API built with Gin + GORM

Sample project to play with:
- [Gin](https://github.com/gin-gonic/gin) (Web Framework)
- [Air](https://github.com/air-verse/air) (Live Reload)
- [GORM](https://github.com/go-gorm/gorm) (ORM)
- [PostgreSQL](https://www.postgresql.org/)

## Requirements

This repo uses [Docker](https://www.docker.com/) to create a standarized environment for development.

## Development

Run `make run` to run the API in development mode.

## Testing

Run `make test` to run the test suite.

## Endpoints

### Todos API

#### `GET /todos`
- **Description**: Retrieve a list of all to-dos.
- **Response Code**: 200
- **Response**:
    ```json
    [
        {
            "id": 1,
            "message": "Buy milk",
        }
    ]
    ```

#### `POST /todos`
- **Description**: Create a new to-do.
- **Response Code**: 201
- **Request Body**:
    ```json
    {
        "message": "Buy milk",
    }
    ```
- **Response**:
    ```json
    {
        "id": 1,
        "message": "Buy milk",
    }
    ```

#### `GET /todos/:id`
- **Description**: Retrieve a to-do by ID.
- **Response Code**: 200
- **Response**:
    ```json
    {
        "id": 1,
        "message": "Buy milk",
    }
    ```

#### `PUT /todos/:id`
- **Description**: Update a to-do by ID.
- **Response Code**: 200
- **Request Body**:
    ```json
    {
        "message": "Buy two galloons of milk",
    }
    ```
- **Response**:
    ```json
    {
        "id": 1,
        "message": "Buy two galloons of milk",
    }
    ```

#### `DELETE /todos/:id`
- **Description**: Delete a to-do by ID.
- **Response Code**: 204
- **Response**: N/A

### URL Shortener API

#### `POST /u/shorten`
- **Description**: Create a short URL.
- **Response Code**: 200
- **Request Body**:
    ```json
    {
        "url": "https://my.super/duper/ultra/long/url",
    }
    ```
- **Response**:
    ```json
    {
        "short": "https://sho.rt/abc"
    }
    ```

#### `GET /u/:shortCode`
- **Description**: Redirect from a short URL to the original URL.
- **Response Code**: 302
- **Response**: N / A

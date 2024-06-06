# GoLang API AuthenticationProject

This is a RESTful API built with GoLang and Fiber framework. The project includes basic CRUD operations for user management.

## Project Structure

```bash
├── .github/
├── configs/
├── controllers/
│   └── user_controller.go
├── models/
│   └── user_model.go
├── responses/
│   └── user_response.go
├── routes/
│   └── user_route.go
├── .env
├── .gitignore
├── README.md
├── docker-compose.yml
├── entrypoint.sh
├── go.mod
├── go.sum
└── main.go
```

## Routes

The routes for the User API are defined in the `routes/user_route.go` file. The following endpoints are available:

- `POST /user`: Create a new user
- `GET /user/:userId`: Retrieve a specific user by ID
- `PUT /user/:userId`: Update a specific user by ID
- `DELETE /user/:userId`: Delete a specific user by ID
- `GET /users`: Retrieve all users

## Getting Started

### Prerequisites

- Go (version 1.16+)
- Docker (optional, for containerization)
- Fiber v2 (Go web framework)

### Installation

1. Clone the repository:

```sh
git clone git@github.com:datalab-api/go_api-rest.git
cd go_api-rest
```

2. Install Go dependencies:

```sh
go mod tidy
```

3. Set up the environment variables. Create a `.env` file in the root directory and add the necessary variables:

```env
# Example .env file content
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```

4. Run the application:

```sh
go run main.go
```

### Using Docker

If you prefer to use Docker, you can start the application with Docker Compose:

1. Build and start the containers:

```sh
docker-compose up --build
```

2. The API should now be running on `http://localhost:8080`.

## Usage

You can use tools like [Postman](https://www.postman.com/) or [cURL](https://curl.se/) to test the endpoints.

### Example Requests

- **Create a User:**

```sh
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john.doe@example.com"}'
```

- **Get a User by ID:**

```sh
curl http://localhost:8080/user/{userId}
```

- **Update a User:**

```sh
curl -X PUT http://localhost:8080/user/{userId} -H "Content-Type: application/json" -d '{"name":"Jane Doe"}'
```

- **Delete a User:**

```sh
curl -X DELETE http://localhost:8080/user/{userId}
```

- **Get All Users:**

```sh
curl http://localhost:8080/users
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure your code follows the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

Feel free to modify the README content to better fit your project's specific details and requirements.

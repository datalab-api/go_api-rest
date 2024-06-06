
# GoLang API Project
---
Welcome to the GoLang API Project! This project is a RESTful API built using GoLang and the Fiber framework. It serves as a foundational example of how to set up a robust and efficient API with basic CRUD (Create, Read, Update, Delete) operations for user management. Whether you're a seasoned developer looking to integrate GoLang and Fiber into your toolkit or a beginner eager to learn more about API development in GoLang, this project provides a clear and structured starting point.

This README will guide you through the project structure, installation steps, and usage instructions to help you get started quickly. You'll also find details on how to contribute to the project and links to relevant documentation.

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

## User Data Model

The User model is defined in `models/user_model.go` and represents the structure of a user in the database. Here is the code:

```go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Address string            `json:"address,omitempty" validate:"required"`
	Phone  string             `json:"phone,omitempty" validate:"required"`
}
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
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
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
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"John Doe", "address":"123 Main St", "phone":"123-456-7890"}'
```

- **Get a User by ID:**

```sh
curl http://localhost:8080/user/{userId}
```

- **Update a User:**

```sh
curl -X PUT http://localhost:8080/user/{userId} -H "Content-Type: application/json" -d '{"name":"Jane Doe", "address":"456 Elm St", "phone":"987-654-3210"}'
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

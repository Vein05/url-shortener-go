# URL Shortener Service
This project is a URL shortener service written in Go. It uses the Gin web framework for HTTP handling, MongoDB for storing URLs, and Godotenv for managing environment variables. Currently, it only created 4 characters for the shortened links, but in the future it will accept custom words.

![Logo](/static/images/logo.jpg =400x300)

## Features

- Accepts long URLs and returns shortened URLs.
- Stores and retrieves URL mappings from a MongoDB database.
- Generates random short URLs.
- Checks if a URL already has a shortened version before creating a new one, returns the already saved response.
- Generates a QR code of the shortened URL. 

## Prerequisites

- Go 1.21 or above
- Gin 
- MongoDB
- A `.env` file with the MongoDB connection string & env package

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/vein05/url-shortener-go.git
    cd url-shortener-go
    ```

2. Create a `.env` file in the root directory and add your MongoDB connection string:

    ```sh
    MONGOURL=your_mongodb_connection_string
    ```

3. Install the necessary Go modules:

    ```sh
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run main.go
    ```

The server will start on port `8080`.

## Endpoints

### `POST /add-url`

**Description:** Accepts a long URL and returns a shortened URL.

**Request Body:**

```json
{
  "url": "https://example.com"
}
```

**Response:**

```json
{
  "newURL": "abcd"
}
```

## Static Files

The static files for the frontend are served from the `./static/` directory.

## Project Structure

```
.
├── main.go
├── go.mod
├── go.sum
├── .env
└── static
    └── index.html
```

## Code Overview

### Main Function

The `main` function initializes the environment variables, connects to MongoDB, sets up the Gin server, and defines the routes.

### MongoDB Connection

The MongoDB connection is established using the connection string from the `.env` file. The `AbodeDB` database is used for storing URL mappings.

### Routes

- `POST /add-url`: Handles the URL shortening request.

### Utility Functions

- `getNewURL`: Generates a new short URL if the provided URL does not already have one.
- `checkUrl`: Checks if a URL already exists in the database.
- `getRandom`: Generates a random 4-character string for the short URL.

## TODO

- Create Custom words for the shortened urls + populate with the new logo
- Make the frontend more clear and concise to work with.
- ~~Make a dynamic url handler in GO.~~
## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to contribute to this project by submitting issues or pull requests. This project isn't completed and more changes are going to come.
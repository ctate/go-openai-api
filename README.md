# go-openai-api 🧠

This project is a Go-powered API server built with [Gin](https://github.com/gin-gonic/gin) and [Go OpenAI](https://github.com/sashabaranov/go-openai) that interacts with the OpenAI API, enabling seamless streaming of responses.

## Features

- Fast HTTP server using Gin
- Streamlined interaction with OpenAI API
- Real-time streaming of responses

## Prerequisites

- Go (version 1.21.x or higher)
- An active OpenAI API Key

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ctate/go-openai-api.git
cd go-openai-api
```

### 2. Set up Environment Variables

Before running the project, add an `.env` file with these environment variable:

```bash
API_PORT="8080"
OPEN_API_KEY="your_openai_api_key"
```

### 3. Run the Server

To start the server:

```bash
go run .
```

By default, the server will start on `http://localhost:8080`.

## Endpoints

- **POST** `/chat`: Streams chat completions from OpenAI API

## Usage Example

Here's a quick example using `curl` to interact with the API:

```bash
curl -XPOST --no-buffer http://localhost:8080/chat -H "Content-Type: application/json" -d '{"prompt":"Lorem ipsum"}'
```

## Contributing

Contributions are welcome! Feel free to submit pull requests or create issues for bugs and feature requests.

## License

[MIT](https://github.com/ctate/go-openai-api/blob/main/LICENSE)

# Simple Go Web App

This is a simple Go web application that serves a static HTML page along with information about the server's port and environment.

## Build and Run

### Prerequisites

- Docker
- Go (if you want to run locally without Docker)

### Build with Docker

```docker build -t sample-webapp .```

### Run with Docker
```docker run -p 3000:3000 yuni2/sample-webapp```

### Build and Run Locally
```bash
go build -o webapp main.go
./webapp
```

## Usage
Once the application is running, open your web browser and navigate to http://localhost:3000. You will see a simple HTML page displaying information about the app's port and environment.

## Docker Images
The Docker image is built in two stages. The first stage uses the golang:1.21-alpine3.19 image to build the Go application, and the second stage uses alpine:3.19 as the base image to create a minimal image with just the built executable.

## Project Structure
main.go: The main Go code for the web server.
static/index.html: The static HTML page with embedded JavaScript to fetch and display server information.

## Environment Variables
PORT: Specifies the port on which the server will run. Default is 3000.
ENV: Specifies the environment of the server. Default is dev.

## Additional Information
The HTML page includes styling for a clean and simple layout.
JavaScript is used to fetch and display the server's port and environment dynamically.
Feel free to customize and extend this project according to your needs!

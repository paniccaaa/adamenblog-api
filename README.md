# AdamenBlog API

This repository contains the source code for the AdamenBlog API. AdamenBlog API is a Go application for managing a blog.
https://github.com/paniccaaa/AdamenBlog 

## Stack:

- **cleanenv**: Environment configuration reader for Golang.
- **go-chi**: Router for building Go HTTP services.
- **PostgreSQL**: Relational database system used for data storage, hosted on Supabase.
- **slog**: Logging for Go, leveled logging.
- **Docker**
  
## Endpoints:

- **GET /posts**: Retrieve a list of all posts.
- **GET /posts/{id}**: Retrieve a post by its identifier.
- **POST /posts**: Create a new post.
- **PATCH /posts/{id}**: Update post information.
- **DELETE /posts/{id}**: Delete a post.

## Installation and Local Deployment

Before running the application, make sure Docker and Go are installed on your system.

1. **Clone the repository:**

    ```bash
    git clone https://github.com/paniccaaa/adamenblog-api.git
    cd adamenblog-api
    ```

2. **Create the .env file:**

    In the root directory of the project, create a .env file and specify the necessary environment variables, for example:

    ```makefile
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=admin
    DB_PASSWORD=secret
    DB_NAME=adamenblog
    ```
    
3. **Run the application locally**

    ```bash
    # Build and run the application locally
    make run
    ```

    Alternatively, to run the compiled binary file:

    ```bash
    # Build the binary file and execute it
    make buildapp && make run_bin
    ```

    The application will be available at `localhost:8080`.

## Docker

1. Dockerfile
   ```bash
   # Build the Docker image
   docker build -t adamenblog-api .

   # Run the Docker container
   docker run -d -p 8085:8080 --name adamenblog-api adamenblog-api
   ```
2. Docker Hub
   ```bash
   docker pull paniccaaa/adamenblog-api
   docker run -d -p 8085:8080 --name adamenblog-api paniccaaa/adamenblog-api
   ```
The application will be available at `localhost:8085`.

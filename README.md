### Configuration

1.  **Configure your environment:**

2.  **Build the Docker images:**

    ```bash
    docker compose build
    ```

3.  **Start the Docker containers:**

    ```bash
    docker compose up -d
    ```

4.  **Run database migrations:**

    ```bash
    go run cmd/migrate/main.go -command=up
    ```

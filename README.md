# otel-prometheus-study

This project is a Go-based RESTful API that serves as a study in implementing OpenTelemetry and Prometheus for 
monitoring and observability. The application provides basic CRUD functionality for managing customers, products, 
stores, and stock.

# Project Structure

The project follows a clean architecture pattern, separating concerns into distinct layers:

- `cmd/api`: Contains the main application entry point.
  - `internal`: Contains the core application logic, divided into:
  - `app`: Application services layer.
  - `handler`: HTTP handlers and routing.
  - `domain`: Core domain entities and business logic.
  - `infra`: Infrastructure concerns, such as database connections and external service integrations.
- `pkg`: Shared libraries and helpers.

# Features

- **RESTful API**: Provides endpoints for managing customers, products, stores, and stock.
- **PostgreSQL Integration**: Uses a PostgreSQL database for data persistence.
  - **Prometheus Metrics**: Exposes a /metrics endpoint for Prometheus to scrape.
- **Structured Logging**: Implements structured logging for better log management and analysis.
- **Clean Architecture**: Follows a clean architecture pattern for maintainability and testability.

# Prerequisites

- Go 1.24.4 or higher
- Docker and Docker Compose
- PostgreSQL client

# Installation

1. Clone the repository:
    ```Bash
    git clone https://github.com/your-username/otel-prometheus-study.git
    cd otel-prometheus-study
    ```
2. Install dependencies:
    ```Bash
    go mod tidy
    ```

3. Set up the database:

The project uses a PostgreSQL database. You can use the provided `docker-compose.yml` file to start a pre-configured 
PostgreSQL instance:
    ```Bash
    docker-compose up -d
    ```

This will start a PostgreSQL container and expose it on port `5432`.

4. Run the database migrations:

The `internal/infra/postgres/sql/init.sql` file contains the database schema. You can use a PostgreSQL client to run 
this script and create the necessary tables.

## Running the Application

To run the application, you can use the following command:

```Bash
go run cmd/api/main.go
```

The API server will start on port `8000`.

## API Endpoints

The following API endpoints are available:

- Customers
    ```
    GET /customers: List all customers
    POST /customers: Create a new customer
    GET /customers/:id: Get a customer by ID
    ```
- Products
    ```
    GET /products: List all products
    POST /products: Create a new product
    GET /products/:id: Get a product by ID
    PUT /products/:id: Update a product by ID
    ```
- Stock
    ```
    GET /stocks: List all stock
    POST /stocks: Create a new stock item
    GET /stocks/:id: Get a stock item by ID
    PUT /stocks/:id: Update a stock item by ID
    ```
- Stores
    ```
    GET /stores: List all stores
    POST /stores: Create a new store
    GET /stores/:id: Get a store by ID
    PUT /stores/:id: Update a store by ID
    ```
- Monitoring

The application exposes a /metrics endpoint on port 2112 for Prometheus to scrape. To view the metrics, you can access 
http://localhost:2112/metrics in your browser.

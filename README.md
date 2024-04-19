# Crypto Price Tracker

This project is a Go application that connects to a WebSocket server, fetches cryptocurrency prices, stores them in a PostgreSQL database, and leverages Debezium to track changes and send them to a RabbitMQ queue. The application is designed to be run using Docker Compose.

## Features

- Connects to a WebSocket server to receive real-time cryptocurrency price updates
- Stores cryptocurrency prices in a PostgreSQL database
- Includes a migrator module to create the necessary database table if it doesn't exist
- Updates existing records or creates new ones when new price data is received
- Utilizes Debezium to capture changes made to the database table
- Sends captured changes to a RabbitMQ queue for further processing

## Prerequisites

- Docker
- Docker Compose

## Installation

1. Clone the repository:<br>
```https://github.com/rmmbdev/crypto-price-tracker.git```
2. Navigate to the project directory:<br>
```cd crypto-price-tracker```

# Usage

Build and start the application using Docker Compose:<br>
```docker-compose up -d```

This command will build the necessary Docker images and start the containers for the application, PostgreSQL database, Debezium, and RabbitMQ.

2. The application will connect to the WebSocket server and start receiving cryptocurrency price updates.
3. The prices will be stored in the PostgreSQL database, and Debezium will capture any changes made to the database table.
4. Captured changes will be sent to the configured RabbitMQ queue for further processing.

## Stopping the Application
To stop the application and its associated services, run:<br>
```docker-compose down```

## Future Work / TODOs

- Implement processing of messages in the RabbitMQ queue
- Improve error handling and logging
- Implement unit tests and integration tests
- Add monitoring and alerting capabilities

## License
This project is licensed under the [MIT License](LICENSE).

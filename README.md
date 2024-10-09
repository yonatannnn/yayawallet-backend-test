YaYa Wallet Webhook Processor
# Overview

This project implements a webhook endpoint for YaYa Wallet. The webhook receives notifications of transactions, verifies their authenticity using HMAC signatures, and processes the payload. The solution is designed following a layered architecture with controllers, use cases, services, and repositories.

## Features

- HMAC (SHA256) Signature Verification
- Replay Protection (5-minute window)
- MongoDB repository for saving payloads
- Unit tests with mocks using Testify

## Project Structure

- **controller**: Handles incoming HTTP requests.
- **usecases**: Contains business logic for processing webhooks.
- **services**: Implements HMAC signature generation and verification.
- **repository**: MongoDB repository for saving payloads.
- **models**: Contains data models and interfaces.

## Setup and Usage

### Requirements

- Golang 1.19+
- Gin Framework
- MongoDB

### Steps to Run

Clone the repository:

```bash
git clone https://github.com/yonatannnn/yayawallet-backend-test.git
cd yourproject
```

Install dependencies:

```bash
go mod tidy
```

Set the environment variables:

```bash
export SECRET_KEY=your_secret_key
export PORT=your_port
export MONGODB_URL=your_mongodb_url
```

Run the application:

```bash
go run main.go
```

Use the following cURL command to test the webhook endpoint:

```bash
curl -X POST http://localhost:$PORT/webhook \
-H "Content-Type: application/json" \
-H "YAYA-SIGNATURE: valid_signature" \
-d '{"id":"12345","amount":1000,"currency":"USD","created_at_time":1625097600,"timestamp":1625097600,"cause":"Payment","full_name":"John Doe","account_name":"john.doe@example.com","invoice_url":"http://example.com/invoice/12345"}'
```

### Running Tests

To run the tests:

```bash
go test ./controller/
go test ./services/
```

## Assumptions

- The secret key for verifying the HMAC signature is provided through an environment variable named `SECRET_KEY`.
- The payload timestamp must not exceed a 5-minute (300 seconds) delay to be processed, otherwise, it's considered too old.

## Key Features

- **Signature Verification**: The solution uses HMAC with SHA256 to verify the payload's authenticity.
- **Time Validation**: Ensures that the webhook event is recent (within 5 minutes).
- **Error Handling**: Handles various types of errors such as invalid JSON, signature mismatch, or internal errors.

## Testing

- **Unit Tests**: Unit tests for the controller using the `testify` library and mocks ensure that each component is tested in isolation.
- **Mocking**: A mock version of the `WebhookUseCase` is used to simulate different responses during testing.
- **Gin Framework**: The tests utilize the Gin framework's `httptest` utilities to mock HTTP requests and responses.

### How to Test

#### Run Tests

Use the `go test ./controller/` command to run the test suite.  

#### Test Webhook Endpoint

Make a POST request to `/webhook` with a valid JSON body and a `YAYA-SIGNATURE` header containing a valid HMAC signature.

Example request:

```bash
curl -X POST http://localhost:$PORT/webhook \
-H "Content-Type: application/json" \
-H "YAYA-SIGNATURE: valid_signature" \
-d '{"id":"12345","amount":1000,"currency":"USD","created_at_time":1625097600,"timestamp":1625097600,"cause":"Payment","full_name":"John Doe","account_name":"john.doe@example.com","invoice_url":"http://example.com/invoice/12345"}'
```

## Project Description

### Project Structure

This solution implements a webhook processing service for YaYa Wallet. It follows a layered architecture to ensure separation of concerns and improve maintainability.

### Technologies Used

- **Golang**: Chosen for its performance, strong typing, and simplicity.
- **Gin Framework**: Used for handling HTTP requests and responses.
- **Testify**: A testing framework used for structuring unit tests and mocks.
- **HMAC (SHA-256)**: For verifying the authenticity of webhook payloads.
- **MongoDB**: For storing webhook payloads.

### Approach

#### Webhook Handler (Controller)

- Listens for incoming HTTP POST requests.
- Parses the request body and validates the JSON payload.
- Extracts the `YAYA-SIGNATURE` header and passes the payload and signature to the use case layer for further processing.
- Returns a 200 status code after validation but before saving the payload.

#### Use Case Layer

- This layer contains the core business logic.
- It first verifies the signature using the HMAC method to ensure the webhook payload is authentic.
- It checks whether the payload is within the allowed time window (i.e., the event is not older than 5 minutes).
- If valid, it saves the payload using the repository.

#### Service Layer

- Handles cryptographic operations such as creating and verifying HMAC signatures.
- Constructs the payload in a specific format and computes the HMAC signature to match it with the received one.
- Provides methods for saving the payload to the repository.

#### Repository Layer

- Implements saving the payload to MongoDB.

### Signature Verification

The payload is signed using HMAC with SHA256 to ensure the integrity and authenticity of the data. The service recreates the signed payload from the webhook body and compares the HMAC with the received signature.

### Replay Protection

The system checks the webhook’s timestamp and rejects requests older than 5 minutes to avoid replay attacks.

### Running the Project

1. Clone the project from GitHub.
2. Set the environment variables:

    ```bash
    export SECRET_KEY=your_secret_key
    export PORT=your_port
    export MONGODB_URL=your_mongodb_url
    ```

3. Run the server:

    ```bash
    go run main.go
    ```

4. Use the provided endpoint `/webhook` to handle webhook requests.

## License

This project is licensed under the MIT License.

_YaYa Wallet Webhook Processor
# Overview

This project implements a webhook endpoint for YaYa Wallet. The webhook receives notifications of transactions, verifies their authenticity using HMAC signatures, and processes the payload. The solution is designed following a layered architecture with controllers, use cases, services, and repositories.

## Features

- HMAC (SHA256) Signature Verification
- Replay Protection (5-minute window)
- Mock repository for logging payloads
- Unit tests with mocks using Testify

## Project Structure

- **controller**: Handles incoming HTTP requests.
- **usecases**: Contains business logic for processing webhooks.
- **services**: Implements HMAC signature generation and verification.
- **repository**: Mock repository for saving payloads.
- **models**: Contains data models and interfaces.

## Setup and Usage

### Requirements

- Golang 1.19+
- Gin Framework

### Steps to Run

Clone the repository:

```bash
git clone https://github.com/yourusername/yourproject.git
cd yourproject
```

Install dependencies:

```bash
go mod tidy
```

Set the `SECRET_KEY` environment variable:

```bash
export SECRET_KEY=your_secret_key
```

Run the application:

```bash
go run main.go
```

Use the following cURL command to test the webhook endpoint:

```bash
curl -X POST http://localhost:8080/webhook \
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

- The payload will have a valid timestamp within a 5-minute window.
- The HMAC signature is provided using the `SECRET_KEY` environment variable.

## Testing Strategy

Unit tests are written using the Testify framework. Mocks are used to simulate different service behaviors, ensuring that the system handles various scenarios such as successful processing, invalid signatures, and internal server errors.

## License

This project is licensed under the MIT License.

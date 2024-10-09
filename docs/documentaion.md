Webhook Processing Solution
# Overview

This project is a solution to YaYa Wallet's coding test. It involves building a webhook endpoint using Golang that receives and processes webhook notifications. The endpoint verifies the webhook signature, processes the payload, and saves the data using a repository pattern.

The solution has been structured into the following layers:

- **Controller**: Manages the HTTP request and sends responses.
- **Use Case**: Handles business logic, including verifying signatures and saving the payload.
- **Service**: Provides methods for signature creation and verification, and interacts with the repository.
- **Repository**: Handles the actual saving of data (currently a mock that logs the data).

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
curl -X POST http://localhost:8080/webhook \
-H "Content-Type: application/json" \
-H "YAYA-SIGNATURE: valid_signature" \
-d '{"id":"12345","amount":1000,"currency":"USD","created_at_time":1625097600,"timestamp":1625097600,"cause":"Payment","full_name":"John Doe","account_name":"john.doe@example.com","invoice_url":"http://example.com/invoice/12345"}'
```

## Project Description (for project_description.txt)

### Project Structure

This solution implements a webhook processing service for YaYa Wallet. It follows a layered architecture to ensure separation of concerns and improve maintainability.

### Technologies Used

- **Golang**: Chosen for its performance, strong typing, and simplicity.
- **Gin Framework**: Used for handling HTTP requests and responses.
- **Testify**: A testing framework used for structuring unit tests and mocks.
- **HMAC (SHA-256)**: For verifying the authenticity of webhook payloads.

### Approach

#### Webhook Handler (Controller)

- Listens for incoming HTTP POST requests.
- Parses the request body and validates the JSON payload.
- Extracts the `YAYA-SIGNATURE` header and passes the payload and signature to the use case layer for further processing.

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

- Currently implemented as a mock that logs the received payload. In a real application, this would be responsible for storing the payload in a database or persistent storage.

### Signature Verification

The payload is signed using HMAC with SHA256 to ensure the integrity and authenticity of the data. The service recreates the signed payload from the webhook body and compares the HMAC with the received signature.

### Replay Protection

The system checks the webhook’s timestamp and rejects requests older than 5 minutes to avoid replay attacks.

### Running the Project

1. Clone the project from GitHub.
2. Set the `SECRET_KEY` environment variable:

    ```bash
    export SECRET_KEY=your_secret_key
    ```

3. Run the server:

    ```bash
    go run main.go
    ```

4. Use the provided endpoint `/webhook` to handle webhook requests.
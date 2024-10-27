# Error Reporter

A simple Go application that captures error reports and sends alert emails using the SendGrid API. This project demonstrates how to handle error reporting, logging, and sending notifications via email.

## Features

- Capture error reports via HTTP requests.
- Send alert emails to specified recipients using SendGrid.
- Load configuration from a JSON file and environment variables.
- Log error reports to a file for future reference.
- Securely store sensitive information like API keys using encryption.

## Prerequisites

- Go (version 1.16 or higher)
- SendGrid account (for sending emails)
- A valid SendGrid API key

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/oghenevwede/Twilio-Articles.git
   cd error-reporter
   ```

2. **Install dependencies**:

   Make sure you have Go modules enabled and run:

   ```bash
   go mod tidy
   ```

3. **Set up your environment**:

   Create a `.env` file in the root directory of the project and add your SendGrid API key:

   ```plaintext
   SENDGRID_API_KEY=your_sendgrid_api_key_here
   ```

4. **Configure email settings**:

   Update the `config/config.json` file with your team emails and default email:

   ```json
   {
       "team_emails": {
           "TeamA": "test1@gmail.com",
           "TeamB": "test2@gmail.com",
           "TeamC": "test3@gmail.com"
       },
       "default_email": "test2@gmail.com",
       "sendgrid_api_key": "your_encrypted_api_key_here"
   }
   ```

   Use the provided encryption script to encrypt your SendGrid API key before placing it in the `config.json`.

5. **Run the application**:

   Start the server:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

## Usage

To report an error, send a POST request to the `/errors` endpoint with a JSON payload. For example:

json
{
"application": "MyApp",
"error_type": "database",
"severity": "high",
"message": "Unable to connect to the database.",
"context": {
"user_id": 123,
"transaction_id": "abc-123"
},
"notify": ["TeamA"]
}

### Example using `curl`

bash
curl -X POST http://localhost:8080/errors \
-H "Content-Type: application/json" \
-d '{
"application": "MyApp",
"error_type": "database",
"severity": "high",
"message": "Unable to connect to the database.",
"context": {
"user_id": 123,
"transaction_id": "abc-123"
},
"notify": ["TeamA"]
}'


## Logging

Error reports are logged to a file named `error.log` in the root directory of the project.

## Security

- Ensure that your `.env` and `config.json` files are not committed to version control. Add them to your `.gitignore` file.
- Use encryption for sensitive information like API keys.



## Acknowledgments

- [SendGrid](https://sendgrid.com/) for providing the email service.
- [Go](https://golang.org/) for the programming language.
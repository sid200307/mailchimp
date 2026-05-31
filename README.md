# Mailchimp Clone in Go 🚀

A high-performance bulk email dispatcher built with Go, featuring worker pools, HTML email templates, SMTP integration, concurrent email processing, and personalized email campaigns.

## Features

* Concurrent email sending using Goroutines
* Worker Pool Architecture
* CSV-based recipient loading
* Personalized email templates
* HTML Email Support
* SMTP Integration
* Mailpit Testing Support
* Gmail SMTP Support
* Error Handling and Retry Foundation
* Scalable Design for Bulk Campaigns

## Tech Stack

* Go (Golang)
* Goroutines
* Channels
* WaitGroups
* HTML Templates
* SMTP
* Mailpit
* Gmail SMTP

---

## Project Structure

```bash
mailchimp/
│
├── main.go
├── emailworker.go
├── csvloader.go
├── template.go
├── email.tmpl
├── users_1000.csv
└── README.md
```

---

## How It Works

1. Load recipients from CSV.
2. Push recipients into a channel.
3. Worker pool consumes recipients concurrently.
4. Generate personalized HTML emails.
5. Send emails through SMTP.
6. Log successes and failures.

---

## Worker Pool Architecture

```text
CSV File
    │
    ▼
Recipient Channel
    │
    ▼
 ┌─────────────┐
 │ Worker 1    │
 │ Worker 2    │
 │ Worker 3    │
 │ ...         │
 │ Worker N    │
 └─────────────┘
    │
    ▼
 SMTP Server
```

---

## Email Template Example

```html
From: your-email@gmail.com
To: {{.Email}}
Subject: Hello, {{.Name}}
MIME-Version: 1.0
Content-Type: text/html; charset=UTF-8

<html>
<body>
    <h1>Hello {{.Name}}</h1>
    <p>Welcome to our email campaign.</p>
</body>
</html>
```

---

## Running With Mailpit

### Start Mailpit

```bash
docker run -d --name mailpit -p 1025:1025 -p 8025:8025 axllent/mailpit
```

### Mailpit Dashboard

```text
http://localhost:8025
```

SMTP Configuration:

```go
smtpHost := "localhost"
smtpPort := "1025"
```

---

## Running With Gmail SMTP

Enable:

* 2-Step Verification
* Google App Passwords

SMTP Configuration:

```go
smtpHost := "smtp.gmail.com"
smtpPort := "587"

auth := smtp.PlainAuth(
    "",
    "your-email@gmail.com",
    "your-app-password",
    smtpHost,
)
```

---

## Example CSV

```csv
Name,Email
John,john@example.com
Alice,alice@example.com
Bob,bob@example.com
```

---

## Future Improvements

* PostgreSQL Integration
* Campaign Management
* Open Tracking
* Click Tracking
* Retry Queue (DLQ)
* Rate Limiting
* REST API
* Dashboard
* Email Analytics
* SendGrid Integration
* AWS SES Integration

---

## Learning Outcomes

This project demonstrates:

* Goroutines
* Channels
* Concurrency Patterns
* Worker Pools
* SMTP Communication
* HTML Templates
* Bulk Email Processing
* Scalable Backend Design

---

## Author

Siddharth Rajesh  Doshi

Python Developer | Golang Backend Developer

GitHub:
https://github.com/sid200307

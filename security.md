# Security Policy

## Supported Versions

Use this section to tell people about which versions of your project are currently being supported with security updates.

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of our software seriously. If you believe you have found a security vulnerability in the **Go Gin Doc** project, please report it to us as described below.

### How to Report

Please do **not** report security vulnerabilities through public GitHub issues.

Instead, please report them via email to `security@walonfoundation.org` (replace with actual email if available, or generic placeholder).

You should include:
- A description of the vulnerability.
- Steps to reproduce the issue.
- Any relevant logs or screenshots.

### Response Timeline

We will acknowledge receipt of your report within 48 hours and will strive to send you a more detailed response within 7 days indicating the next steps in handling your report.

### Disclosure Policy

We ask that you do not disclose the vulnerability to the public until we have had the opportunity to verify and fix the issue. We will credit you for your discovery in our release notes once the issue is resolved.

## Security Features

This project implements the following security measures:

-   **JWT Authentication**: All protected routes require a valid JSON Web Token.
-   **Password Hashing**: User passwords are hashed using `bcrypt` before storage (implied by `golang.org/x/crypto`).
-   **Environment Variables**: Sensitive configuration (secrets, DB credentials) is managed via `.env` files and not hardcoded.

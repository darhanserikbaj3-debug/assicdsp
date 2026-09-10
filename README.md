# Builder Pattern — Email Service (Go)

## 1. Overview
This project demonstrates the **Builder Design Pattern** in Go. 
The selected domain is an **Email Construction Service**.

Using the same construction process managed by the `EmailDirector`, the system builds two different representations:
1. **Email Object (`*Email`)** — An immutable struct representing a structured email object with internal fields and string representation.
2. **Email Preview (`string`)** — A text-formatted preview representation suitable for logs or UI display.

---

## 2. Class / Package Structure

The project is located in the `builder` package and consists of:

- **`Email` (`email.go`)**: The product struct. Fields are unexported (`sender`, `recipient`, `subject`, `body`) to ensure immutability.
- **`EmailBuilder` (`builder.go`)**: Interface declaring construction steps (`SetSender`, `SetRecipient`, `SetSubject`, `SetBody`).
- **`ConcreteEmailBuilder` (`email.builder.go`)**: Concrete builder that constructs a structured `Email` object with validation in `GetResult()`.
- **`EmailPreviewBuilder` (`preview_builder.go`)**: Concrete builder that constructs a formatted string representation of the email.
- **`EmailDirector` (`director.go`)**: Director managing ready-to-use configurations (`MakeWelcomeEmail`, `MakePasswordResetEmail`). It depends strictly on the `EmailBuilder` interface.
- **`main.go`**: Client entry point demonstrating how the director generates both representations.

---

## 3. Prerequisites
- Go 1.18+ installed on your system.

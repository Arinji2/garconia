# Garconia Landing Page Backend API

This is the **backend API** for the Garconia landing page. It provides essential functionality for managing user sign-ups, sending emails, and ensuring system health.

---

## Tech Stack

- **Golang**: The primary language for the backend API.
- **SQLite**: A lightweight database solution for persistent storage.
- **Resend**: A service used for sending emails.

---

## API Routes

The following routes are available:

1. **`/health` (GET)**

   - Acts as a health check to verify that the API is running.

2. **`/api/add` (POST)**

   - Adds a new user to the database.
   - Sends a verification email to the user.

3. **`/api/verify` (POST)**
   - Verifies a user.
   - Sends a confirmation email upon successful verification.

---

## Environment Variables

The backend requires the following environment variables to be configured:

- **`RESEND_API_KEY`**  
  Your API key for the Resend email service.

- **`WEBHOOK_URL`**  
  A Discord webhook URL to receive real-time notifications of new user sign-ups.

- **`FRONTEND_URL`**  
  The URL of the frontend application. This is used for constructing links or cross-origin requests.

Example:  
If your frontend is hosted at `https://frontend.com`, you should set:

```plaintext
FRONTEND_URL=https://frontend.com
```

# Database Folder for Garconia Landing API

This folder contains the **database migrations file** for the Garconia landing API.

---

## Contents

- **`migrations.sql`**:  
  This file contains all the necessary SQL commands for setting up and migrating the database.

---

## Important Notes

- **DO NOT EDIT OR ADD TO THIS FILE**:  
  The `migrations.sql` file is managed entirely by the API. Any manual changes may cause unexpected behavior or errors.

- The API will automatically handle database setup:
  - If the database is missing, the API will create a new one using `migrations.sql`.
  - Any errors during migration are gracefully handled by the API.

---

With this setup, the database is fully automated and requires minimal maintenance.

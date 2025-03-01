# Garconia Law Bot

This is the **official bot** for handling Garconia's legal system, making it easy to access constitutional laws and amendments directly within Discord.

---

## Tech Stack

- **Go**: A fast and efficient language for backend services.
- **PocketBase**: A lightweight backend for managing legal data.
- **Discord API**: For seamless integration with Discord servers.

---

## Hosting

The bot runs on a **dedicated server**, ensuring reliable uptime and quick responses.

---

## Environment Variables

The bot requires the following environment variables to be configured:

```env
TOKEN=(Bot Token)
GUILD_ID=(ID of the Guild which the bot is in)
ADMIN_EMAIL=(Super user email of the PocketBase Instance)
ADMIN_PASSWORD=(Super user password of the PocketBase Instance)
BASE_DOMAIN=(Base domain of the PocketBase Instance)
ALLOWED_ROLES=(Comma-separated list of allowed role IDs for the data-refresh command)
ALLOWED_CHANNELS=(Comma-separated list of allowed channel IDs for the commands to be run in)
```

---

## Commands

### Refresh Data

- **Command:** `/refresh-data`
- **Description:** Refreshes the bot’s stored legal data.

### Get Clauses

- **Command:** `/get-clauses`
- **Description:** Retrieves specific clauses from the constitution.
- **Options:**
  - `article-number` (Required): Article number of the clause.
  - `clause-number` (Optional): Clause number (autocomplete enabled).

### Get Articles

- **Command:** `/get-articles`
- **Description:** Retrieves articles from the constitution.
- **Options:**
  - `article-number` (Optional): Article number of the clause (autocomplete enabled).

### Get Amendments

- **Command:** `/get-amendments`
- **Description:** Retrieves constitutional amendments.
- **Options:**
  - `article-number` (Required): Article number of the constitution.
  - `clause-number` (Required): Clause number of the article.
  - `amendment-number` (Required): Amendment number of the clause.

---

## Built by Arinji

This project is proudly built as part of the **Garconia Monolithic Infrastructure** by [Arinji](https://www.arinji.com/). Check out my website for more cool projects!

# 📊 FISCUS - Usage Logging Backend for AI Visual Assistance System

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

**FISCUS** is a lightweight and modular backend logging service designed for the **EYEPATH** AI visual assistance system.  
It collects, stores, and serves user-specific session data such as inference durations and photo counts.  
Built in **Go**, FISCUS supports both SQLite and MySQL databases and communicates with the core system via REST API.

---

## 🌟 Project Overview

**EYEPATH** is a real-time collision prediction service tailored for visually impaired users. It combines mobile AI inference and haptic feedback to enhance user mobility.  
FISCUS serves as the **Usage Logging Service**, responsible for:

- Tracking inference session metadata (start time, end time, photo count)
- Providing monthly statistics for end-users
- Integrating with the Argus (AI) and Tablula (auth) modules
- Offering REST APIs to store and retrieve logs

---

## 🏗️ Architecture & Ecosystem

FISCUS is part of a broader distributed architecture:

- `EYEPATH-APP`: Android mobile application
- `ARGUS`: Real-time AI inference server (WebSocket)
- `TABLULA`: Authentication and user profile server
- `FISCUS`: Logging and usage metrics server

Each component operates independently, connected via token-secured HTTP/WebSocket communication.

---

## 🧰 Tech Stack

- **Language**: Go 1.24.4
- **Framework**: Gin (HTTP server)
- **Database**: SQLite3 (default) or MySQL (optional)
- **Security**: JWT access token required in headers
- **Logging**: go-logging structured formatter

---

## 📂 Repository Structure

```
fiscus/
├── cmd/
│   ├── main.go          # Entry point - starts Gin server
│   ├── api.go           # API route registration & handlers
│   ├── driver.go        # DB initialization and table setup
│   └── usage_log.go     # Core DB logic (insert/select)
├── scripts/
│   └── init_db.sh       # MySQL schema and user setup script
├── go.mod / go.sum      # Go module dependencies
└── README.md
```

---

## 🚀 API Specification

### 1️⃣ POST `/logs` - Save Usage Log

Used by the Argus AI server to record session data when a session ends.

#### 📤 Request Format

```http
POST /logs
Content-Type: application/json
```

```json
{
  "user_id": "andylim1022",
  "start_time": "2025-06-10 13:00:00",
  "end_time": "2025-06-10 13:30:00",
  "photos": 5
}
```

| Field       | Type     | Description                |
|-------------|----------|----------------------------|
| `user_id`   | string   | Unique user identifier     |
| `start_time`| datetime | Session start (UTC)        |
| `end_time`  | datetime | Session end (UTC)          |
| `photos`    | integer  | Number of captured images  |

#### 📥 Response (Success)

```json
{ "message": "Log saved successfully" }
```

---

### 2️⃣ GET `/usage/:userId` - Monthly Summary

Returns total minutes and photos used by a user in a given month.

#### 📤 Request Format

```http
GET /usage/andylim1022?year=2025&month=6
```

#### 📥 Response (Success)

```json
{
  "user_id": "andylim1022",
  "year": 2025,
  "month": 6,
  "used_minutes": 120,
  "photo_count": 423423
}
```

---

## 🔐 Authentication

FISCUS requires JWT authentication for all endpoints.

- **Header Format**:  
  `Authorization: Bearer <access_token>`

- Tokens are issued and validated by the **Tablula** module.

---

## 🛠️ How to Run

### 🔧 1. Install Dependencies

```bash
go mod tidy
```

### 🧪 2. Run with SQLite (default)

```bash
go run cmd/main.go
```

This uses a local `usage_log.db` file and creates the table automatically.

### 🐬 3. MySQL Setup (Optional)

To use MySQL:

- Edit `scripts/init_db.sh` and run:
  ```bash
  ./scripts/init_db.sh
  ```

- Then adjust connection logic in `cmd/driver.go`

---

## 🧠 How It Works

- **Argus AI service** sends usage data via POST `/logs` after each WebSocket session.
- **Tablula** verifies token claims and user validity.
- **FISCUS** stores sessions in `usage_log` table, including start/end times and image counts.
- Clients (e.g., Android app) call `/usage/:userId` to visualize monthly data.

---

## 🧪 Example Log Record

```
user_id: andylim1022
start_time: 2025-06-10 13:00:00
end_time:   2025-06-10 13:30:00
photos:     100223
```

Results in: `30 minutes used`, `100223 photos`

---

## 🧾 License

Licensed under [Apache License 2.0](LICENSE)

---

## 👥 Contributors

This project was developed as part of the **Embedded System course**  
at Dankook University, Department of Mobile System Engineering.

- Kim Woosung  
- Lee Youngjoo  
- Lim Seokbeom

🔗 GitHub: [EYEPATH-EMBEDDED/fiscus](https://github.com/EYEPATH-EMBEDDED/fiscus.git)

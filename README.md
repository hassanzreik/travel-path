# ✈️ Travel Path

`travel-path` is a Go web service that reconstructs a complete travel itinerary from a list of unordered flight ticket pairs.  
It uses the [Echo](https://echo.labstack.com/) framework and includes robust validation and full test coverage.

---

## 📁 Project Structure

```
travel-path/
├── api/
│   ├── handlers/
│   │   ├── travel_path.go           # HTTP handler logic
│   │   └── travel_path_test.go      # Tests for handler logic
│   └── routes.go                    # Registers routes
├── main.go                          # Entry point: starts Echo server
├── start.sh                         # Bash script to run tests then server
```

---

## 📦 Requirements

- Go 1.20+
- Bash (to run `start.sh`)
- curl (optional, for testing)

---

## 🚀 How to Run

1. Make the startup script executable:

   ```bash
   chmod +x start.sh
   ```

2. Start the server:

   ```bash
   ./start.sh
   ```

   > This runs all tests first. If tests pass, the Echo server will start on port `8081`.

---

## 🧪 Example API Request with curl

```bash
curl --location 'http://localhost:8081/travel-path' \
--header 'Content-Type: application/json' \
--data '[["LAX", "DXB"], ["JFK", "LAX"], ["SFO", "SJC"], ["DXB", "SFO"]]'
```

**Expected Response:**

```json
{
  "path": ["JFK", "LAX", "DXB", "SFO", "SJC"]
}
```

---

## ❌ Error Responses

| Scenario                     | Status | Error Message                             |
|-----------------------------|--------|-------------------------------------------|
| Empty input                 | 400    | `input is empty`                          |
| Malformed ticket            | 400    | `invalid ticket format...`                |
| Empty dep/dest              | 400    | `departure and destination must not...`   |
| Departure == Destination    | 400    | `departure cannot be the same as...`      |
| Duplicate departure         | 400    | `duplicate departure found`               |
| No valid starting point     | 400    | `no valid starting point found`           |
| Cycle in path               | 400    | `cycle detected in path`                  |
| Incomplete itinerary        | 400    | `incomplete itinerary path`               |

---

## 🧼 Run All Tests

To ensure everything is working correctly:

```bash
go test ./...
```

---

## 🧠 Notes

- Clean, testable, modular Go code
- Built with Echo framework
- Validates and reconstructs a travel itinerary
- Includes `start.sh` to enforce test success before running

---

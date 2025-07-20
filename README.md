# 📊 Binance L2 Order Book Collector

A lightweight Go service that connects to Binance's WebSocket API and streams **Level 2 (L2) depth data** (order book updates) into **Redis** for fast, structured access.  

Ideal for high-frequency trading research, real-time dashboards, or building your own quant infrastructure.

---

## 🚀 Features

- 🔌 Connects to Binance WebSocket stream (`@depth@100ms`)
- 🧠 Parses raw JSON into structured Go objects
- 🧾 Stores bid/ask data in Redis using symbol-based keys
- 🛡️ Includes `/health` HTTP endpoint for container health checks
- 🐳 Docker + Docker Compose ready

---

## 🧠 Financial Concepts

This collector captures **Level 2 depth data**, which includes:
- Multiple bid and ask levels (not just best price)
- Real-time changes to the order book
- Key for understanding market microstructure, liquidity, and price action

Example use cases:
- Market-making bots
- Order book imbalance indicators
- Quant backtesting and replay
- Live dashboards

---

## 🛠️ Project Structure

```bash
.
├── cmd/
│   └── main.go             # Starts the WebSocket listener + health server
├── internal/
│   ├── ws/                 # WebSocket client (connect, read, forward)
│   │   └── client.go
│   ├── parser/             # JSON parsing of Binance depth messages
│   │   └── depth.go
│   ├── handler/            # Business logic: writes to Redis
│   │   └── handler.go
│   └── store/              # Redis client wrapper
│       └── redis.go
├── Dockerfile
├── docker-compose.yml
└── README.md

## 🛠️ Phase 1: Project Setup and Architecture

### ✅ Tasks:
- [ ] Initialize Go module: `go mod init marketflow`
- [ ] Create project directory structure based on Hexagonal Architecture:
  ```
  marketflow/
  ├── cmd/
  │   └── marketflow/
  ├── internal/
  │   ├── app/
  │   ├── domain/
  │   ├── adapter/
  │   │   ├── http/
  │   │   ├── storage/
  │   │   ├── cache/
  │   │   └── exchange/
  │   ├── config/
  │   └── util/
  ├── test/
  ├── configs/
  ├── README.md
  └── go.mod
  ```
- [ ] Setup configuration loading (YAML or JSON)
- [ ] Add usage/help command line flag (`--help`)
- [ ] Implement graceful shutdown (SIGINT, SIGTERM)
- [ ] Setup structured logging using `log/slog`

---

## 🌐 Phase 2: Connect to One Exchange Source

### ✅ Tasks:
- [ ] Load and run one Docker image (e.g., `exchange1`)
- [ ] Implement TCP client to read data from exchange on port `40101`
- [ ] Parse incoming messages (ensure data includes: pair, price, timestamp, exchange)
- [ ] Output received data to log for verification
- [ ] Implement reconnection logic on failure (failover)

---

## 🔄 Phase 3: Concurrency – Fan-Out & Worker Pool

### ✅ Tasks:
- [ ] Implement Fan-Out pattern for one exchange
- [ ] Create a Worker Pool with 5 workers
- [ ] Define `MarketData` struct and processing logic
- [ ] Use channels to communicate between listener → workers → aggregator
- [ ] Implement batching logic for storing data periodically

---

## 🧪 Phase 4: Test Mode (Generator Pattern)

### ✅ Tasks:
- [ ] Create synthetic data generator (Test Mode)
- [ ] Match format of real exchange data
- [ ] Launch multiple generators (simulating 3 exchanges)
- [ ] Toggle between Test Mode and Live Mode via internal config
- [ ] Expose `/mode/test` and `/mode/live` HTTP endpoints

---

## 🚀 Phase 5: Redis Caching Integration

### ✅ Tasks:
- [ ] Connect to Redis
- [ ] Design keys like: `latest:{exchange}:{symbol}`
- [ ] Store recent prices per symbol per exchange
- [ ] Maintain data for at least 60 seconds
- [ ] Periodically clean old entries

---

## 🗃️ Phase 6: PostgreSQL Storage

### ✅ Tasks:
- [ ] Connect to PostgreSQL
- [ ] Create schema and table:
  ```sql
  CREATE TABLE market_data (
      pair_name TEXT,
      exchange TEXT,
      timestamp TIMESTAMP,
      average_price FLOAT,
      min_price FLOAT,
      max_price FLOAT
  );
  ```
- [ ] Every minute, compute average, min, max from Redis
- [ ] Batch insert to PostgreSQL

---

## 🔧 Phase 7: Expand to Multiple Exchanges (Fan-In Pattern)

### ✅ Tasks:
- [ ] Add listeners for all 3 exchange sources
- [ ] Implement fan-in to aggregate all workers' results
- [ ] Maintain isolation per exchange, but aggregate for combined views

---

## 🌐 Phase 8: API Development

### ✅ Tasks:

#### Price Data Endpoints
- [ ] `GET /prices/latest/{symbol}`
- [ ] `GET /prices/latest/{exchange}/{symbol}`
- [ ] `GET /prices/highest/{symbol}`
- [ ] `GET /prices/highest/{exchange}/{symbol}?period={duration}`
- [ ] `GET /prices/lowest/{symbol}`
- [ ] `GET /prices/lowest/{exchange}/{symbol}?period={duration}`
- [ ] `GET /prices/average/{symbol}`
- [ ] `GET /prices/average/{exchange}/{symbol}?period={duration}`

#### Mode Switchers
- [ ] `POST /mode/test`
- [ ] `POST /mode/live`

#### Health Check
- [ ] `GET /health`

---

## 📊 Phase 9: Monitoring, Logging, Error Handling

### ✅ Tasks:
- [ ] Add `Info`, `Warn`, `Error` logs with context
- [ ] Log Redis/Pg failures but continue running
- [ ] Include timestamps, pair names, exchange names in logs
- [ ] Return clear HTTP error messages (400, 500, etc.)

---

## 📦 Phase 10: Final Touches

### ✅ Tasks:
- [ ] Ensure formatting via `gofumpt`
- [ ] Add unit tests for core logic
- [ ] Validate graceful shutdown works
- [ ] Print usage on `--help`
- [ ] Check Redis and PostgreSQL fallbacks
- [ ] Optimize goroutine handling
- [ ] Cleanup unused data in Redis
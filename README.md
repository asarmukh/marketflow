# MarketFlow

## Overview

MarketFlow is a real-time cryptocurrency market data processing system built with Go that implements hexagonal architecture principles. The system processes concurrent data streams from multiple exchanges, provides caching with Redis, stores aggregated data in PostgreSQL, and exposes a RESTful API for querying price information.

## Architecture

This project follows **hexagonal architecture** (ports and adapters) with clear separation of concerns:

- **Domain Layer**: Core business logic and entities (`internal/core/domain/`)
- **Application Layer**: Use cases and service orchestration (`internal/core/services/`)
- **Ports**: Interface definitions (`internal/core/ports/`)
- **Adapters**: External integrations
  - **Inbound**: HTTP API and CLI (`internal/adapters/inbound/`)
  - **Outbound**: Database, cache, and exchange clients (`internal/adapters/outbound/`)

## Features

🏗️ **Hexagonal Architecture** - Clean separation of business logic from infrastructure
🔄 **Real-time Data Processing** - Concurrent data streams from multiple exchanges
⚡ **High-Performance Concurrency** - Fan-in/Fan-out patterns with worker pools
💾 **Dual Storage** - PostgreSQL for persistence, Redis for caching
🔄 **Mode Switching** - Live mode (real exchanges) and Test mode (synthetic data)
📊 **Data Aggregation** - Minute-level price statistics (avg/min/max)
🛡️ **Fault Tolerance** - Automatic reconnection and Redis fallback mechanisms
🌐 **RESTful API** - Comprehensive endpoints for price queries
🔍 **Health Monitoring** - System status and connectivity checks
📝 **Structured Logging** - Contextual logging with slog
🛑 **Graceful Shutdown** - Proper resource cleanup on termination

## Supported Trading Pairs

- BTCUSDT
- DOGEUSDT  
- TONUSDT
- SOLUSDT
- ETHUSDT

## Installation & Usage

### Prerequisites

- Go 1.22+
- Docker & Docker Compose
- PostgreSQL
- Redis

### Quick Start

```bash
# Clone the repository
git clone <repository-url>
cd marketflow

# Start dependencies
docker-compose up -d postgres redis

# Load exchange simulators (adjust for your architecture)
docker load -i data/exchange1_amd64.tar
docker load -i data/exchange2_amd64.tar
docker load -i data/exchange3_amd64.tar

# Run exchange simulators
docker run -p 40101:40101 --name exchange1 -d exchange1-amd64
docker run -p 40102:40102 --name exchange2 -d exchange2-amd64
docker run -p 40103:40103 --name exchange3 -d exchange3-amd64

# Build and run MarketFlow
go build -o marketflow .
./marketflow --port 8080
```

### CLI Usage

```bash
Usage:
  marketflow [--port <N>]
  marketflow --help

Options:
  --port N     Port number (default: 8080)
```

## API Endpoints

### Market Data API

**Latest Prices**
- `GET /prices/latest/{symbol}` – Latest price across all exchanges
- `GET /prices/latest/{exchange}/{symbol}` – Latest price from specific exchange

**Highest Prices**
- `GET /prices/highest/{symbol}` – Highest price across all exchanges  
- `GET /prices/highest/{exchange}/{symbol}` – Highest price from specific exchange
- `GET /prices/highest/{symbol}?period={duration}` – Highest price in time period
- `GET /prices/highest/{exchange}/{symbol}?period={duration}` – Highest price from exchange in period

**Lowest Prices**
- `GET /prices/lowest/{symbol}` – Lowest price across all exchanges
- `GET /prices/lowest/{exchange}/{symbol}` – Lowest price from specific exchange  
- `GET /prices/lowest/{symbol}?period={duration}` – Lowest price in time period
- `GET /prices/lowest/{exchange}/{symbol}?period={duration}` – Lowest price from exchange in period

**Average Prices**
- `GET /prices/average/{symbol}` – Average price across all exchanges
- `GET /prices/average/{exchange}/{symbol}` – Average price from specific exchange
- `GET /prices/average/{exchange}/{symbol}?period={duration}` – Average price from exchange in period

### Mode Management API

- `POST /mode/test` – Switch to Test Mode (synthetic data generation)
- `POST /mode/live` – Switch to Live Mode (real exchange data)
- `GET /mode` – Get current operating mode

### System Health API

- `GET /health` – System health status and connectivity checks

### Supported Time Periods

Use these duration formats in `period` query parameters:
- `1s`, `3s`, `5s`, `10s`, `30s` (seconds)
- `1m`, `3m`, `5m`, `10m`, `30m` (minutes)

### Error Response Format

All API errors return structured JSON responses:

```json
{
  "error": "Not Found",
  "message": "Price data not found: no recent data available",
  "code": 404,
  "timestamp": "2024-01-15T10:30:00Z"
}
```

## Configuration

The application reads configuration from `config.yaml`. Key settings include:

- **PostgreSQL**: Database connection details
- **Redis**: Cache connection and TTL settings  
- **Exchanges**: Connection details for live data sources
- **Aggregation**: Time windows and batch processing settings

## Concurrency Patterns

MarketFlow implements several Go concurrency patterns:

- **Fan-in**: Aggregates multiple exchange streams
- **Fan-out**: Distributes processing across worker pools
- **Worker Pools**: 5 workers per exchange for balanced processing
- **Generator**: Produces synthetic test data
- **Pipeline**: Orchestrates data flow through processing stages

## Data Flow

1. **Data Ingestion**: Connect to exchanges (live) or generate data (test)
2. **Processing**: Worker pools process price updates concurrently
3. **Caching**: Store recent prices in Redis for fast retrieval
4. **Aggregation**: Calculate minute-level statistics (avg/min/max)
5. **Persistence**: Batch write aggregated data to PostgreSQL
6. **API**: Serve real-time and historical price queries

## Development

### Project Structure

```
marketflow/
├── internal/
│   ├── core/                 # Business logic layer
│   │   ├── domain/          # Entities and value objects
│   │   ├── ports/           # Interface definitions  
│   │   └── services/        # Application services
│   ├── adapters/            # External adapters
│   │   ├── inbound/         # HTTP API, CLI
│   │   └── outbound/        # Database, cache, exchanges
│   ├── infrastructure/      # Cross-cutting concerns
│   └── config/             # Configuration management
├── pkg/                    # Shared packages
├── data/                   # Exchange Docker images
└── logs/                   # Application logs
```

### Contributing

1. Follow hexagonal architecture principles
2. Use dependency injection through interfaces
3. Write comprehensive tests for business logic
4. Ensure proper error handling and logging
5. Follow Go formatting standards (gofumpt)

## License

kbayazov and asarmukh

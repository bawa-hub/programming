# Concurrent Web Scraper

A production-ready web scraper built with Go that demonstrates advanced concurrency patterns.

## Features

- **Worker Pool**: Fixed number of workers processing scraping jobs
- **Rate Limiting**: Respectful scraping with configurable rate limits
- **Circuit Breaker**: Fault tolerance for handling service failures
- **Context Support**: Cancellation and timeout handling
- **Pipeline Processing**: Data flows through scraping → processing → storage
- **Pub/Sub Events**: Real-time event notifications
- **Graceful Shutdown**: Clean shutdown on interrupt signals

## Architecture

```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐
│   Job Queue │───▶│ Worker Pool  │───▶│   Results   │
└─────────────┘    └──────────────┘    └─────────────┘
                           │
                           ▼
┌─────────────┐    ┌──────────────┐    ┌─────────────┐
│   Events    │◀───│   Pub/Sub    │◀───│  Processor  │
└─────────────┘    └──────────────┘    └─────────────┘
```

## Components

### 1. Worker Pool (`worker_pool.go`)
- Manages 3 concurrent workers
- Processes jobs from the queue
- Handles result processing and data transformation

### 2. Rate Limiter (`rate_limiter.go`)
- Limits requests to 2 per second
- Prevents overwhelming target servers
- Token bucket algorithm implementation

### 3. Circuit Breaker (`circuit_breaker.go`)
- Opens circuit after 3 consecutive failures
- 30-second timeout before retry
- Prevents cascading failures

### 4. Web Scraper (`scraper.go`)
- HTTP client with 10-second timeout
- Title extraction from HTML
- Error handling and status reporting

### 5. Data Processor (`processor.go`)
- Word count calculation
- Keyword extraction
- Data structuring

### 6. Pub/Sub System (`pubsub.go`)
- Event-driven architecture
- Real-time notifications
- Decoupled components

## Usage

```bash
# Run the scraper
go run .

# The scraper will:
# 1. Start 3 workers
# 2. Process 5 test URLs
# 3. Apply rate limiting
# 4. Handle failures with circuit breaker
# 5. Process and display results
# 6. Shutdown gracefully on Ctrl+C
```

## Concurrency Patterns Demonstrated

1. **Goroutines**: Worker pool and event handlers
2. **Channels**: Job queue, results, and events
3. **Select**: Non-blocking operations and context handling
4. **Sync Package**: WaitGroup for coordination
5. **Context**: Cancellation and timeouts
6. **Advanced Patterns**: Worker pool, pipeline, pub/sub, rate limiting, circuit breaker

## Configuration

- **Workers**: 3 concurrent workers
- **Rate Limit**: 2 requests per second
- **Circuit Breaker**: 3 failures, 30-second timeout
- **HTTP Timeout**: 10 seconds per request
- **Job Timeout**: 30 seconds total

## Example Output

```
=== Concurrent Web Scraper ===
Added job: job-1
Added job: job-2
[Worker] 10:30:15: Worker 1 started job job-1
[Results] 10:30:16: Result processed for https://httpbin.org/html
[Data] 10:30:16: Data processed for https://httpbin.org/html

=== Processed Data ===
URL: https://httpbin.org/html
Title: Herman Melville - Moby Dick
Word Count: 1250
Keywords: [the, and, of, to, a]
Processed At: 10:30:16
====================
```

## Learning Objectives

This project demonstrates:
- Real-world application of Go concurrency
- Production-ready error handling
- Scalable architecture patterns
- Performance optimization techniques
- Clean code organization

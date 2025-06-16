# Gologger

**Gologger** is a real-time, concurrent log monitoring tool written in Go. It tails multiple log files concurrently, parses them for error and warning messages, and periodically reports aggregated statistics.

---

## 🔧 Features

- 📂 Monitor multiple log files concurrently  
- 🧵 Uses Goroutines and Channels for concurrency  
- 🔍 Detects `ERROR` and `WARN` log levels via regex  
- 📊 Aggregates error/warning counts per file  
- 🕒 Prints summaries every 10 seconds  
- ⛔ Graceful shutdown via Ctrl+C (SIGINT/SIGTERM)

---

## 📦 Architecture

```
┌────────────┐
│ Tailer N   │◄────┐
└────────────┘     │
       ▼           │
  (LogLine chan)   │
       ▼           │
   ┌────────┐      │
   │ Parser │◄─────┘
   └────────┘
       ▼
(ParsedLog chan)
       ▼
 ┌────────────┐
 │ Aggregator │
 └────────────┘
```

- **Tailer**: One goroutine per log file  
- **Parser**: Matches `ERROR` and `WARN` lines  
- **Aggregator**: Counts and logs summaries

---

## 🚀 Usage

### Build

```sh
make build
```

### Run

```sh
make run
```

Or manually:

```sh
./gologger -files="testdata/log1.log,testdata/log2.log"
```

### Simulate log input

In another terminal:

```sh
echo "ERROR Something broke" >> testdata/log1.log
echo "WARN Disk space low" >> testdata/log2.log
```

---

## 📁 Project Structure

```
gologger/
├── cmd/gologger/         # Main entry point
│   └── main.go
├── internal/
│   ├── tailer/           # Tails files concurrently
│   ├── parser/           # Matches log levels
│   └── aggregator/       # Aggregates and prints stats
├── pkg/types/            # Shared data types
├── testdata/             # Sample log files
├── Makefile              # Build and run commands
└── go.mod
```

---

## 🛑 Graceful Shutdown

Press `Ctrl+C`  
→ Stops tailers  
→ Flushes stats  
→ Exits cleanly

---

## 🧪 Future Improvements

- Log level customization  
- Timestamp parsing  
- JSON/HTML export  
- Web dashboard

---

## 🛠 Requirements

- Go 1.20+  
- Unix-like system (for file appending via shell)


Built for hands-on concurrency practice in Go.


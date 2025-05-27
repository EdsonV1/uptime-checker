# ⏱️ Uptime Checker

A lightweight CLI tool written in Go to check the availability of websites. It supports checking a single URL or multiple URLs from a file and can perform repeated checks over a specified interval and duration.

---

## 🚀 Features

- Check the HTTP status of a single URL or multiple URLs from a file.
- Repeat checks over a specified interval and duration.
- Simple CLI interface using standard Go libraries.
- Modular design for easy extension and testing.

---

## 🛠️ Technologies Used

- **Go**: The programming language used.
- **net/http**: To perform HTTP requests.
- **flag**: To parse command-line flags.
- **time**: For managing intervals and durations.

---

## 📦 Installation

1. **Clone the repository**

```bash
git clone https://github.com/EdsonV1/uptime-checker.git
cd uptime-checker
```

2. **Build the project**

```bash
go build -o uptime-checker
```

---

## ⚙️ Usage

### Check a single URL

```bash
./uptime-checker -url https://example.com
```

### Check multiple URLs from a file

```bash
./uptime-checker -file urls.txt
```

### Run the check repeatedly

```bash
./uptime-checker -url https://example.com -interval 10s -duration 1m
```

- `-interval`: How often to repeat the check (e.g., `10s`, `1m`).
- `-duration`: Total duration to keep checking (e.g., `1m`, `5m`).

### Help

```bash
./uptime-checker -help
```

---

## 💡 Example

Create a file `urls.txt`:

```
https://example.com
https://google.com
https://nonexistent.website
```

Run:

```bash
./uptime-checker -file urls.txt -interval 30s -duration 2m
```

---

## 📌 Future Enhancements

- Concurrency for faster checks
- Alert system (e.g., email, Slack)
- Logging to a file
- Docker support

---

## 📄 License

MIT License
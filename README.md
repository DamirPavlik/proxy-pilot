# ProxyPilot – Steering Your Traffic with Precision

ProxyPilot is a lightweight and flexible reverse proxy server built in Go. It allows you to route traffic to multiple backend services with ease, making it perfect for microservices architecture, load balancing, and API routing.

---

### Features
- **Dynamic Proxy Configuration**: Easily configure multiple routes and destination servers via a YAML configuration file.
- **Built-in Health Check**: `/ping` endpoint to ensure the server is running.
- **Simple and Fast**: Lightweight reverse proxy powered by Go's `net/http` package.
- **Customizable**: Modify or extend the routing and proxy logic for your specific needs.

---

### Usage

1. **Run Backend Services**:  
   Use Docker to spin up mock services:
   
   ```bash
   make run

2. **Start The Server**:  
   Start the go server
   
   ```bash
   go run cmd/main.go

3. **Ping Health Check:**:  
   Test the server is running:
   
   ```bash
   curl -I http://localhost:8080/ping

4. **Proxy Requests:**:  
   Forward requests to a backend service:
   
   ```bash
   curl -I http://localhost:8080/server1

5. **Stop Backend Services:**:  
   Use the following command to stop mock services:
   
   ```bash
   make stop

---

### Makefile Commands

| Command                 | Description                                      |
|-------------------------|--------------------------------------------------|
| `make run-proxy`        | Run the Go server                        |
| `make run`              | Start backend mock services via Docker           |
| `make stop`             | Stop all running backend mock services           |
| `make help`             | Display all available Makefile commands  

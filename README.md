# ProxyPilot – Steering Your Traffic with Precision

ProxyPilot is a lightweight and flexible reverse proxy server built in Go. It allows you to route traffic to multiple backend services with ease, making it perfect for microservices architecture, load balancing, and API routing.

---

## Features
- **Dynamic Proxy Configuration**: Easily configure multiple routes and destination servers via a YAML configuration file.
- **Built-in Health Check**: `/ping` endpoint to ensure the server is running.
- **Simple and Fast**: Lightweight reverse proxy powered by Go's `net/http` package.
- **Customizable**: Modify or extend the routing and proxy logic for your specific needs.

---

### Makefile Commands

| Command                 | Description                                      |
|-------------------------|--------------------------------------------------|
| `make run-proxy`        | Run the ProxyPilot server                        |
| `make run`              | Start backend mock services via Docker           |
| `make stop`             | Stop all running backend mock services           |
| `make help`             | Display all available Makefile commands  

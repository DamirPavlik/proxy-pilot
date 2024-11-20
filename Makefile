## run: start multiple httpbin containers
.PHONY: run
run:
	docker run --rm -d -p 9001:80 --name server1 kennethreitz/httpbin 
	docker run --rm -d -p 9002:80 --name server2 kennethreitz/httpbin 
	docker run --rm -d -p 9003:80 --name server3 kennethreitz/httpbin 
	docker run --rm -d -p 9004:80 --name server4 kennethreitz/httpbin 
	docker run --rm -d -p 9005:80 --name server5 kennethreitz/httpbin 

## stop: stop all running httpbin containers
.PHONY: stop
stop:
	docker stop server1
	docker stop server2
	docker stop server3
	docker stop server4
	docker stop server5

## run-proxy: runs the go proxy server
.PHONY: run-proxy
run-proxy:
	go run cmd/main.go

## help: displays this help message
.PHONY: help
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
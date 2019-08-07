run: build
	docker run -p 8080:8080 --rm ps1.fun

build:
	docker build -t ps1.fun -f ops/Dockerfile .

run: build
	docker run -p 8080:8080 --rm ps1.fun

build:
	docker build -t ps1.fun -f ops/Dockerfile .

try: build_try_image
	docker run -it --rm ps1.fun.try -c 'curl -s host.docker.internal:8080 | bash; bash'

build_try_image:
	docker build -t ps1.fun.try -f util/Dockerfile ./util

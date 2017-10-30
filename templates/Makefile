all: build test

build:
	@echo ">> Building image"
	docker build -t dockerizedphp/{{ .Name }} latest/

test:
	@echo ">> Run test"
	docker run --rm -ti dockerizedphp/{{ .Name }} --help 2>/dev/null; true

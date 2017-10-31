PROJECT = bookshelf
WORKDIR = /go/src/github.com/nicovogelaar/$(PROJECT)
PKGS = ./app ./bookshelf

all: glide-install builder build up

up:
	docker-compose up -d

down:
	docker-compose down

server:
	docker-compose up --build server

db:
	docker-compose up --build db

build:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		$(PROJECT)-builder \
		go build -ldflags '-w -s' -a -installsuffix cgo -o bin/server

lint:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		$(PROJECT)-builder \
		golint $(PKGS)

test:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		$(PROJECT)-builder \
		/bin/bash -c "golint $(PKGS); go test $(PKGS)"

test-travis:
	golint $(PKGS)
	go test $(PKGS)

builder:
	docker build -t $(PROJECT)-builder -f data/Dockerfile.build .

clean:
	rm -r vendor/
	docker-compose stop
	docker-compose rm -f
	docker-compose down --rmi all --volumes --remove-orphans

glide-install:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		instrumentisto/glide install

glide-get:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		instrumentisto/glide get $(pkg)

glide-init:
	docker run --rm -it \
		-v $(PWD):$(WORKDIR) \
		-w $(WORKDIR) \
		instrumentisto/glide init
PROJECT = bookshelf
WORKDIR = /go/src/github.com/nicovogelaar/$(PROJECT)
PKGS = ./app ./bookshelf

all: glide-install build-image-builder build start

start:
	docker-compose up -d

stop:
	docker-compose down

compose-server: build
	docker-compose build server
	docker-compose up server

compose-db:
	docker-compose build db
	docker-compose up db

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
		go test $(PKGS)

build-image-builder:
	docker build -t $(PROJECT)-builder -f data/Dockerfile.builder .

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
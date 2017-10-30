all: install dev-server

install:
	npm install

build:
	./node_modules/.bin/webpack

dev-server:
	./node_modules/.bin/webpack-dev-server

lint:
	./node_modules/.bin/eslint --ext .html,.js,.vue ./src
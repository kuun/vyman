all: build web

build:
	go build -ldflags "-linkmode 'external' -extldflags '-static'"


web:
	cd webapp; npm run build

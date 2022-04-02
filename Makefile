
build: example/example.go
	mkdir -p temp
	go build -o temp/example example/example.go

copy: build
	scp temp/example optiplex:/tmp

deploy: copy
	ssh optiplex /tmp/example -action deploy

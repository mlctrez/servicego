


build: example/example.go
	mkdir -p temp
	go build -o temp/example example/example.go

copy: build
	scp temp/example optiplex:/tmp

deploy: copy
	ssh optiplex /tmp/example -action deploy

tags:
	git tag -l

version:
	echo "$${VER:?re run with VER=v1.x.x}"
	git add .
	git commit -m "version $VER"
	git tag -a $VER -m "$VER"
	git push origin HEAD $VER

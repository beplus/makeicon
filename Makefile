VERSION=$(shell ./makeicon -v)
HARDWARE=$(shell uname -m)

build:
	go build

release: build
	mkdir release
	GOOS=linux go build -o release/makeicon
	cd release && tar -zcf makeicon_$(VERSION)_Linux_$(HARDWARE).tgz makeicon
	GOOS=darwin go build -o release/makeicon
	cd release && tar -zcf makeicon_$(VERSION)_Darwin_$(HARDWARE).tgz makeicon
	rm release/makeicon

clean:
	rm -rf release

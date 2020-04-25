
BINARIES = test
CMD_PATH = github.com/shrek/golearn/cmd
PKG_PATH = github.com/shrek/pkg

test:
	@mkdir -p bin
	go build -a -o bin/test $(CMD_PATH)/test

lint:
	golangci-lint run --verbose

unittest:
	go clean testcache; \
	go test -v $(PKG_PATH)/...

funtionaltest:
	echo "not implemented yet"

utcover:
	go test -cover $(PKG_PATH)/...

clean:
	@rm -rf bin/

.PHONY:
	all clean lint test

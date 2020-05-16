
BINARIES = test kelly
CMD_PATH = github.com/shrek/golearn/cmd
PKG_PATH = github.com/shrek/pkg

all: bandit

test:
	@mkdir -p bin
	go build -a -o bin/test $(CMD_PATH)/test

kelly:
	@mkdir -p bin
	go build -a -o bin/kelly $(CMD_PATH)/kelly

bandit:
	@mkdir -p bin
	go build -a -o bin/bandit $(CMD_PATH)/bandit

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

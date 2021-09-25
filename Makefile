BIN=app

all: build

build:
	go build -o $(BIN)

clean:
	rm -rf $(BIN)

.PHONEY: all build clean

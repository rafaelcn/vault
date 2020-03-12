CC=`which go`
SRC=main.go criptor.go hashes.go

run:
	$(CC) run $(SRC)

build:
	$(CC)  build -o vault $(SRC)
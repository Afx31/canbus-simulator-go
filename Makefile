BINARY_NAME=bin/canbus-simulator-go
BINARY_DIR=./

# Compiles executable
compile:
	go build -o ${BINARY_NAME} ${BINARY_DIR}


# Compiles executable and runs
rebuild: compile
	./${BINARY_NAME}


# Runs already compiled executable
run: compile
	./${BINARY_NAME}


# Cleans and removes executable
purge:
	go clean
	rm ${BINARY_NAME}
GO=CGO_ENABLED=0 GOOS=linux go
CLEAN=${GO} clean
TEST=${GO} test
BUILD=${GO} build

all: clean build test
clean:
	${CLEAN}
build:
	${BUILD} -o gve github.com/thunpin/gve/cmd
test:
	${TEST} -cover -v ./pkg/...

.PHONY: build
build:
	@mkdir -p dist
	go build -o ./dist/licence-check ./cmd/licenceCalculator
	@chmod +x ./dist/licence-check

.PHONY: run
run: build
	./dist/licence-check

.PHONY: clean
clean:
	@rm -rf dist

.PHONY: test
test:
	go test ./...

.PHONY: test-small
test-small: build
	IN_FILE="./tests/testdata/sample-small.csv" \
	./dist/licence-check

.PHONY: fmt
fmt:
	fmt ./...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
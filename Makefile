.PHONY: build
build:
	@mkdir dist
	go build -o ./dist/licence-check ./cmd/licenceCalculator

.PHONY: run
run: 
	./dist/licence-check

.PHONY: clean
clean:
	@rm -rf dist

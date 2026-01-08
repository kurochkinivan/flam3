COVERAGE_FILE_OUT ?= coverage.out
COVERAGE_FILE_HTML ?= coverage.html

TARGET ?= flam3
VERSION ?= v1.0.0

S1 ?= 2408048046953245368
S2 ?= 14740139351423643016

## build binary file
.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p bin
	@go build -ldflags="-X 'main.version=${VERSION}'" -o bin/${TARGET} ./cmd/${TARGET}

## test: run all tests with coverage
.PHONY: test
test: 
	@echo "▶ Запуск тестов..."
	@go test -coverpkg='gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/...' \
		--race \
		-count=1 \
		-coverprofile='$(COVERAGE_FILE_OUT)' ./...
	@echo "▶ Результаты покрытия:"
	@go tool cover -func='$(COVERAGE_FILE_OUT)' | grep ^total | tr -s '\t'
	@go tool cover -html='$(COVERAGE_FILE_OUT)' -o ${COVERAGE_FILE_HTML}
	@echo "✓ HTML отчет сохранен: ${COVERAGE_FILE_HTML}"

## flam3-json: run flam3 with config.json
.PHONY: flam3-json
flam3-json:
	@go run cmd/flam3/main.go --config='input/config.json'

## flam3-cli: run flam3 with cli args
.PHONY: flam3-cli 
flam3-cli:
	@go run cmd/flam3/main.go \
  		--width 1920 \
  		--height 1080 \
  		--seed 2.1324512 \
  		-i 2500 \
  		-o output/result.png \
  		-t 8 \
  		-f swirl:10.0,horseshoe:0.7 \
  		-ap 1.0,1.0,1.0,1.0,1.0,1.0/0.3,1.0,-0.2,0.4,1.0,1.0

.PHONY: flam3-random
flam3-random:
	@go run cmd/flam3-rand/main.go --number=200 --output-dir=output

.PHONY: flam3-random-reproduce
flam3-random-reproduce:
	@go run ./cmd/flam3-rand/main.go --seed1=${S1} --seed2=${S2} --number=1 --output-dir=output/reproduce
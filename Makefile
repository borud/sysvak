all: vet lint test build

build: sysvak

sysvak:
	@cd cmd/sysvak && go build -o ../../bin/sysvak

test:
	@go test ./...

vet:
	@go vet ./...

lint:
	@revive ./...

kommuner:
	@curl -s https://data.ssb.no/api/klass/v1/versions/1160.csv?language=nb -o pkg/sysvak/kommuner.csv
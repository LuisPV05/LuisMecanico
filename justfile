set windows-shell := ["sh", "-cu"]

check:
    gofmt -e . > /dev/null

test:
    go test ./...

calidad:
    golangci-lint run
    find internal -name '*.go' | xargs wc -l | awk '$2 != "total" && $1 > 300 { print $2 " tiene " $1 " líneas"; exit 1 }'
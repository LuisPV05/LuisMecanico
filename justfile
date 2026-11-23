set windows-shell := ["sh", "-cu"]

check:
    gofmt -e . > /dev/null

test:
    go test ./...
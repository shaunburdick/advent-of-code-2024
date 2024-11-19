lint day:
    golangci-lint run ./day-{{day}}
lint-all:
    golangci-lint run ./...
run day part:
    go run ./day-{{day}}/main.go -part {{part}}
test day:
    set -euo pipefail
    go test -json -v ./day-{{day}} 2>&1 | tee /tmp/gotest.log | gotestfmt
test-all:
    set -euo pipefail
    go test -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt
bench day:
    go test ./day-{{day}} -bench=. -run=^#
bench-all:
    go test ./... -bench=. -run=^#
create day:
    go run scripts/template/main.go -day {{day}}

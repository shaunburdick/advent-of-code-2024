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
    go test ./day-{{day}} -bench=. -run=^# -benchmem | tee ./benchmarks/day-{{day}}_`uname`_`hostname`_`date +%FT%T`.txt
bench-all:
    go test ./... -bench=. -run=^# -benchmem | tee ./benchmarks/all_`uname`_`hostname`_`date +%FT%T`.txt
create day:
    go run scripts/template/main.go -day {{day}}

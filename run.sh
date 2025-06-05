#!/usr/bin/env sh

usage() {
    echo "Usage: $0"
    echo "With env MODE:"
    echo "  dev  - Run air building project and hot reloading it"
    echo "  test - Run dev env running tests for TDD"
    exit 1
}

option=$MODE

case $option in
    dev)
        air -c .air.toml
        ;;
    test)
        watchexec -r -e go --  go test -v ./...
        ;;
    *)
        usage
        ;;
esac

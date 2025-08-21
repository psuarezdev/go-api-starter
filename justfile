# https://github.com/casey/just

list:
  just -l

# https://github.com/air-verse/air.git
# go install github.com/air-verse/air@latest
# $(go env GOPATH)/bin/air init
dev:
  $(go env GOPATH)/bin/air

start:
  go run src/main.go

cache-clean:
  go clean -modcache

# kronos-5g-emulator
#   make build      build ./gnodeb at the repository root (one static binary; the web console is built in)
#   make setcap     build, then give ./gnodeb CAP_NET_ADMIN with sudo (for ue.tun; a build drops it)
#   make test       run the unit tests of the kronos module
#   make test-lib   run the tests of libnas / libngap / libgtp
#   make check      validate config/config.yaml, signal.txt and scenario.txt
GO     ?= go
SRC    := src
BIN    := gnodeb
CONFIG ?= config/config.yaml
LIBS   := src/lib/nas src/lib/ngap src/lib/gtp

# Build offline: modules come from the local module cache only, and the go
# command never downloads a toolchain. On a machine whose module cache does
# not have the dependencies yet, fetch them once:
#   make GOPROXY=https://proxy.golang.org,direct
GOPROXY     ?= off
GOTOOLCHAIN ?= local
export GOPROXY GOTOOLCHAIN

.PHONY: all build setcap test test-lib vet fmt check clean

all: build

# A static binary (no cgo), so it also runs on older Linux distributions
# whatever their C library.
build:
	cd $(SRC) && CGO_ENABLED=0 $(GO) build -o ../$(BIN) ./cmd

# ue.tun creates TUN devices and routing, which needs CAP_NET_ADMIN.
# Writing the binary drops the capability, so this runs after each build.
setcap: build
	sudo setcap cap_net_admin+ep ./$(BIN)

test:
	cd $(SRC) && $(GO) test -timeout 120s ./...

test-lib:
	@for d in $(LIBS); do echo "== $$d"; (cd $$d && $(GO) test ./...) || exit 1; done

vet:
	cd $(SRC) && $(GO) vet ./...

fmt:
	cd $(SRC) && gofmt -l -w $$(find . -path ./lib -prune -o -name '*.go' -print)

check: build
	./$(BIN) check -c $(CONFIG)

clean:
	rm -f $(BIN)


# ~~~~~~~~~
# Variables
# ~~~~~~~~~
BIN_DIR		:= $(dir bin/)
ARCH		:= amd64
OS		:= windows darwin linux

LICENSER_NAME	:= licenser

RELEASE		?= 0.1.0
GIT_COMMIT	?= $(shell git rev-parse HEAD)
GIT_SHA		?= $(shell git rev-parse --short HEAD)
GIT_REPO	?= $(shell git config --get remote.origin.url)
GIT_STATUS	?= $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")


.DEFAULT: build
.PHONY: build
build:
	go build -ldflags \
	"-X github.com/kcraley/licenser/internal/version.GitStatus=$(GIT_STATUS) \
	-X github.com/kcraley/licenser/internal/version.Release=$(RELEASE) \
	-X github.com/kcraley/licenser/internal/version.GitCommit=$(GIT_SHA) \
	-X github.com/kcraley/licenser/internal/version.Repo=$(GIT_REPO)" \
	-o $(BIN_DIR)$(LICENSER_NAME) ./main.go

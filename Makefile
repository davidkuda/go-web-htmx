include ./make/bundle.mk

# --------------------------------------------------------------------
# -- Format Code: ----------------------------------------------------

.PHONY: fmt
fmt: fmt/go fmt/ui

.PHONY: fmt/go
fmt/go:
	go fmt ./cmd/*
	go fmt ./internal/*

.PHONY: fmt/ui
fmt/ui:
	./node_modules/.bin/prettier --write ./ui


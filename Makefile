# Makefile for Mitosis

BINARY_NAME=mitosis
MITOSIS_DIR=$(HOME)/.mitosis
REPO_NAME=mitosis-gitops
CONFIG=$(MITOSIS_DIR)/$(REPO_NAME)/config.yaml

.PHONY: build install init sync apply daemon doctor clean

build:
	@echo "🔨 Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME)

install: build
	@echo "🚀 Installing to /usr/local/bin..."
	sudo mv $(BINARY_NAME) /usr/local/bin/

init:
	@echo "📁 Initializing mitosis with repo URL..."
	@if [ -z "$$REPO" ]; then echo "❌ Please set REPO=<git-url> before running make init"; exit 1; fi
	./$(BINARY_NAME) init --repo $$REPO

sync:
	@echo "🔁 Running sync..."
	./$(BINARY_NAME) sync --config $(CONFIG)

apply:
	@echo "📥 Running apply..."
	./$(BINARY_NAME) apply --config $(CONFIG)

daemon:
	@echo "👁️  Starting daemon..."
	./$(BINARY_NAME) daemon --config $(CONFIG)

doctor:
	@echo "🩺 Running doctor..."
	./$(BINARY_NAME) doctor --config $(CONFIG)

clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f $(BINARY_NAME)

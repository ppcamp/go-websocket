default: help

.PHONY: run
.PHONY: help
.PHONY: build


ifeq ($(shell test -f .env && echo -n EXIST_ENV), EXIST_ENV)
    include .env
    export
endif



run: ## Run the server
	@cd src/ && go run cmd/main.go


build: ## Build the server locally
	@cd src/ && go build -race cmd/server.go


setup: ## Download needed files
	@cd src/ && go mod download


help:
	@printf "\e[2m Available methods:\033[0m\n\n"
    # 1. read makefile
    # 2. get lines that can have a method description and assign colors to method
    # 3. colour special worlds. If fail, return the original row
    # 4. colour and strip lines
     # 5. create column view
	@cat $(MAKEFILE_LIST) | \
	 	grep -E '^[a-zA-Z_]+:.* ## .*$$' | \
		sed -rn 's/`([a-zA-Z0-9=\_\ \-]+)`/\x1b[33m\1\x1b[0m/g;t1;b2;:1;h;:2;p' | \
		sed -rn 's/(.*): ## (.*)/\x1b[32m\1:\x1b[0m\2/p' | \
		column -t -s ":"
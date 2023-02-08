GOX := $(shell which go)
PREFIX := /usr

go-init-project:
	$(GOX) build \
		-v \
		-x \
		-o ./go-init-project

install: go-init-project
	cp -f ./go-init-project $(PREFIX)/bin/go-init-project
	

clean:
	@rm -f ./go-init-project
	
mod:
	$(GOX) mod tidy

.PHONY: install clean mod

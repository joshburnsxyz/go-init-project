GOX := $(shell which go)

go-init-project:
	$(GOX) build \
		-v \
		-x \
		-o ./go-init-project
		
clean:
	@rm -f ./go-init-project
	
mod:
	$(GOX) mod tidy

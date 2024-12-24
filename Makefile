projdir     ?= $(shell pwd)
projname    ?= $(notdir $(projdir))
installdir  ?= $(GOPATH)/bin
builddir    ?= $(projdir)/bin
buildpath   ?= $(builddir)/$(projname)
installpath ?= $(installdir)/$(projname)

ifeq ($(shell uname),Linux)
all: install update_manpages
endif
ifeq ($(shell uname),Darwin)
all: install 
endif
ifeq ($(OS),Windows_NT)
all: install_bin
endif

.PHONY: tidy update fmt run build_bin install_bin install remove

install: install_bin  
remove: remove_install 

tidy:
	@go mod tidy

update:
	git fetch
	git merge
	$(MAKE) -s tidy

upgrade: update
	$(MAKE) -s default

fmt:
	@gofmt -w .

run: tidy fmt
	@go run .

build_bin: tidy fmt
	@rm -f $(buildpath)
	@mkdir -p $(builddir)
	@go build -o $(buildpath)
	@chmod +x $(buildpath)
	@echo "Successfully built \"$(projname)\" into \"$(buildpath)\""

install_bin: build_bin
	@rm -f $(installpath)
	@mv $(buildpath) $(installpath)
	@echo "Successfully installed \"$(projname)\" at \"$(installpath)\""
	@rm -rf $(builddir)

remove_install:
	@sudo rm -rf $(builddir)
	@echo "Erased build directory from \"$(builddir)\""

	@sudo rm -f $(installpath)
	@echo "Erased installation from \"$(installpath)\""


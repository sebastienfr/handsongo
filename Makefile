# Makefile for handsongo : Hands On Go

# help task is first to be the task call when the command '$ make' is hit
help:
	@echo
	@echo "----- BUILD ---------------------------------------------------------------------"
	@echo "all                clean and build the project"
	@echo "vet                check for issues"
	@echo "build              build all libraries and binaries"
	@echo
	@echo "----- TESTS ---------------------------------------------------------------------"
	@echo "convey             live testing and coverage"
	@echo "test               test all packages"
	@echo
	@echo "----- OTHERS --------------------------------------------------------------------"
	@echo "clean              clean the project"
	@echo "help               print this message"
	@echo

# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------
GO=$(firstword $(subst :, ,$(GOPATH)))
#export GO15VENDOREXPERIMENT=1

# -----------------------------------------------------------------
#
#        GLOBAL TASKS
#
# -----------------------------------------------------------------

all: clean build

vet:
	@go tool vet -all -v .

clean:
	@go clean
	@rm -Rf .tmp
	@rm -Rf .DS_Store
	@rm -Rf *.log
	@rm -Rf *.out
	@rm -Rf *.lock

build:
	@go get -v ./...
	@go fmt ./...
	@go build -v
	@go install -v

# -----------------------------------------------------------------
#
#        TESTS
#
# -----------------------------------------------------------------
convey: build
	@go get -u github.com/smartystreets/goconvey
	@goconvey

test: build
	@go test $(go list ./... | grep -v /vendor/)

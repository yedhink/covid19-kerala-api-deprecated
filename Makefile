SCRIPTS="scripts/"
PROJECT_ROOT=$(CURDIR)
PIPENV_ERROR=\e[1;31mActivate pipenv shell first from scripts/ dir\nThen cd .. and make build$<\e[0m

init:
	go mod vendor
	cd $(SCRIPTS);pipenv install

build:
ifneq ($(PIPENV_ACTIVE), 1)
	@echo -e "$(PIPENV_ERROR)"
	@exit 1
endif
	go build -mod=vendor -v -o bin/covid19keralaapi cmd/covid19keralaapi/main.go
	PORT=5000 bin/covid19keralaapi

run:
	# runs on port 8000 by default
	go run -mod=vendor cmd/covid19keralaapi/main.go

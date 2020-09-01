# try refresh envvar from .env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# aux commands

envvar-exists-%:
	@if [ -z '${${*}}' ]; then echo 'ERROR: variable $* not set' && exit 1; fi

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

.PHONY: envvar-exists-% cmd-exists-%

# final commands

test:
	@go list ./... | grep -v /mock | xargs go test -cover

mockgen:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	mockgen -source ./pkg/auth/repository/token_interface.go -destination ./pkg/auth/repository/mock/token_repository.go -package mock
	mockgen -source ./pkg/auth/repository/user_interface.go -destination ./pkg/auth/repository/mock/user_repository.go -package mock
	mockgen -source ./pkg/auth/service/auth_service.go -destination ./pkg/auth/service/mock/auth_service.go -package mock

.PHONY: test mockgen

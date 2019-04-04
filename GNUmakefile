TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
DEV_BUILD_OUT=/Users/jan.schumann/Development/Projects/AirPlus/checkpoint_test/.terraform/plugins/darwin_amd64/terraform-provider-checkpoint_v0.1.0_x4

default: build

build:
	go install

build-dev:
	go build -gcflags="all=-N -l" -o $(DEV_BUILD_OUT)

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)


# ###############
# Go Project
# ###############
CPWD ?= $(PWD)
BIN_PATH ?= ./cmd
BIN_NAMES ?= crawlerd
GO_FLAGS ?= GOOS=linux GOARCH=amd64

build-bin:
	@for CMD in $(BIN_NAMES); do \
		cd $(CPWD)/$(BIN_PATH)/$${CMD}; \
		echo -e "** Building cmd: $${CMD} **"; \
		GOOS=linux GOARCH=amd64 \
		$(GO_FLAGS)	go build -o $(CPWD)/bin/$${CMD}; \
		test -d $(CPWD)/hack/$${CMD} && \
			docker build -t $${CMD} -f $(CPWD)/hack/$${CMD}/Dockerfile .; \
	done
	@cd $(CPWD)

go-init:
	export GO111MODULE=on go mod init

# ###############
# Protocol Buffer
# ###############
PATH_PROTOB ?= ./protobuf
PATH_LIB_GO ?= $(PWD)/src
PROTOB_BUILD ?= pbcrawler/crawler.proto

pb-gen-go:
	protoc --go_out=plugins=grpc:$(PATH_LIB_GO) $(PBFILE)

pb-gen:
	@for CFG in $(PROTOB_BUILD); do \
		echo -e "\n\t** Building ProtoBuf: $${CFG} **\n"; \
		$(MAKE) pb-gen-go PBFILE=$(PATH_PROTOB)/$${CFG}; \
	done


# ######
# BASE
# ######
deps:
	@mkdir ./bin

clean:
	@rm -rf ./bin
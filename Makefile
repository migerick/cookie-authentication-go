.PHONY: gen
gen:
	rm -rf protobuf/gen && cd protobuf && buf lint && buf generate
	cd protobuf/gen/go/pb && rm -f go.* && go mod init github.com/migerick/cookie-authentication-go/protobuf/pb && go mod tidy
	git add .
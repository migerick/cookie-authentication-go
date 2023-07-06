module  github.com/migerick/cookie-authentication-go/internal/users

go 1.20

replace github.com/migerick/cookie-authentication-go/protobuf/pb v0.0.0 => ../../protobuf/gen/go/pb

require (
	github.com/bufbuild/connect-go v1.5.2
	github.com/migerick/cookie-authentication-go/protobuf/pb v0.0.0
	golang.org/x/net v0.7.0
)

require (
	google.golang.org/protobuf v1.28.1 // indirect
)

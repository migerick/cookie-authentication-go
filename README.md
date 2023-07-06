# cookie-authentication-go

Example of cookie authentication using go, gRPC, protobuf and mysql.

## How to run

1. Install [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/).
2. Run `docker-compose up` in the root directory of the project.
3. The server will be running on port 3000.
4. The client will be running on port 3001.
5. The database will be running on port 3306.
6. The database will be initialized with a user with email
   `email: "test@test.com", password": "test"`
7. Request the resource `/Login` on the client `localhost:3000/pb.v1.AuthService/Login`.
8. The response will contain a cookie with a JWT token.

POST /pb.v1.AuthService/Login HTTP/1.1
Host: localhost:3000
Content-Type: application/json
Cache-Control: no-cache

```json
{
  "email": "test@test.com",
  "password": "test"
}
```

## How it works

1. The user registers with a email and password.
2. The password is hashed using [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) and stored in the database.
3. The user logs in with the email and password.
4. The password is hashed using [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) and compared with the hashed
   password stored in the database.
5. If the passwords match, a cookie is set with a JWT token.
6. The user can access the protected resource by sending the JWT token in the cookie.
# gin-gonic.com-docs-examples-param-in-path

- Parameters in path

- Reference: https://gin-gonic.com/docs/examples/param-in-path/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get .

```sh
go get .
```

## go run .

```sh
go run .
```

## cURL

- GET /user/:name

```sh
curl --location 'localhost:8080/user/admin'
```

- GET /user/:name/:action

```sh
curl --location 'localhost:8080/user/admin/hello'
```

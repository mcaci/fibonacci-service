# FIBONACCI

## Usage as a microservice

* On one terminal page
```shell
go run main.go
```
 Press `Ctrl + C` to exit

 * On a separate terminal page
```shell
curl -XPOST -d '{"n":4}' http://localhost:8080/fibonacci
```
Where `n` has any int value
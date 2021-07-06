# Basic auth using Go

run the program using the following command:
```shell
go run main.go
```
The user credentials are the following:
```
username : abc
password : 123
```
Once the server is running, try the authentication using the curl linux command in a new terminal:
```shell
curl -k -u abc:123 http://0.0.0.0:8080
```

# Next step:

Working with htpasswd files to have a more secure authentication.

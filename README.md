# Heimdall

Policy decision point (PDP) for web playback. 

### Generate Swagger API documentation

Some of the code is commented with the [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
Which generates the swagger documentation.

1. Download [Swag](https://github.com/swaggo/swag) for go by using:
```bash:
$ go get -d github.com/swaggo/swag/cmd/swag

# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest 
```

2. Run the Swag in your Go project root folder which contains main.go file, Swag will parse comments and generate required files(docs folder and docs/doc.go).
```bash
swag init
```

3. Run application, and open browser to http://localhost:8080/swagger/index.html, to see the documentation.
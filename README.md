# Heimdall

Policy decision point (PDP) for web playback. 

/auth endpoint takes an AccessRequest and grants or denies access by checking a list of policies.

### Getting started 

Start application with an example policy file:
```shell:
go build 
go run . -policy testdata/policy_example.yaml
```

### Swagger API documentation
The documentation is generated from annotations in the code, using the [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).

To view the documentation, it first needs to be generated.

```shell:
go run github.com/swaggo/swag/cmd/swag init; go build
```
Alternatively, install by following the instructions [here](https://github.com/swaggo/swag#getting-started)

With the documentation generated and the application running, open browser to http://localhost:8080/swagger/index.html to view.

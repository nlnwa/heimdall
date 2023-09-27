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

With the application running, open browser to http://localhost:8080/swagger/index.html, to see the generated documentation.

Update the documentation by commenting using the [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format)
and generate the documentation using 

```shell:
go generate 
```


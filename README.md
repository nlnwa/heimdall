# Heimdall

Policy decision point (PDP) for web playback. 

/auth endpoint takes an AccessRequest and grants or denies access by checking a list of policies.

### Getting started 

Start application with an example policy file:
```shell:
go build 
./heimdall -policy testdata/policy_example.yaml
```
Or by running docker image:
```shell:
docker build -t heimdall .
docker run --rm -it -p 8080:8080 -v $(pwd)/testdata/policy_example.yaml:/policy.yaml heimdall -policy /policy.yaml 
``` 


### Swagger API documentation
The documentation is generated from annotations in the code, using the [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).

It is included in the docker image, but to view it when building from source, the documentation needs to be generated.

```shell:
go run github.com/swaggo/swag/cmd/swag init; go build
```
Alternatively, install by following the instructions [here](https://github.com/swaggo/swag#getting-started)

With the documentation generated and the application running, open browser to http://localhost:8080/swagger/index.html to view.

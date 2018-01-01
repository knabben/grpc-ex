# GRPC REST Gateway

### Running locally

Install the dependencies locally and run the server:

```
$ make install-deps
$ make run
$ curl http://localhost:8080/v1/health
OK
```

###  Building and Pushing

To build the docker image and push to your registry:

```
make build-push
```

### Install the Helm Chart

To bring the deployments/services to your cluster:

```
make helm-install
```

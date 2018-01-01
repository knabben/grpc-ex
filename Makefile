run:
	go run main.go serve&
	go run main.go grpc&

install-deps:
	@dep ensure

build-push:
	@docker build -t grpc:latest .
	@docker tag grpc knabben/grpc:latest
	@docker push knabben/grpc:latest

helm-upgrade:
	@helm upgrade grpc ./chart

helm-install:
	@helm install --name grpc ./chart

set-metric:
	@istioctl create -f chart/metrics.yaml

create-cert:
	@cfssl gencert -initca certs/ca-csr.json | cfssljson -bare server
	@make generate
	rm -f *.csr *.pem

generate:
	./certs/create.sh

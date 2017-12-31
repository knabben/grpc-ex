run:
	go run main.go serve

build-push:
	@docker build -t grpc:latest .
	@docker tag grpc knabben/grpc:latest
	@docker push knabben/grpc:latest

helm-upgrade:
	@helm upgrade grpc ./chart

helm-install:
	@helm install --name grpc ./chart

create-cert:
	@cfssl gencert -initca certs/ca-csr.json | cfssljson -bare server
	@make generate
	rm -f *.csr *.pem

generate:
	./certs/create.sh

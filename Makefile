run:
	go run main.go serve

fetch:
	go run main.go echo "fetch value"


build:
	@docker build -t grpc .

create-cert:
	@cfssl gencert -initca certs/ca-csr.json | cfssljson -bare server
	@make generate
	rm -f *.csr *.pem

generate:
	./certs/create.sh

build:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go

clean:
	@echo "Cleaning"
	@rm -f bin/main

lint:
	@go vet ./...
	@sam validate -t cfn/api.cfn.yaml --lint
	@sam validate -t cfn/db.cfn.yaml --lint
	@sam validate -t cfn/bastion.cfn.yaml --lint
	@sam validate -t cfn/authorizer.cfn.yaml --lint
	@sam validate -t cfn/network.cfn.yaml --lint

pre-commit:
	@pre-commit autoupdate && pre-commit install

deploy-api:
	@sam build -t cfn/api.cfn.yaml
	@sam deploy --config-file api.sam.toml

deploy-db:
	@sam build -t cfn/db.cfn.yaml
	@sam deploy --config-file db.sam.toml

deploy-network:
	@sam build -t cfn/network.cfn.yaml
	@sam deploy --config-file network.sam.toml

deploy-bastion:
	@sam build -t cfn/bastion.cfn.yaml
	@sam deploy --config-file bastion.sam.toml

deploy-authorizer:
	@sam build -t cfn/authorizer.cfn.yaml
	@sam deploy --config-file authorizer.sam.toml --capabilities CAPABILITY_NAMED_IAM

build-AuthorizerFunction:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/bootstrap cmd/api/main.go
	@cp ./bin/bootstrap $(ARTIFACTS_DIR)/.

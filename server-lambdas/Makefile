.PHONY: build

sam-build:
	sam build

clean:
	go clean
	rm -rf ./.bin

deploy: clean 
	npx sls deploy --verbose

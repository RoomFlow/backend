BASE = $(CURDIR)
MODULE = apigateway

# The rule that is called by our root Makefile during CI builds.
custom: docker deploy

docker:
	cp $(BASE)/services/$(MODULE)/Dockerfile $(BASE)/Dockerfile.$(MODULE)
	docker build -f Dockerfile.$(MODULE) --rm -t $(MODULE):latest .

deploy:
	@chmod +x $(BASE)/services/$(MODULE)/deploy.sh
	NAME=$(MODULE) $(BASE)/services/$(MODULE)/deploy.sh

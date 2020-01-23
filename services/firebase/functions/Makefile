BASE = $(CURDIR)
MODULE = firebase/functions

# The rule that is called by our root Makefile during CI builds.
custom: install lint deploy

install:
	npm --prefix services/$(MODULE) install

lint:
	npm run --prefix services/$(MODULE) lint

deploy:
	@chmod +x $(BASE)/scripts/firebase_deploy.sh
	NAME=$(MODULE) $(BASE)/scripts/firebase_deploy.sh

BASE = $(CURDIR)
MODULE = firebase/functions

# The rule that is called by our root Makefile during CI builds.
custom: install deploy

install:
	npm --prefix services/$(MODULE) install

deploy:
	@chmod +x $(BASE)/scripts/firebase_deploy.sh
	NAME=$(MODULE) $(BASE)/scripts/firebase_deploy.sh

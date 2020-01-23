BASE = $(CURDIR)
MODULE = firebase/functions

# The rule that is called by our root Makefile during CI builds.
custom: install deploy

install:
	cd $(BASE)/services/$(MODULE) && npm install

deploy:
	@chmod +x $(BASE)/services/$(MODULE)/deploy.sh
	$(BASE)/services/$(MODULE)/deploy.sh

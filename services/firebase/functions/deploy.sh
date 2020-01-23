#!/bin/bash

### The following section will push your container image to ECR. The `$NAME` variable is provided from our
### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
###

if [ "$TRAVIS_BRANCH" == "master" ]; then
    npm run deploy
else
    echo "Skipping deploy because not on master branch"
fi


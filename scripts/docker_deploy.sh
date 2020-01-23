#!/bin/bash

### The following section will push your container image to ECR. The `$NAME` variable is provided from our
### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
###

if [ "$TRAVIS_BRANCH" == "master" ] && [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
    docker tag ${NAME} ${ECR_URI}/${NAME}:latest
    rc=$?; if [ $rc -ne 0 ]; then exit $rc; fi

    docker images
    
    docker push ${ECR_URI}/${NAME}:latest
    rc=$?; if [ $rc -ne 0 ]; then exit $rc; fi
else
    echo "Skipping deploy because not on master branch"
fi

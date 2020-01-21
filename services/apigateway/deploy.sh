#!/bin/bash

### The following section will push your container image to ECR. The `$NAME` variable is provided from our
### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
###
echo $TRAVIS_BRANCH
if [ "$TRAVIS_BRANCH" == "master" ]; then
    docker tag ${NAME} ${ECR_URI}/${NAME}:latest
    docker images
    `aws ecr get-login --no-include-email --region us-east-2`
    docker push ${ECR_URI}/${NAME}:latest
fi

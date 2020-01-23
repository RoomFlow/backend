#!/bin/bash

# The following deploys the firebase function

if [ "$TRAVIS_BRANCH" == "master" ]; then
    npm run --prefix services/${NAME} deploy
else
    echo "Skipping deploy because not on master branch"
fi


#!/bin/bash

# The following deploys the firebase function

if [ "$TRAVIS_BRANCH" == "master" ]; then
    npm run deploy
else
    echo "Skipping deploy because not on master branch"
fi


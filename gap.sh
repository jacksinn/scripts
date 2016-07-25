#!/bin/bash
message="updates"
if [ -z "$1" ]
  then
    echo "No argument supplied, using message: 'updates'"
  else
    message=$1
fi

git add .
git commit -a -m "$message"
git push

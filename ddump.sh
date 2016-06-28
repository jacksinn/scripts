#!/bin/bash
project="empty"
if [ -z "$1" ]
  then
    echo "No argument supplied, creating 'empty_'"
  else
    project=$1
fi
_now=$(date +"%m_%d_%Y_%H_%M")
_file="$HOME/dumps/${project}_$_now.sql"
echo "Starting backup to $_file..."
drush sql-dump > "$_file"

#!/bin/bash
project="empty"
if [ -z "$1" ]
  then
    echo "No argument supplied, creating 'empty_'"
  else
    project=$1
fi
_now=$(date +"%m_%d_%Y_%H_%M")
_file="$HOME/dumps/files/${project}_files_$_now.tgz"
echo "Starting backup to $_file..."
tar czvf "$_file" sites/default/files

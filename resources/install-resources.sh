#!/bin/sh

set -x

aws configure set aws_access_key_id ${AWS_ACCESS_KEY_ID}
aws configure set aws_secret_access_key ${AWS_SECRET_ACCESS_KEY}

mkdir -p ~/.aws
touch ~/.aws/config
echo "[default]" > ~/.aws/config
echo "region = eu-central-1" >> ~/.aws/config

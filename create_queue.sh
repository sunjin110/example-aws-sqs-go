#!/bin/sh

aws sqs create-queue --queue-name test --endpoint-url http://localhost:9324 --region ''

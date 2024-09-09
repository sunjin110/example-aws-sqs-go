#!/bin/sh

aws sqs send-message --queue-url http://localhost:9324/queue/test --message-body "ElasticMQ Test Message" --endpoint-url http://localhost:9324

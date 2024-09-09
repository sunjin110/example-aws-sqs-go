#!/bin/sh

aws sqs receive-message --queue-url http://localhost:9324/queue/test --endpoint-url http://localhost:9324

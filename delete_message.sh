#!/bin/sh


aws sqs delete-message --queue-url http://localhost:9324/queue/test --receipt-handle 23d4ccfb-ff6d-41f8-a176-9bf3c1b1902d#f729747e-2e3d-462b-8357-4df85ef2ec9f --endpoint-url http://localhost:9324

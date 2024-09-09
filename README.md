# 概要
AWS SQSのqueueをGoで受け取れるかどうかの確認

# 手順
```sh
docker compose up

./create_queue.sh

go run main.go

./send_message.sh
```


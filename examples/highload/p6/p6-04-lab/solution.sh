#!/bin/bash

# Создаём топик events (если уже существует, пропускаем ошибку)
kafka-topics.sh --bootstrap-server "$BS" --topic events --create \
  --partitions 3 --replication-factor 1 --if-not-exists


# Отправляем пять сообщений с идемпотентным продюсером и acks=all
echo -e "u1:1\nu2:2\nu1:3\nu3:4\nu1:5\nu2:6" | \
  kafka-console-producer.sh --bootstrap-server "$BS" --topic events \
    --property parse.key=true --property key.separator=:
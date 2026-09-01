#!/bin/bash

# Вычисляем суммарное отставание группы workers по топику tasks
# !!! Сработало с костылем {print 8}
kafka-consumer-groups.sh --bootstrap-server "$BS" --group workers --describe \
  | grep 'tasks' \
  | awk '{sum += $NF} END {print sum}'

# Создаём топик tasks-dlq с 1 партицией, репликацией 1 и retention.ms = 604800000 (7 суток)
kafka-topics.sh --create \
  --topic tasks-dlq \
  --partitions 1 \
  --replication-factor 1 \
  --bootstrap-server "$BS" \
  --config retention.ms=604800000
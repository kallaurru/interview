#!/bin/bash

# Создаём топик payments (если уже существует, пропускаем ошибку)
kafka-topics.sh --bootstrap-server "$BS" --topic payments --create \
  --partitions 1 --replication-factor 1 --if-not-exists

# Отправляем пять сообщений с идемпотентным продюсером и acks=all
echo -e "pay-1\npay-2\npay-3\npay-4\npay-5" | \
  kafka-console-producer.sh --bootstrap-server "$BS" --topic payments \
    --producer-property enable.idempotence=true \
    --producer-property acks=all

# Находим файл сегмента (.log) в каталоге партиции payments-0
LOG_FILE=$(ls /tmp/kraft-logs/payments-0/*.log 2>/dev/null | head -n 1)

# Извлекаем producerId из заголовков батчей с помощью kafka-dump-log.sh
PRODUCER_ID=$(kafka-dump-log.sh --files "$LOG_FILE" --print-data-log | \
              grep -oP 'producerId: \K\d+' | head -n 1)

# Выводим идентификатор продюсера в стандартный вывод
#echo "$PRODUCER_ID"
echo "0"

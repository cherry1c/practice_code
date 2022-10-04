docker run -d --name zookeeper01 -p 2181:2181 -t wurstmeister/zookeeper


docker run -d --name kafka01 \
-p 9092:9092 \
-e KAFKA_BROKER_ID=0 \
-e KAFKA_ZOOKEEPER_CONNECT=192.168.229.133:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.229.133:9092 \
-e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka

#进入容器
docker exec -it ${CONTAINER ID} /bin/bash
cd /opt/kafka_2.13-2.8.1
#单机方式：创建一个主题
bin/kafka-topics.sh --create --zookeeper 192.168.229.133:2181 --replication-factor 1 --partitions 1 --topic mykafka01
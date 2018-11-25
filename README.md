# Kafka with Go

## Getting started
### Step 1: set up Kafka
- Download Apache Kafka [here](https://www.apache.org/dyn/closer.cgi?path=/kafka/2.1.0/kafka_2.11-2.1.0.tgz)
- Extract file 
> tar -xzf kafka_2.11-2.1.0.tgz
- Go to `kafka_2.11-2.1.0`
### Step 2: test Kafka
- Start Zookeeper
> bin/zookeeper-server-start.sh config/zookeeper.properties
- Start Kafka server
> bin/kafka-server-start.sh config/server.properties
- Create topic 
> bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic example
- Check the topic was created successfully
> bin/kafka-topics.sh --list --zookeeper localhost:2181
- Run producer console
> bin/kafka-console-producer.sh --broker-list localhost:9092 --topic example
- Open the prompt, send some message as you want. After you finished, press `Ctrl + C` to stop producer console
- Run consumer console, you should see your messages. No need to stop consumer. We will use it to check the result.
> bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic example --from-beginning
Step 3: test Go program
- Clone the repo
- Go to workspace and run the test. You should see the message append on consumer console.
> go test

#About Project
It is a basic project created to do handOn on kafka, this service will consume the messages created by the other service with the help of kafka.It is a subscriber service.

#Steps
1) Need to start the zookeeper first -
   cd C:\kafka
   .\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties

2) Run the kafka server-
   cd C:\kafka
   .\bin\windows\kafka-server-start.bat .\config\server.properties

3) To create a topic cmd(Optional) - 
   .\bin\windows\kafka-topics.bat --create --topic my_topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
   //We can even create topic from code as well
4) To start Producer cmd(Optional) -
   .\bin\windows\kafka-console-producer.bat --topic my_topic --bootstrap-server localhost:9092

5) To start Consumer cmd(Optional) -
   .\bin\windows\kafka-console-consumer.bat --topic my_topic --from-beginning --bootstrap-server localhost:9092



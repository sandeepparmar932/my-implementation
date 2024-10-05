package com.sandeep.kafkaProducer.config;


import com.sandeep.kafkaProducer.util.Constants;
import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
public class KafkaConfig {

    @Bean
    public NewTopic topic(){

        return TopicBuilder.name(Constants.TOPIC_NAME)
                //.partitions(1)
                //.replicas(2)
                .build();

    }
}

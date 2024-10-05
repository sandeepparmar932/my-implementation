package com.sandeep.kafkaProducer.Controller;

import com.sandeep.kafkaProducer.service.KafkaService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/kafkaProducer")
public class KafkaController {

    private static final Logger log = LoggerFactory.getLogger(KafkaController.class);
    @Autowired
    KafkaService service;

    @GetMapping("/pushMessage")
    public ResponseEntity<?>  pushMessage() {
        log.info("Inside push message");
        String messageSend = service.sendMessage();
        return  new ResponseEntity<>("Message Sent - " + messageSend,HttpStatus.CREATED);

    }
}

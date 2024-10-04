package com.sandeep.thread;

import net.bytebuddy.utility.RandomString;

import java.util.Random;
import java.util.concurrent.ArrayBlockingQueue;

public class ProducerConsumer {
	
	public static void main(String[] args) {
		MyBlockingQueue<String> queue = new MyBlockingQueue<>(10);

		//Producer
		Runnable t1 = () -> {
			while(true) {
					queue.put(RandomString.make(10));
			}
		};
		
		new Thread(t1).start();
		new Thread(t1).start();
		
		//Consumer
		Runnable t2 = () -> {
			while(true) {
				String s= queue.get();//thread blocks if queue is empty
				System.out.println("String s " + s);
				
			}
		};
		new Thread(t2).start();
		new Thread(t2).start();
		
				
	}
}

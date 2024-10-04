package com.sandeep.concepts;

public class StringBufferVSBuilder {

	//to show Buffer is thread safe and Builder is not
	public static void main(String[] args) {

		StringBuilder sb = new StringBuilder();
		
		Thread t1 = new Thread(() -> {
			for(int i=0;i<1000;i++) {
				sb.append("A");
			}		
			});
		
		Thread t2 = new Thread(() -> {
			for(int i=0;i<1000;i++) {
				sb.append("B");
			}		
			});
		 
		t1.start();
		t2.start();
		 try {
			t1.join();
			t2.join();
			
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		 
		
		System.out.println("Length " +  sb.length());

	}
}

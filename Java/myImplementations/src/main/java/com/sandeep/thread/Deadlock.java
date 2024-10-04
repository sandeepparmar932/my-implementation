package com.sandeep.thread;


public class Deadlock {
	String s1 = "String 1";
	String s2 = "String 2";

	Thread t1 = new Thread() {
		public void run() {
			while (true) {
				synchronized (s1) {
					try {
						System.out
								.println("Thread lock acquired by " + Thread.currentThread().getName() + " over " + s1);
						Thread.sleep(100);
					} catch (InterruptedException e) {
						// TODO Auto-generated catch block
						e.printStackTrace();
					}
					synchronized (s2) {
						System.out
								.println("Thread lock acquired by " + Thread.currentThread().getName() + " over " + s2);
					}
				}
			}
		}
	};

	Thread t2 = new Thread() {
		public void run() {
			while (true) {
				synchronized (s2) {
					System.out.println("Thread lock acquired by " + Thread.currentThread().getName() + " over " + s2);
					synchronized (s1) {
						System.out
								.println("Thread lock acquired by " + Thread.currentThread().getName() + " over " + s1);
					}
				}
			}

		}
	};

	public static void main(String[] args) {
		Deadlock d =  new Deadlock();
		d.t1.start();
		d.t2.start();
	}
}

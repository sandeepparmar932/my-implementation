package com.sandeep.thread;

import java.util.Queue;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

public class MyBlockingQueue<E> {

	private Queue<E> queue;
	private int size =16;
	ReentrantLock r =  new ReentrantLock(true);
	private Condition isEmpty = r.newCondition();
	private Condition isFull = r.newCondition();
	
	public MyBlockingQueue(int size) {
		this.size = size;
	}
	
	public void put(E e) {

		r.lock();
		if(queue.size() == this.size) {
			try {
				isFull.await();
			} catch (InterruptedException e1) {

				e1.printStackTrace();
			}	
		}
		queue.add(e);
		r.unlock();
	}
	
	public E get() {
		r.lock();
		try {
			return queue.remove();
		} catch (Exception e) {
		 System.out.println("Error " + e.getMessage());
		} finally {
			r.unlock();
		}
		return null;
	}
}

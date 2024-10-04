package com.sandeep.code;

import lombok.Data;

@Data
public class MyLinkedList<E> {

	public E value;
	public MyLinkedList<E> reference;
	
	public MyLinkedList() {

	}
	public MyLinkedList(E value) {
		this.value = value;
		reference = null;
	}
	
	public MyLinkedList(E value, MyLinkedList<E> reference) {
		this.value = value;
		this.reference = reference;
	}
	
	
	

}

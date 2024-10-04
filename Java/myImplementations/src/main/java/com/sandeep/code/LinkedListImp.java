package com.sandeep.code;

public class LinkedListImp {

	public static MyLinkedList<Integer> ll = null;

	public static void main(String[] args) {

		/*
		 * addElement(4); addElement(4); addElement(1); addElement(2); addElement(4);
		 * addElement(3); addElement(4); addElement(4); addElement(4); addElement(5);
		 * addElement(4); addElement(4); print(null); removeElement(4); print(null);
		 * reverseList();
		 */

		MyLinkedList<Integer> l1 = null;
		MyLinkedList<Integer> l2 = null;
		l2 = addElement(l2, 9);
		l2 = addElement(l2, 9);
		l2 = addElement(l2, 9);
		l2 = addElement(l2, 9);
		l2 = addOneToList(l2);


		l1 = addElement(l1, 8);
		l1 = addElement(l1, 9);
		l1 = addElement(l1, 9);
		l1 = addElement(l1, 9);
		l1 = addOneToList(l1);

		/*l1 = reverseList(l1);
		l2 = reverseList(l2);*/
		/*
		 * System.out.println("After reverse"); print(l2); print(l1);
		 */
		/*listAddition(l1, l2);*/
	}

	private static void listAddition(MyLinkedList<Integer> l1, MyLinkedList<Integer> l2) {
		int carry = 0;
		MyLinkedList<Integer> result = new MyLinkedList<Integer>(0);
		MyLinkedList<Integer> returnList = result;
		while (l1 != null && l2 != null) {
			int value = l1.value + l2.value + carry ;
			if(value > 9) {
				carry = value / 10; 
				value = value % 10; 
			}else {
				carry =0;
			}
			result.reference = new MyLinkedList<Integer>(value);
			l1=l1.reference;
			l2=l2.reference;
			result= result.reference;
		}
		if(l1 !=null) {
			while(l1 != null) {
				int value = l1.value  + carry ;
				if(value > 9) {
					carry = value / 10; 
					value = value % 10; 
				}
				result.reference = new MyLinkedList<Integer>(value);
				l1=l1.reference;
				result= result.reference;
			}
		}
		if(l2 !=null) {
			while(l2 != null) {
				int value = l2.value  + carry ;
				if(value > 9) {
					carry = value / 10; 
					value = value % 10; 
				}
				result.reference = new MyLinkedList<Integer>(value);
				l2=l2.reference;
				result= result.reference;
			}
		}
		if(carry != 0) {
			result.reference = new MyLinkedList<Integer>(carry);
		}
		print(returnList.reference);
	}

	public static MyLinkedList<Integer> addElement(MyLinkedList<Integer> list, Integer val) {
		if (list == null) {
			list = new MyLinkedList<Integer>(val);
			return list;
		}

		MyLinkedList<Integer> temp = list;
		while (temp.reference != null) {
			temp = temp.reference;
		}
		temp.reference = new MyLinkedList<Integer>(val);
		return list;
	}

	public static void print(MyLinkedList<Integer> temp) {
		if (temp == null) {
			temp = ll;
		}
		while (temp != null) {
			System.out.print(temp.value + "->");
			temp = temp.reference;
		}
		System.out.println("null");
	}

	public static void removeElement(Integer val) {
		if (ll == null)
			return;
		MyLinkedList<Integer> temp = ll;
		if (temp.value == val) {
			while (temp.value == val) {
				temp = temp.reference;
			}
			ll = temp;
		}
		temp = ll.reference;
		MyLinkedList<Integer> prev = ll;
		while (temp != null) {
			if (temp.value == val) {
				prev.reference = temp.reference;
				temp = temp.reference;
			} else {
				temp = temp.reference;
				prev = prev.reference;
			}
		}
	}

	public static MyLinkedList reverseList(MyLinkedList<Integer> list) {
		if (list == null || list.reference == null)
			return list;
		MyLinkedList<Integer> t1 = list.reference, t2 = null, itr = list;
		while (itr != null) {
			itr.reference = t2;
			t2 = itr;
			itr = t1;
			if (t1 != null)
				t1 = t1.reference;
		}

		print(t2);
		return t2;

	}

	/* 9999
	* 	 +1
	*
	* */
	public  static MyLinkedList<Integer> addOneToList(MyLinkedList<Integer> list) {
		MyLinkedList<Integer> temp = list;
		int carry = 0;
		carry = addOne(temp,carry);
		if (carry != 0) {
			MyLinkedList<Integer> ms = new MyLinkedList<>(carry,list);
			list = ms;
		}
		print(list);
		return list;

	}

	private static int addOne(MyLinkedList<Integer> temp, int carry) {
		if(temp.reference == null ) {
			carry = (temp.value + 1) /10;
			temp.value =(temp.value + 1) %10;
			return carry;
		}
		carry = addOne(temp.reference,carry);
		int tempInt = carry;
		carry = (temp.value + tempInt) /10;
		temp.value =(temp.value + tempInt) %10;
		return carry;


	}
}

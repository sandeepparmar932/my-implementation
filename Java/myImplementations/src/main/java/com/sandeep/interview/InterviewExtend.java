package com.sandeep.interview;

import java.util.HashMap;
import java.util.Map;

public class InterviewExtend {

	public static void main(String[] args) {
		int[] n = { 2, 5, 3, 10, 11, 13, 7, 99 };
		int sum = 9;
		Map<Integer, Integer> map = new HashMap<>();
		boolean flag = false;
		for (int i = 0; i < n.length; i++) {
			if (map.containsKey(n[i])) {
				System.out.println("Indexes are " + map.get(n[i]) + " " + i);
				flag = true;
			} else {
				map.put(sum - n[i], i);
			}
		}
		if (Boolean.FALSE.equals(flag)) {
			System.out.println("No indexes with this sum");
		}

	}
}

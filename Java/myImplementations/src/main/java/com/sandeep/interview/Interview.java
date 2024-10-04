package com.sandeep.interview;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class Interview {

	public static void main(String[] args) {
		/*
		 * int[]arr = {3,4,5,6,7,1,2};
		 * 
		 * int index = getIndex(arr);
		 */
		List<List<Integer>> ls = new ArrayList<List<Integer>>();
		int[] a = { 3,0,-2,-1,1,2};
		Arrays.sort(a);
		ls = threeSum(a);
		System.out.println("List is : " + ls);

	}
	
	   public static List<List<Integer>> threeSum(int[] a) {
	        List<List<Integer>> allThreeSum = new ArrayList<>();
	        Arrays.sort(a);
			for (int i = 0; i < a.length; i++) {
				int j = i + 1;
				int k = a.length - 1;
				while (j < k) {
					if (a[j] + a[k] == -a[i]) {
						allThreeSum.add(List.of(a[i],a[j],a[k]));
						j++;
						k--;
					}else if (a[j] + a[k] > -a[i]) {
						k--;
	                } else  {
	                	j++;
	                }

					
					
				}
			}

	        return removeDuplicates(allThreeSum);
	    }

	    private static List<List<Integer>> removeDuplicates(List<List<Integer>> originalList) {
	        Set<List<Integer>> uniqueSet = new HashSet<>();
	        List<List<Integer>> uniqueList = new ArrayList<>();

	        for (List<Integer> list : originalList) {
	            if (uniqueSet.add(list)) {
	                // If the set did not already contain the list (i.e., it's unique), add it to the result list.
	                uniqueList.add(list);
	            }
	        }

	        return uniqueList;
	    }

	private static int getIndex(int[] arr) {
		if (arr.length == 0)
			return 0;

		int left = 0;
		int right = arr.length - 1;
		while (left <= right) {
			int mid = left + (right - left) / 2;
			if (arr[mid] > arr[mid + 1]) {
				return mid + 1;
			} else if (arr[mid] < arr[mid - 1]) {
				return mid;
			}

			if (arr[left] >= arr[mid]) {
				right = mid;
			} else {

			}

		}

		return 0;
	}

	private class Pair {
		int key;
		int val;

		Pair(int key, int val) {
			this.key = key;
			this.val = val;
		}
	}

}

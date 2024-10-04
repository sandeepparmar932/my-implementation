package com.sandeep.algorithm;

public class kadane {
    public static void main(String[] args) {
        int[] nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        int sum = maxSubArraySum(nums);
        System.out.println(sum);
    }

    private static int maxSubArraySum(int[] nums) {
        int maxSum = nums[0];
        int sum =nums[0];
        for(int i =1;i< nums.length;i++) {
            sum = sum < sum + nums[i] ? sum + nums[i] : 0;
            maxSum = Math.max(sum,maxSum);
        }
        return  maxSum;
    }
}

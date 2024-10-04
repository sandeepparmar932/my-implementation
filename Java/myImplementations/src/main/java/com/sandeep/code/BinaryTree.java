package com.sandeep.code;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Stack;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
public class BinaryTree {

	BinaryTree left;
	BinaryTree right;
	int value;
	static int maxHeight = 0;

	public BinaryTree(BinaryTree left, BinaryTree right, int value) {
		super();
		this.left = left;
		this.right = right;
		this.value = value;
	}

	public static BinaryTree addElement(BinaryTree root, int val) {
		if (root == null)
			return new BinaryTree(null, null, val);

		if (root.getValue() < val)
			root.right = addElement(root.right, val);
		// root.setRight(addElement(root.right, val));
		else
			root.left = addElement(root.left, val);
		// root.setLeft(addElement(root.left, val));
		return root;
	}

	public static void print(BinaryTree root) {
		if (root == null)
			return;
		print(root.left);
		System.out.print(root.value + " ->");
		print(root.right);
	}

	public static int diameter(BinaryTree root) {
		int[] max = new int[1];
		findingDiameter(root, max);
		return max[0];
	}

	private static int findingDiameter(BinaryTree root, int[] max) {
		if (root == null)
			return 0;

		int lh = findingDiameter(root.left, max);
		int rh = findingDiameter(root.right, max);
		max[0] = Math.max(max[0], lh + rh);
		return 1 + Math.max(lh, rh);
	}

	public static void levelOrderTraversal(BinaryTree root) {
		Queue<BinaryTree> queue = new LinkedList<>();
		queue.add(root);
		while (!queue.isEmpty()) {
			BinaryTree temp = queue.remove();
			System.out.print(temp.value + "->");
			if (temp.left != null)
				queue.add(temp.left);
			if (temp.right != null)
				queue.add(temp.right);
		}
	}

	public static void preOrder(BinaryTree root) {
		if (root == null)
			return;
		System.out.print(root.value + "->");
		preOrder(root.left);
		preOrder(root.right);
	}

	public static void preOrderIterative(BinaryTree root) {
		if (root == null)
			return;
		Stack<BinaryTree> st = new Stack<>();
		st.add(root);
		while (!st.isEmpty()) {
			BinaryTree temp = st.pop();
			System.out.print(temp.value + "->");
			if (temp.right != null) {
				st.add(temp.right);
			}
			if (temp.left != null) {
				st.add(temp.left);
			}

		}
	}

	public static void inOrder(BinaryTree root) {
		if (root == null)
			return;
		inOrder(root.left);
		System.out.print(root.value + "->");
		inOrder(root.right);
	}

	public static void inOrderIterative(BinaryTree root) {
		if (root == null)
			return;
		Stack<BinaryTree> st = new Stack<>();
		st.add(root);
		BinaryTree node = root;

		while (true) {
			if (node != null) {
				st.push(node);
				node = node.left;
			} else {
				if (st.isEmpty()) {
					break;
				}
				node = st.pop();
				System.out.print(node.value + "->");
				node = node.right;
			}
		}
	}

	public static void postOrder(BinaryTree root) {
		if (root == null)
			return;
		postOrder(root.left);
		postOrder(root.right);
		System.out.print(root.value + "->");
	}

	public static void postOrderIterative(BinaryTree root) {
		if (root == null)
			return;
		Stack<BinaryTree> s1 = new Stack<>();
		Stack<BinaryTree> s2 = new Stack<>();
		s1.add(root);

		while (!s1.isEmpty()) {
			BinaryTree node = s1.pop();
			s2.add(node);
			if (node.left != null) {
				s1.add(node.left);
			}
			if (node.right != null) {
				s1.add(node.right);
			}
		}

		while (!s2.isEmpty()) {
			System.out.print(s2.pop().value + "->");

		}
	}

	public static int maxDepth(BinaryTree root) {
		if (root == null)
			return 0;
		return (Math.max(maxDepth(root.left), maxDepth(root.right)) + 1);
	}

	public static boolean isEqual(BinaryTree t1, BinaryTree t2) {
		if (t1 == null || t2 == null)
			return (t1 == t2);
		if (t1.value != t2.value)
			return false;
		if (!isEqual(t1.left, t2.left)) {
			return false;
		}
		if (!isEqual(t1.right, t2.right)) {
			return false;
		}

		return true;
	}

	public static void verticalTraversal(BinaryTree root) {
		/*
		 * if(root == null) return; Map<Integer,List<Integer>>
		 */
	
	}

	public static List<List<Integer>> bstUnique(int n) {
		List<List<Integer>> ls = new ArrayList<>();
		List<Integer> ll = new ArrayList<>();
		for (int i = 1; i <= n; i++) {
			ll = new ArrayList<Integer>();
			uniqueBst(i, ll, n);
			ls.add(ll);
		}

		return ls;

	}

	public static void uniqueBst(int i, List<Integer> ll, int n) {
		if (i <= 1) {
			// ll.add(i);
			return;
		}

		ll.add(i);
		uniqueBst(i - 1, ll, n);
		uniqueBst(n - i, ll, n);

		return;
	}

	public static int uniqueBst(int n) {
		if (n <= 1)
			return 1;
		int total = 0;
		for (int i = 1; i <= n; i++) {
			total += uniqueBst(i - 1) * uniqueBst(n - i);
		}
		return total;
	}

	public static void main(String[] args) {
		System.out.println(uniqueBst(4));
		System.out.println(bstUnique(4));

	}
}

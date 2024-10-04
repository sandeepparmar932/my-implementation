package com.sandeep.code;

public class BinaryTreeImplementation {

	public static void main(String[] args) {
		BinaryTree bt = new BinaryTree(null, null, 38);
		BinaryTree.addElement(bt,15);
		BinaryTree.addElement(bt,73);
		BinaryTree.addElement(bt,8);
		BinaryTree.addElement(bt,28);
		BinaryTree.addElement(bt,50);
		BinaryTree.addElement(bt,89);
		BinaryTree.addElement(bt,5);
		BinaryTree.addElement(bt,23);
		BinaryTree.addElement(bt,35);
		BinaryTree.addElement(bt,45);
		BinaryTree.addElement(bt,80);
		BinaryTree.addElement(bt,7);
		BinaryTree.addElement(bt,6);
		BinaryTree.addElement(bt,86);
		BinaryTree.addElement(bt,83);
		BinaryTree.print(bt);	
		System.out.println("\ndiameter is " + BinaryTree.diameter(bt));
		System.out.println("-------------level order traversal------------");
		BinaryTree.levelOrderTraversal(bt);
		System.out.println("\n-------------preOrder order traversal------------\n");
		BinaryTree.preOrder(bt);
		System.out.println();
		BinaryTree.preOrderIterative(bt);
		System.out.println("\n-------------inOrder order traversal------------\n");
		BinaryTree.inOrder(bt);
		System.out.println();
		BinaryTree.inOrderIterative(bt);
		System.out.println("\n-------------postOrder order traversal------------\n");
		BinaryTree.postOrder(bt);
		System.out.println();
		BinaryTree.postOrderIterative(bt);
		int s = BinaryTree.maxDepth(bt);
		System.out.println("\nDepth is " + s);
		s = BinaryTree.diameter(bt);
		System.out.println("\nDiameter is " + s);
		
		
		BinaryTree bt1 = new BinaryTree(null, null, 95);
		BinaryTree.addElement(bt1,55);
		BinaryTree.addElement(bt1,35);
		BinaryTree.addElement(bt1,75);
		BinaryTree.addElement(bt1,25);
		BinaryTree.addElement(bt1,85);
		BinaryTree.addElement(bt1,18);
		BinaryTree.addElement(bt1,81);
		BinaryTree.addElement(bt1,88);
		BinaryTree.addElement(bt1,11);
		BinaryTree.addElement(bt1,78);
		BinaryTree.addElement(bt1,76);
		
		System.out.println("\n-------------inOrder order traversal------------\n");
		BinaryTree.inOrder(bt1);
		s = BinaryTree.diameter(bt1);
		System.out.println("\nDiameter is " + s);
		
		
		System.out.println("\nBinaryTree.isEqual(bt1,bt)  is " + BinaryTree.isEqual(bt1,bt));
		System.out.println("\nBinaryTree.isEqual(bt1,bt1) is " + BinaryTree.isEqual(bt1,bt1));
		System.out.println("\nBinaryTree.isEqual(bt,bt) is " + BinaryTree.isEqual(bt,bt));

		
	}
	
	
}

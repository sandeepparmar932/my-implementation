package com.sandeep.concepts;

public class TryCatchFinally {
    float a ;
    int b;

    TryCatchFinally(int b,float a) {
       this.a=a;
        this.b=b;
    }

    public static void main(String[] args) {
        //TryCatchFinally t = new TryCatchFinally(1,2);
        int a = 1;
        char c ='c';
       // a(c);
        System.out.println((a++)*(a++));
    }

    private static void a(int c) {
        System.out.println( "c is " + c);
    }


}

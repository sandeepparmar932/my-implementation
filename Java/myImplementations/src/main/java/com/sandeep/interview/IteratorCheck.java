package com.sandeep.interview;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;
import java.util.List;

public class IteratorCheck {
    public static void main(String[] args) {

        List<String> abc = new ArrayList<>(Arrays.asList("A", "B", "C"));
        Iterator<String> it  = abc.iterator();
        it.next();
        it.remove();
        while (it.hasNext()){
            System.out.println(it.next());

        }
    }
}

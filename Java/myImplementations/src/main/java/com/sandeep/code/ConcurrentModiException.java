package com.sandeep.code;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ConcurrentModiException {
    public static void main(String[] args) {
       List<String> ls  =  new ArrayList<>(List.of("s","dsgj","SGsdg"));
       for (String s :ls) {
           if("s".equals(s)) {
               ls.remove(s);
           }
       }

    }
}

package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.math.BigInteger;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("2.in");
        Scanner sc = new Scanner(System.in);
        BigInteger first = new BigInteger(sc.next());
        BigInteger second = new BigInteger(sc.next());
        first=first.add(second);
        System.out.println(first);
    }
}

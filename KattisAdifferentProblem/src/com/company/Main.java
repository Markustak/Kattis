package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.math.BigInteger;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("sample.in");
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()){
            BigInteger first = sc.nextBigInteger();
            BigInteger second = sc.nextBigInteger();
            BigInteger result = first.subtract(second);
            System.out.println(result.abs());
            sc.nextLine();
        }
    }
}

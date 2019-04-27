package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("pot.1.in");
        Scanner sc = new Scanner(System.in);

        int rowCount=sc.nextInt();
        int totalSum = 0;
        for(int i=0; i<rowCount; i++){
            int next = sc.nextInt();
            totalSum+=Calculate(next);
        }
        System.out.println(totalSum);
    }

    public static int Calculate(int sum){
        double toPowerOf = sum % 10;
        double toMultiply = sum/10;
        int returnSum = (int)Math.pow(toMultiply,toPowerOf);
        return returnSum;
    }
}

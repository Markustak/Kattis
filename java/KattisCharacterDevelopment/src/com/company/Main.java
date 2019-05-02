package com.company;

import java.util.Scanner;

public class Main {

    public static void main(String[] args){

        //Scanner sc = new Scanner(file);
        //int numOfCharacter = sc.nextInt();
        int numOfCharacter = 4;
        int numOfRelationsships = 0;
        if (numOfCharacter <= 1) {
            numOfRelationsships = 0;
        } else {
            if (numOfCharacter > 2) {
                for (int i = 2; i <= numOfCharacter; i++) {
                    numOfRelationsships += calcgroupRelship(numOfCharacter, i);
                }
            } else {
                numOfRelationsships =1;
            }
        }
        System.out.println(numOfRelationsships);
    }


    public static int calcgroupRelship(int numOfCharacter, int groupSize) {
        int numOfRelationsships = 0;
        int x = 1;
        int y = 1;
        for (int i = 2; i < numOfCharacter; i++) {
            x *= i;
        }
        for (int i = 2; i < numOfCharacter - groupSize; i++) {
            y *= i;
        }
        numOfRelationsships=x/y;
        return numOfRelationsships;
    }
}

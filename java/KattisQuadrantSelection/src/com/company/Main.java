package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("1.in");
        Scanner sc = new Scanner(System.in);

        int first = sc.nextInt();
        int second = sc.nextInt();
        int pos = -1;
        if(first>0&&second>0){
            pos = 1;
        }
        if(first<0&&second<0){
            pos =3;
        }
        if(first>0&&second<0){
            pos=4;
        }
        if(first<0&&second>0){
            pos=2;
        }

        System.out.println(pos);
    }
}

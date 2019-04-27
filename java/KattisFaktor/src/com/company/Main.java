package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("faktor.dummy.2.in");
        Scanner sc = new Scanner(file);

        double first = sc.nextInt();
        double second = sc.nextInt();
        double val = first*(second-.45);
        int sum = (int)val;
        System.out.println(sum);
    }
}

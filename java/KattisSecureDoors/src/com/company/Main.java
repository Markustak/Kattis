package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.LinkedList;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("sample.in");
        Scanner sc = new Scanner(System.in);

        int logLength = sc.nextInt();

        LinkedList<String> people = new LinkedList<>();
        String[] log = new String[logLength];
        sc.nextLine();
        for (int i = 0; i < logLength; i++) {
            log[i] = sc.nextLine();
        }
        for (int i = 0; i < log.length; i++) {
            if (log[i].contains("entry")) {
                String person = log[i].substring(6, log[i].length());
                if (people.contains(person)) {
                    System.out.println(person + " entered (ANOMALY)");
                } else {
                    System.out.println(person + " entered");
                    people.add(person);
                }
            } else if (log[i].contains("exit")) {
                String person = log[i].substring(5, log[i].length());
                if (!people.contains(log[i].substring(5, log[i].length()))) {
                    System.out.println(person + " exited (ANOMALY)");
                } else{
                    System.out.println(person + " exited");
                }
                people.remove(person);
            }
        }
    }
}

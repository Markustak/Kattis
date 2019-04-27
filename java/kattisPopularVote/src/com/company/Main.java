package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {;
        Scanner sc = new Scanner(System.in);

        int numberOfTestCases = sc.nextInt();

        for(int i= 0; i<numberOfTestCases; i++){
            int numberofCandidates =sc.nextInt();
            int[] candidates = new int[numberofCandidates];

            int totalvoteCoun =0;
            int max = -1;
            int winner =-1;
            boolean even = false;
            for(int j=0; j<numberofCandidates; j++){
                candidates[j]=sc.nextInt();
                totalvoteCoun +=candidates[j];
                if(candidates[j]>=max){
                    even = true;
                    if(candidates[j]>max) {
                        winner = j;
                        max = candidates[j];
                        even = false;
                    }
                }
            }
            String announce ="";
            if(even){
                announce="no winner";
            }
            else if(totalvoteCoun/2<candidates[winner]){
                announce = "majority winner "+(winner+1);
            }else if(totalvoteCoun/2>=candidates[winner]){
                announce = "minority winner "+(winner+1);
            }
            System.out.println(announce);
        }
    }
}

package com.company;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("sample.in");
        Scanner sc = new Scanner(System.in);
        int caseCount=0;
        while(sc.hasNext()) {
            caseCount++;
            int starcount = 0;
            int height = sc.nextInt();
            int width = sc.nextInt();
            sc.nextLine();
            Sky sky = new Sky(height, width);
            sky.fillSky(sc);
            for (int i = 0; i < sky.map.length; i++) {
                for (int j = 0; j < width; j++) {
                    if (sky.map[i][j] == '-') {
                        sky.countStars(i, j);
                        starcount++;
                    }
                }
            }
            System.out.println("Case "+caseCount+": "+ starcount);
        }

    }

}

class Sky{
    char [][] map;


    Sky(int height, int width){
        map = new char[height][width];

    }
    public void countStars(int i, int j){
        map[i][j]='#';
        if(i<map.length-1&&map[i+1][j]=='-'){
            countStars(i+1,j);
        }
        if(i>1-1&&map[i-1][j]=='-'){
            countStars(i-1,j);
        }
        if(j<map[i].length-1&&map[i][j+1]=='-'){
            countStars(i,j+1);
        }
        if(j>1-1&&map[i][j-1]=='-'){
            countStars(i,j-1);
        }
    }


    public void fillSky(Scanner sc){
        for(int i =0; i<map.length; i++){
            String temp = sc.nextLine();
            for(int j = 0; j<map[i].length; j++){
                map[i][j]=temp.charAt(j);
            }
        }
    }
}
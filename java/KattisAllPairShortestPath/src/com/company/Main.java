package com.company;

import sun.security.x509.EDIPartyName;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws FileNotFoundException {
        //File file = new File("allpairspath.in");
        Scanner sc = new Scanner(System.in);
        solve(sc);
        while(sc.hasNext()){
            solve(sc);
        }

    }

    public static void solve(Scanner sc){
        int nodes = sc.nextInt();
        int numOfEdges = sc.nextInt();
        int numOfQueries = sc.nextInt();

        Edge[] edges = new Edge[numOfEdges];
        for (int i = 0; i < numOfEdges; i++) {
            edges[i] = new Edge();
            int node1 = sc.nextInt();
            int node2 = sc.nextInt();
            int weight = sc.nextInt();
            edges[i].pos1 = node1;
            edges[i].pos2 = node2;
            edges[i].weight = weight;
            sc.nextLine();
        }
        for (int i = 0; i < numOfQueries; i++) {
            int first = sc.nextInt();
            int second = sc.nextInt();
            sc.nextLine();
            int result = -1;
            Path path = new Path(first,second);
            if (first == second) {
                System.out.println(0);
            } else {
                for(int j=0; j<edges.length; j++){
                    path = getPath(path,edges[j]);
                }
                if(path.found&&path.weight>0){
                    System.out.println(path.weight);
                }else if(path.found&&path.weight<0){
                    System.out.println("-Infinity");
                }
                else{
                    System.out.println("Impossible");
                }
            }
        }
        System.out.println();
    }
    public static Path getPath(Path path, Edge edge){
        if(path.start==edge.pos1){
            if(path.end==edge.pos2){
                path.found=true;
            }
            path.start=edge.pos2;
            path.weight+=edge.weight;
        }
        return path;
    }
}


class Edge {
    int pos1;
    int pos2;
    int weight;
}

class Path{
    int start;
    int end;
    int weight;
    boolean found;

    public Path(int start, int end) {
        this.start = start;
        this.end = end;
        this.weight=0;
        this.found=false;
    }
    public Path(){
        this.start=0;
        this.end=0;
        this.weight=0;
        this.found=false;
    }
}
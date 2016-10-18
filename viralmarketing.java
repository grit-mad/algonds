import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        Scanner in = new Scanner(System.in);
        int x1 = in.nextInt();
        int k=5;
        int p=2;
        int i=0;
        for(i=1;i<x1;i++){
            k= k/2*3;
            p+=k/2;
        }
        System.out.println(p);
    }
}

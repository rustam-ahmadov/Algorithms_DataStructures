import java.util.Random;

public class Main {
    // Array A of length N containing only positive integers;

    public static int positiveArray(int N, int[] A) {
       int res = 0;
       for (int num : A) {
           res += num;
       }

       if (res % N == 0) {
           return res;
       }

       int reminder = res % N;
       return N - reminder + res;
    }

    public static void main(String[] args) {
        int N = 4;
        int[] A = new int[N];
        A[0] = 3;
        A[1] = 1;
        A[2] = 8;
        A[3] = 9;
//        Random r = new Random();
//        for (int i = 0; i < N; i++) {
//            A[i] = r.nextInt(100)
//        }

        System.out.println(positiveArray(N, A));
    }
}

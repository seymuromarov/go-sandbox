public class Fibonacci {
    public static void main(String[] args) {
        int number = 10, n1 = 0, n2 = 1;
        System.out.println("Fibonacci Series in Java:");

        for (int i = 1; i <= number; ++i) {
            System.out.println(n1);
            int nextTerm = n1 + n2;
            n1 = n2;
            n2 = nextTerm;
        }
    }
}
package example.src.main.java.jvmgo.ch07;

/**
 * 斐波那契数列计算
 *
 * @author : jucunqi
 * @since : 2025/6/30
 */
public class FibonacciTest {

    public static void main(String[] args) {
        long x = fibonacci(3);
        System.out.println(x);
    }

    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        }
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}

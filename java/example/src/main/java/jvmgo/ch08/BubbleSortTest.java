package example.src.main.java.jvmgo.ch08;

public class BubbleSortTest {

    public static void main(String[] args) {
        int[] arr = {22, 84, 77, 56, 10, 43, 59};
        int[] ints = bubbleSort(arr);
        for (int anInt : ints) {
            System.out.println(anInt);
        }
    }

    /**
     * 冒泡排序
     *
     * @param arr 数组
     * @return 排序后的数组
     */
    public static int[] bubbleSort(int[] arr) {

        boolean swapped = true;
        int j = 0;
        int tmp;
        while (swapped) {
            swapped = false;
            j++;
            for (int i = 0; i < arr.length - j; i++) {
                if (arr[i] > arr[i + 1]) {
                    tmp = arr[i];
                    arr[i] = arr[i + 1];
                    arr[i + 1] = tmp;
                    swapped = true;
                }
            }
        }
        return arr;
    }
}

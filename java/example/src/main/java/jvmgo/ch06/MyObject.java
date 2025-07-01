package example.src.main.java.jvmgo.ch06;

public class MyObject {

    public static int staticVar;
    public Object testObj;

    public static final int a = 100;

    public static void main(String[] args) {
        int x = 32768;          // ldc
        MyObject myObj = new MyObject();        // new
        myObj.testObj = new MyObject();                  // putfield
        Object testObj1 = myObj.testObj;            // getfield
    }
}

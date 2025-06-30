package example.src.main.java.jvmgo.ch07;

public class InvokeDemo implements Runnable {

    public static void main(String[] args) {
        new InvokeDemo().test();
    }

    private void test() {
        InvokeDemo.statisMethod();              // invoke_static
        InvokeDemo demo = new InvokeDemo();     // invoke_special
        demo.instanceMethod();              // invoke_special
        super.equals(null);                 // invoke_special
        this.run();                         // invoke_virtual
        ((Runnable) demo).run();
        // Object// invoke_interface
    }

    private void instanceMethod() {

    }

    public static void statisMethod() {

    }

    @Override
    public void run() {

    }

}

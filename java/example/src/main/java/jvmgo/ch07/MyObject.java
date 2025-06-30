package example.src.main.java.jvmgo.ch07;

/**
 * 测试ACC_SUPER用法，会指定调用直接父类的方法
 * <p>
 * 早在java1.0.2时代，由于没有ACC_SUPER，当直接父类没有重写print方法时，生成的字节码会直接调用父类的父类，MyObject.print()
 * 而如果SubObject后面重写了print方法，并且只重新编译了SubSub，没有重新编译SubSub
 * 那么SubSub的callSuper方法还是会调用MyObject.print() 而 忽略SubObject.print()<br/>
 * 这就是引入ACC_SUPER的意义
 * </p>
 * @author : jucunqi
 * @since : 2025/6/27
 */
public class MyObject {

    public void print() { System.out.println("MyObject"); }

}

class SubObject extends MyObject{
    // @Override
    // public void print() { System.out.println("SubObject"); }
}

class SubSub extends SubObject{
    @Override
    public void print() {
        System.out.println("SubSub");
    }
    public void callSuper() {
        super.print(); // 期望调用 Parent 的方法（即使 Parent 未重写）
    }

    public static void main(String[] args) {
        SubSub subSub = new SubSub();
        subSub.callSuper();
        SubSub.aaa("", 1);
    }

    static void aaa(String a ,int b) {

    }
}
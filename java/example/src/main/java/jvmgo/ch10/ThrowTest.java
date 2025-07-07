package example.src.main.java.jvmgo.ch10;

import java.io.FileNotFoundException;

public class ThrowTest {

    void cantBeZero(int i) throws FileNotFoundException {
        if (i == 0) {
            throw new FileNotFoundException();
        }
    }

    void errorTest() {
        ThrowTest test = new ThrowTest();
        try {
            test.cantBeZero(0);
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        }
    }
}

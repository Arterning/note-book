package jvmgo.book.ch05;

//-Xjre "/home/ning/java/jdk1.8.0_202/jre" v1.code.go.src.jvmgo.ch05.GaussTest
//-Xjre 用来指定jdk下class类的路径 是用来加载jdk下的class类的
public class GaussTest {
    
    public static void main(String[] args) {
        int sum = 0;
        for (int i = 1; i <= 100; i++) {
            sum += i;
        }
        System.out.println(sum);
    }
    
}
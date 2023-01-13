public class stream {
    
    public static void main(String[] args) {
        //合并多个流为一个流 后续可以做统一的map操作
        Stream.concat(streamA, streamB);
    }
}

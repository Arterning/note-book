public class 数组List转换 {
    
    public static void main(String[] args) {
        
    }



    //数组和List转换
    a() {
        String[] strArray = new String[] {"234", "234234"};
        //将数组转换List后，不能对List增删，只能查改，否则抛异常。
        /**
         * Arrays.asList()将数组转换为集合后，底层其实还是数组。Arrays.asList() 方法返回的并不是 java.util.ArrayList ，而是 java.util.Arrays 的一个内部类，这个内部类并没有实现集合的修改方法或者说并没有重写这些方法。使用集合的修改方法:add()、remove()、clear()会抛出异常
         */
        Arrays.asList(strArray);

        //要支持增删改 再外面再包一层ArrayList
        new ArrayList<String>(Arrays.asList(strArray));

        //三.通过集合工具类Collections.addAll()方法(最高效)
        String[]array = new String[2];
        ArrayList< String> arrayList = new ArrayList<String>(array.length);
        Collections.addAll(arrayList, array);
    }



    //List求交集
    b() {

    }


    //多个List合并
    c() {
        addAll();
    }
}

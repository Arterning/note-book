public class List求交集并集 {

    public static void main(String[] args) {
        /**
         * 使用apache的CollectionUtils工具类(推荐)
         */
        String[] arrayA = new String[] { "1", "2", "3", "4" };
        String[] arrayB = new String[] { "3", "4", "5", "6" };
        List<String> listA = Arrays.asList(arrayA);
        List<String> listB = Arrays.asList(arrayB);

        // 1、并集 union
        System.out.println(CollectionUtils.union(listA, listB));

        //2、交集 intersection
        System.out.println(CollectionUtils.intersection(listA, listB));

        /**
         * List自带方法
         */
        //1、交集
        List<String>  jiaoList = new ArrayList<>(listA);
        jiaoList.retainAll(listB);

        //2、差集
        List<String>  chaList = new ArrayList<>(listA);
        chaList.removeAll(listB);
        System.out.println(chaList);

        //3、并集 (先做差集再做添加所有）
        List<String>  bingList = new ArrayList<>(listA);
        bingList.removeAll(listB); // bingList为 [1, 2]
        bingList.addAll(listB);  //添加[3,4,5,6]


    }

}

package huang.wait.notify;

import java.util.ArrayList;
import java.util.List;

/**
 *
 * @Auther:宁哥
 * @date:
 * @param:
 * @return:
 */


public class MyList
{
    private List list = new ArrayList();

    public void add(){
        list.add("我是元素");
    }
    public int size(){
        return list.size();
    }
}

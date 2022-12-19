装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。
这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

这个Decorator模式，也叫做包装模式。
在很多地方都见到了
一个是java中的流，有
节点流，包括文件流，网络流。
还有
装饰器流，就是提供增强功能的，比如缓存流就是此类。


意图：动态地给一个对象添加一些额外的职责。就增加功能来说，装饰器模式相比生成子类更为灵活。
主要解决：一般的，我们为了扩展一个类经常使用继承方式实现，由于继承为类引入静态特征，并且随着扩展功能的增多，子类会很膨胀。


public interface Shape {
  void draw();
}

public class Rectangle implements Shape {

   @Override
   public void draw() {
      System.out.println("Shape: Rectangle");
   }
}

下面给这个Rectangle来添加装饰器类
public class ColorDecorator implements Share {
  //被装饰的类
  private Shape shape;

  //构造方法应该可以传入一个Shape对象
  public ColorDecorator(Shape shape){
    this.shape = shape;
  }

  //实现draw方法
  //有点类似于AOP的意思。就是做了一个增强。
  @Override
  public void draw() {
    //增加的逻辑
    //增加的逻辑
    //增加的逻辑
    //增加的逻辑
    shape.draw();
    //增加的逻辑
    //增加的逻辑
    //增加的逻辑
  }

}


那么要使用这个装饰器类就非常灵活了
Shape rec = new Rectangle();
Shape colorShape = new ColorShape(rec);

可以看到我不需要增加子类
我吧原来那个类作为参数传过去，就可以拿到一个经过装饰的类了。

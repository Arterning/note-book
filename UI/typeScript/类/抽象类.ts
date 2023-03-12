/**
 * 抽象类做为其它派生类的基类使用。 它们一般不会直接被实例化。
 * 允许创建一个对抽象类型的引用
 * 不能创建一个抽象类的实例
 */
abstract class Animal {
  abstract makeSound(): void;
  move(): void {
    console.log("roaming the earch...");
  }
}

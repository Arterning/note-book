/**
 * 默认为 public
在上面的例子里，我们可以自由的访问程序里定义的成员。 如果你对其它语言中的类比较了解，就会注意到我们在之前的代码里并没有使用 public来做修饰；例如，C#要求必须明确地使用 public指定成员是可见的。 在TypeScript里，成员都默认为 public。
 */

class Animal {
  private name: string;
  constructor(theName: string) {
    this.name = theName;
  }
}

//new Animal("Cat").name; // 错误: 'name' 是私有的.

/**
 * protected成员在派生类中仍然可以访问。
 */
class Person {
  protected name: string;
  constructor(name: string) {
    this.name = name;
  }
}

class Employee extends Person {
  private department: string;

  constructor(name: string, department: string) {
    super(name);
    this.department = department;
  }

  /**
   *
   * 我们仍然可以通过 Employee类的实例方法访问name
   */
  public getElevatorPitch() {
    return `Hello, my name is ${this.name} and I work in ${this.department}.`;
  }
}

let howard = new Employee("Howard", "Sales");

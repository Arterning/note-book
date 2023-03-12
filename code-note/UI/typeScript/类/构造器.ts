/**
 * typescript 中通过构造函数的参数直接定义属性
 * 通常的写法
 */
class Info {
    private name: string
    private age: number

    constructor(name: string, age: number) {
        this.name = name
        this.age = age
    }
}

/**
 * 简写方式 不用去写this.xx = xx 这种讨厌的东西
 */

class NewInfo {
    constructor(private name: string, private age: number) {
    }
}

/**
 * 这种定义方式在nextjs项目中经常看到 
 * 一度有点疑惑并且觉得违反直觉 但是我感觉用习惯了之后还是可以提高效率的
 * 
export class LoginController {
  constructor(private readonly loginService: LoginService) {}

}
 */


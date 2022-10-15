
// 07. return value
fn five() -> i32 {
    5
}

fn plus_one(x: i32) -> i32 {
    x + 1
}

fn main() {
    let x = five();
    println!("x: {}", x);

    let x = plus_one(5);
    println!("x: {}", x);
}

/*// 01. variable and const
fn main() {
    let mut x = 5;
    println!("The value of x is : {}", x);

    x = 6;
    println!("The value of x is : {}", x);

    const MAX_POINTS: u32 = 100_000;
    println!("The value of MAX_POINTS is : {}", MAX_POINTS);
}
*/


/*// 02. shadow 
fn main() {
    let x = 5;

    let x = x + 1;

    let x = x*2;

    println!("The value of x is : {}", x);

    let spaces = "    ";
    let spaces = spaces.len();
    println!("The length of spaces is : {}", spaces);
}
*/


/*// 03. variable type
fn main() {
    let guess: u32 = "42".parse().expect("Not a number!");

    println!("The value of guess is : {}", guess);

    let x = 2.0;        // f64
    let y:f32 = 3.0;    // f32
    println!("x: {}; y:{}", x, y);

    let sum = 5 + 10;
    let difference = 95.5 - 4.3;
    let product = 4*30;
    let quotient = 56.7 / 32.2;
    println!("sum: {}; difference: {}; product: {}; quotient: {}",
            sum, difference, product, quotient);

    let t = true;
    let f: bool = false;
    println!("t: {};\t f: {}", t, f);

    let c = 'z';
    println!("c : {}", c);
}
*/


/*// 04. tuple and array
fn main() {
    let tup:(i32, f64, u8) = (500, 6.4, 1);
    println!("tup: {:?}", tup);

    let (x, y, z) = tup;
    println!("x: {}; y: {}; z: {}", x, y, z);

    let x:(i32, f64, u8) = (500, 6.4, 1);
    let five_hundred = x.0;
    let six_point_four = x.1;
    let one = x.2;
    println!("five_hundred: {}; six_point_four: {}; one: {}", 
            five_hundred, six_point_four, one);

    let a = [1, 2, 3, 4, 5];
    println!("a : {:?}", a);

    let months = ["January", "February", "March", "April", "May", 
        "June", "July", "August", "September", "October", "November", "December"];
    println!("months: {:?}", months);

    let a:[i32; 5] = [1, 2, 3, 4, 5];
    println!("a : {:?}", a);

    let a = [3; 5];
    println!("a : {:?}", a);    // equal : a = [3, 3, 3, 3, 3];

    let a = [1, 2, 3, 4, 5];
    let first = a[0];
    let second = a[1];
    println!("first: {}; second: {}", first, second);
}
*/


/*// 05. functions
fn main() {
    println!("Hello, world!");

    another_function(5, 6);
}

fn another_function(x: i32, y: i32) {
    println!("x: {}; y: {}", x, y);
}
*/


/*// 06. statement and expression
fn main() {
    let x = 5;

    let y = {
        let x = 3;
        x + 1
    };

    println!("x: {}; y: {}", x, y);
}
*/



// 05. for
fn main() {
    let a = [10, 20, 30, 40, 50];

    for element in a.iter() {
        println!("value is: {}", element);
    }

    for number in (1..4).rev() {
        println!("number: {}", number);
    }
    println!("LFITOFF!");
}

/*// 01. basic loop
fn main() {
    loop {
        println!("again!");
    }
}
*/

/*// 02. loop with return value
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("result: {}", result);
}
*/


/*// 03. basic while
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("number: {}", number);

        number -= 1;
    }

    println!("LIFTOFF");
}
*/


/*// 04. array while
fn main() {
    let a = [10, 20, 30, 40, 50];

    let mut index = 0;
    while index < 5 {
        println!("value is: {}", a[index]);

        index += 1;
    }
}
*/




// 03. if and let
fn main() {
    let condition = true;
    let number = if condition {
        5
    }
    else {
        6
    };
    
    println!("number: {}", number);
}


/*// 01. basic if 
fn main() {
    let number = 3;
    if number < 5 {
        println!("condition was true");
    }
    else {
        println!("condition was false");
    }
}
*/


/*// 02. else if
fn main() {
    let number = 6;

    if number%4 == 0 {
        println!("number is divisible by 4");
    }
    else if number%3 == 0 {
        println!("number is divisible by 3");
    }
    else if number%2 == 0 {
        println!("number is divisible by 2");
    }
    else {
        println!("number is not divisible by 4, 3, or 2");
    }
}
*/

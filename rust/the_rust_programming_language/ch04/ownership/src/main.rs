


// 05. change bowrrowing variable
fn main() {
    let mut s = String::from("hello");

    change(&mut s);

    println!("s: {}", s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}


/*// 04. borrowing
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
*/


/*// 03. return tuple
fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();

    (s, length)
}
*/


/*// 02. give ownership
fn main() {
    let s1 = give_onwership();

    let s2 = String::from("hello");
    let s3 = takes_and_give_back(s2);

    println!("s1: {}; s3:{}", s1, s3);
}

fn give_onwership() -> String {
    let some_string = String::from("hello");

    some_string
}

fn takes_and_give_back(a_string: String) -> String {
    a_string
}
*/


/*// 01. take ownership
fn main() {
    let s = String::from("hello");
    takes_ownership(s);

    let x = 5;
    make_copy(x);
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn make_copy(some_integer: i32) {
    println!("{}", some_integer);
}
*/
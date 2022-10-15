

// 04. slice of other type
fn main() {
    let a = [1, 2, 3, 4, 5];

    let slice = &a[1..3];
    println!("{:?}", slice);
}


/*// 03. revise first_word
fn main() {
    let s = String::from("hello world");
    let word = first_word(&s);
    println!("word: {}", word);
    let word = first_word(&s[..]);
    println!("word: {}", word);

    let s_literal = "hello world";
    let word = first_word(&s_literal[..]);   
    println!("word: {}", word);
    let word = first_word(s_literal);   //可以接收字符串字面量的切片作为参数 
    println!("word: {}", word);
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for(i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }

    &s[..]
}
*/


/*// 02. slice
fn main() {
    let s = String::from("hello world");

    let hello = &s[0..5];
    let word = &s[6..11];

    println!("hello: {}; word: {}", hello, word);

    let slice = &s[0..2];
    println!("slice: {}", slice);
    let slice = &s[..2];
    println!("slice: {}", slice);

    let len = s.len();
    let slice = &s[3..len];
    println!("slice: {}", slice);
    let slice = &s[3..];
    println!("slice: {}", slice);

    let len = s.len();
    let slice = &s[0..len];
    println!("slice: {}", slice);
    let slice = &s[..];
    println!("slice: {}", slice);
}
*/


/*// 01. first word
fn main() {
    let mut s = String::from("hello world");

    let word = first_word(&s);

    s.clear();

    println!("word: {}", word);
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for(i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}
*/
/**
 * 
 * 
 * 
'a 生命周期声明的含义

Hi, I found your answer not simple enough for Rust newbies like me. After looking around for quite some time, I'd like to offer a simpler explanation to the question: "Why is 'a there?" ---- Because the struct definition ties it to a referenced object (in this case, every struct Person instance is referencing a &str) you want to specificly declare an arbitary lifetime and tie these two things together: You want a struct Person instance to only live as long as its referenced object (hence Person<'a> and name: &'a str) so dangling references after each other's death is avoided. 

#[derive(Debug)]
struct Person<'a> {
    name: &'a str,
    age: u8
}

fn main() {
    let name = "Peter";
    let age = 27;
    let peter = Person { name, age };

    // Pretty print
    println!("{:#?}", peter);
}
 */


/**
 * 
 * refer
 * https://stackoverflow.com/questions/47640550/what-is-a-in-rust-language
 */
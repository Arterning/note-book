
// 03. The Twelve Days of Christmas
fn main() {
    let first = ["first", "second", "third", "fourth",
                 "fifth", "sixth", "seventh", "eighth",
                 "ninth", "tenth", "11th", "12th"];
    let third = ["A partridge in a pear tree", 
                 "Two turtle-doves",
                 "Three French hens",
                 "Four calling birds",
                 "Five golden rings (five golden rings)",
                 "Six geese a laying",
                 "Seven swans a swimming",
                 "Eight maids a milking",
                 "Nine ladies dancing",
                 "Ten lords a-leaping",
                 "11 pipers piping",
                 "12 drummers drumming",
                ];
    const N: u32 = 12;

    for i in 0..N {
        println!("On the {} day of Christmas", first[i as usize]);
        println!("My true love sent to me");
        for j in (0..i+1).rev() {
            println!("{}", third[j as usize]);
        }        
        println!("");
    }
}


/*// 1. temperature conversion
fn main() {
    // temperature_f_to_c
    let f = 86.0;
    let c = temperature_f_to_c(f);
    println!("temperature F({}) to C({}).", f, c);
    
    // temperature_c_to_f
    let c = 40.0;
    let f = temperature_c_to_f(c);
    println!("temperature C({}) to F({}).", c, f);
}

fn temperature_f_to_c(f: f64) -> f64 {
    let c = (f-32.0)/1.8;
    c
}

fn temperature_c_to_f(c: f64) -> f64 {
    let f = (c*1.8) + 32.0;
    f
}
*/


/*// 02. fabonacci
fn main() {
    for e in 1..10 {
        println!("fibonacci {} : {}", e, fibonacci(e));
    }
}

fn fibonacci(n: u32) -> u32 {
    if n == 0 {
        return 0;
    }
    else if n == 1 {
        return 1;
    }

    let mut i=2;
    let mut prev0 = 0;
    let mut prev1 = 1;
    let mut res = prev0 + prev1;
    while i<=n {
        i += 1;
        res = prev0 + prev1;
        prev0 = prev1;
        prev1 = res;
    }
    res
}
*/


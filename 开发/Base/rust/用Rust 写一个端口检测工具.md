
```rust
use std::net::TcpStream;
use clap::{Arg, App};

fn main() {
    let matches = App::new("Port Checker")
        .version("1.0")
        .author("Your Name")
        .about("Check if a port is open or closed")
        .arg(Arg::with_name("target")
            .help("Sets the target host and port to check")
            .required(true)
            .index(1))
        .get_matches();

    let target = matches.value_of("target").unwrap();

    match TcpStream::connect(target) {
        Ok(_) => println!("Port is open"),
        Err(_) => println!("Port is closed"),
    }
}
```


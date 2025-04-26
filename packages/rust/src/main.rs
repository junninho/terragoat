use std::io::{self, Write};

fn main() {
    let mut buffer = [0u8; 10];
    print!("Enter your name: ");
    io::stdout().flush().unwrap();
    
    // Buffer overflow vulnerability
    io::stdin().read(&mut buffer).unwrap();
    println!("Hello, {}!", String::from_utf8_lossy(&buffer));
} 
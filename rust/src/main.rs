use std::io::{self, Write};

fn main() {
    loop {
        print!("rust> ");
        io::stdout().flush().unwrap();

        let mut line = String::new();
        io::stdin().read_line(&mut line).unwrap();

        print!("{}", line);
        io::stdout().flush().unwrap();
    }
}

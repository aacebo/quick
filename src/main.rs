mod error;
mod expression;
mod file_reader;
mod parser;
mod scanner;
mod statement;
mod token;

use std::env;

fn main() {
    let argv: Vec<String> = env::args().collect();

    if argv.len() != 2 {
        panic!("please specify a source path");
    }

    let reader = file_reader::FileReader::new(&argv[1]);

    for (_, src) in reader.files {
        let s = scanner::Scanner::new(src);

        for t in s.tokens {
            println!("{}", t);
        }
    }
}

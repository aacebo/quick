mod file_reader;

use std::env;

fn main() {
    let argv: Vec<String> = env::args().collect();

    if argv.len() != 2 {
        panic!("please specify a source path");
    }

    let reader = file_reader::FileReader::new(&argv[1]);
    dbg!(argv);
    dbg!(reader);
}

use glob::glob;
use std::{collections::HashMap, path::PathBuf};

#[derive(Debug)]
pub struct FileReader {
    pub files: HashMap<PathBuf, String>,
}

impl FileReader {
    pub fn new(path: &String) -> Self {
        let mut files: HashMap<PathBuf, String> = HashMap::new();

        for entry in
            glob(path.as_str()).expect(format!("failed to read pattern \"{}\"", path).as_str())
        {
            let k = match entry {
                Ok(p) => p,
                Err(e) => panic!("{}", e),
            };

            let v = match std::fs::read_to_string(&k) {
                Ok(v) => v,
                Err(e) => panic!("{}", e),
            };

            files.insert(k, v);
        }

        return Self { files };
    }
}

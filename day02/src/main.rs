use std::collections::HashMap;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let values = reader.lines().map(|line| {
        line.unwrap_or_default().trim_end().to_owned()
    }).collect();

    println!("Checksum: {}", checksum_box_ids(values));
    Ok(())
}

#[derive(PartialEq, Debug)]
struct Characteristics {
    has_two: bool,
    has_three: bool,
}

fn analyse_box_id(s: &str) -> Characteristics {
    let mut character_counts: HashMap<char, i32> = HashMap::new();

    for ch in s.chars() {
        let counter = character_counts.entry(ch).or_insert(0);
        *counter += 1
    }

    let mut result = Characteristics { has_two: false, has_three: false };
    for n in character_counts.values() {
        if *n == 2 {
            result.has_two = true
        } else if *n == 3 {
            result.has_three = true
        }
    }

    result
}

fn checksum_box_ids(ids: Vec<String>) -> i32 {
    let mut twos = 0;
    let mut threes = 0;

    for id in ids {
        let characteristics = analyse_box_id(&id);
        if characteristics.has_two {
            twos += 1
        }
        if characteristics.has_three {
            threes += 1
        }
    }

    twos * threes
}

#[test]
fn test_box_id() {
    assert_eq!(analyse_box_id("ab"),
               Characteristics { has_two: false, has_three: false });
    assert_eq!(analyse_box_id("aabb"),
               Characteristics { has_two: true, has_three: false });
    assert_eq!(analyse_box_id("aabbb"),
               Characteristics { has_two: true, has_three: true });
    assert_eq!(analyse_box_id("aaabbb"),
               Characteristics { has_two: false, has_three: true });
    assert_eq!(analyse_box_id("aaaabbbb"),
               Characteristics { has_two: false, has_three: false });
}

#[test]
fn test_checksum() {
    assert_eq!(checksum_box_ids(vec!["abcdef".to_string(),
                                     "bababc".to_string(),
                                     "abbcde".to_string(),
                                     "abcccd".to_string(),
                                     "aabcdd".to_string(),
                                     "abcdee".to_string(),
                                     "ababab".to_string()]), 12)
}

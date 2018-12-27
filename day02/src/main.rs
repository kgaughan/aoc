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

    println!("Checksum: {}", checksum_box_ids(&values));

    if let Some(common) = scan_for_consecutive(&values) {
        println!("Common: {}", common)
    } else {
        println!("No consecutive IDs found.")
    }
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

fn checksum_box_ids(ids: &Vec<String>) -> i32 {
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

fn check_if_consecutive(s1: &str, s2: &str) -> Option<String> {
    let mut common = "".to_string();
    let mut different = 0;

    for (ch1, ch2) in s1.chars().zip(s2.chars()) {
        if ch1 == ch2 {
            common.push(ch1);
        } else {
            different += 1;
        }
    }

    if different != 1 {
        None
    } else {
        Some(common)
    }
}

fn scan_for_consecutive(ids: &Vec<String>) -> Option<String> {
    for (i, id1) in ids.iter().enumerate() {
        if i < ids.len() - 1 {
            for j in (i + 1)..ids.len() {
                if let Some(common) = check_if_consecutive(&id1, &ids[j]) {
                    return Some(common)
                }
            }
        }
    }
    None
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
                                     "ababab".to_string()].as_ref()), 12)
}

#[test]
fn test_consecutive() {
    assert_eq!(check_if_consecutive("abcde", "axcye"), None);
    assert_eq!(check_if_consecutive("fghij", "fguij"), Some("fgij".to_string()));
}

#[test]
fn test_scan_consecutive() {
    assert_eq!(scan_for_consecutive(vec!["abcde".to_string(),
                                         "fghij".to_string(),
                                         "klmno".to_string(),
                                         "pqrst".to_string(),
                                         "fguij".to_string(),
                                         "axcye".to_string(),
                                         "wvxyz".to_string()].as_ref()),
               Some("fgij".to_string()))
}

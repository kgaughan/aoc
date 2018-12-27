use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
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

fn checksum_box_ids(ids: Vec<&str>) -> i32 {
    let mut twos = 0;
    let mut threes = 0;

    for id in ids {
        let characteristics = analyse_box_id(id);
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
    assert_eq!(checksum_box_ids(vec!["abcdef",
                                     "bababc",
                                     "abbcde",
                                     "abcccd",
                                     "aabcdd",
                                     "abcdee",
                                     "ababab"]), 12)
}

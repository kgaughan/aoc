fn main() {
    println!("Hello, world!");
}

#[derive(PartialEq, Debug)]
struct Characteristics {
    has_two: bool,
    has_three: bool,
}

fn analyse_box_id(_s: &str) -> Characteristics {
    return Characteristics {
        has_two: false,
        has_three: false,
    }
}

#[test]
fn tests_box_id() {
    assert_eq!(analyse_box_id("aabbb"),
               Characteristics { has_two: true, has_three: false })
}

#[macro_use] extern crate text_io;

use std::cmp;

fn main() {
    println!("Hello, world!");
}

#[derive(PartialEq, Debug)]
struct Rect {
    id: String,
    x: i32,
    y: i32,
    w: i32,
    h: i32,
}

impl Rect {
    fn new(id: &str, x: i32, y: i32, w: i32, h: i32) -> Rect {
        Rect { id: id.to_string(), x: x, y: y, w: w, h: h }
    }

    fn bottom_right(&self) -> (i32, i32) {
        (self.x + self.w, self.y + self.h)
    }

    fn overlap(&self, other: &Rect) -> Option<i32> {
        let (x2, y2) = self.bottom_right();
        let (x4, y4) = other.bottom_right();

        let x5 = cmp::max(self.x, other.x);
        let y5 = cmp::max(self.y, other.y);
        let x6 = cmp::min(x2, x4);
        let y6 = cmp::min(y2, y4);

        // Degenerate triangle: no overlap
        if x5 >= x6 || y5 >= y6 {
            None
        } else {
            Some((x6 - x5) * (y6 - y5))
        }
    }
}

fn parse_rect(s: &str) -> Rect {
    let id: String;
    let x: i32;
    let y: i32;
    let w: i32;
    let h: i32;

    scan!(s.bytes() => "#{} @ {},{}: {}x{}",
          id, x, y, w, h);

    Rect::new(&id, x, y, w, h)
}

#[test]
fn test_bottom_right() {
    assert_eq!(Rect::new("", 0, 0, 1, 1).bottom_right(),
               (1, 1));
    assert_eq!(Rect::new("", 7, 5, 1, 1).bottom_right(),
               (8, 6));
    assert_eq!(Rect::new("", 12, 8, 8, 2).bottom_right(),
               (20, 10));
}

#[test]
fn test_no_overlap() {
    let r1 = Rect::new("", 0, 0, 1, 1);
    let r2 = Rect::new("", 1, 1, 3, 3);
    assert_eq!(r1.overlap(&r2), None);
    assert_eq!(r2.overlap(&r1), None);
}

#[test]
fn test_overlaps() {
    // Overlap from TL <-> BR
    let r1 = Rect::new("", 5, 5, 5, 5);
    let r2 = Rect::new("", 7, 7, 5, 5);
    assert_eq!(r1.overlap(&r2), Some(9));
    assert_eq!(r2.overlap(&r1), Some(9));

    // Overlap from BL <-> TR
    let r3 = Rect::new("", 9, 5, 5, 5);
    assert_eq!(r2.overlap(&r3), Some(9));
    assert_eq!(r3.overlap(&r2), Some(9));

    // Fully contained
    let r4 = Rect::new("", 8, 8, 2, 2);
    assert_eq!(r2.overlap(&r4), Some(4));
}

#[test]
fn test_parse() {
    assert_eq!(parse_rect("#1 @ 1,3: 4x4"),
               Rect::new("1", 1, 3, 4, 4));
    assert_eq!(parse_rect("#2 @ 3,1: 4x4"),
               Rect::new("2", 3, 1, 4, 4));
    assert_eq!(parse_rect("#3 @ 5,5: 2x2"),
               Rect::new("3", 5, 5, 2, 2));
}

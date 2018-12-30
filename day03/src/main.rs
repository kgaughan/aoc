#[macro_use] extern crate text_io;

use std::cmp;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::vec;

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    // .sort() further down requres this be mutable.
    let mut rects: vec::Vec<Rect> = reader.lines().map(|line| {
        parse_rect(line.unwrap_or_default().as_str())
    }).collect();
    // Has to be separate to keep the type inference engine happy.
    rects.sort();

    println!("Total overlap: {}", sum_overlaps(&rects));

    Ok(())
}

#[derive(PartialEq, Eq, Ord, Debug)]
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

        // Degenerate rectangle: no overlap
        if x5 >= x6 || y5 >= y6 {
            None
        } else {
            Some((x6 - x5) * (y6 - y5))
        }
    }
}

impl PartialOrd for Rect {
    fn partial_cmp(&self, other: &Rect) -> Option<cmp::Ordering> {
        // Enough ordering for sorting the rectangles for a plane sweep
        // (https://en.wikipedia.org/wiki/Sweep_line_algorithm)
        Some(self.x.cmp(&other.x))
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

fn sum_overlaps(rects: &Vec<Rect>) -> i32 {
    let mut total = 0;
    for (i, r1) in rects.iter().enumerate() {
        if i < rects.len() - 1 {
            for j in (i + 1)..rects.len() {
                if let Some(overlap) = r1.overlap(&rects[j]) {
                    total += overlap
                }
            }
        }
    }
    total
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

#[test]
fn test_sum() {
    let rects = vec![Rect::new("1", 1, 3, 4, 4),
                     Rect::new("2", 3, 1, 4, 4),
                     Rect::new("3", 5, 5, 2, 2)];
    assert_eq!(sum_overlaps(&rects), 4)
}

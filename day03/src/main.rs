#[macro_use] extern crate text_io;

use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::vec;

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let rects: vec::Vec<Rect> = reader.lines().map(|line| {
        parse_rect(line.unwrap_or_default().as_str())
    }).collect();

    println!("Total overlap: {}", sum_overlaps(&rects));

    Ok(())
}

#[derive(PartialEq, Debug)]
struct Rect {
    id: String,
    x: usize,
    y: usize,
    w: usize,
    h: usize,
}

impl Rect {
    fn new(id: &str, x: usize, y: usize, w: usize, h: usize) -> Rect {
        Rect { id: id.to_string(), x: x, y: y, w: w, h: h }
    }

    fn top_left(&self) -> (usize, usize) {
        (self.x, self.y)
    }

    fn bottom_right(&self) -> (usize, usize) {
        (self.x + self.w, self.y + self.h)
    }
}

fn parse_rect(s: &str) -> Rect {
    let id: String;
    let x: usize;
    let y: usize;
    let w: usize;
    let h: usize;

    scan!(s.bytes() => "#{} @ {},{}: {}x{}",
          id, x, y, w, h);

    Rect::new(&id, x, y, w, h)
}

fn sum_overlaps(rects: &Vec<Rect>) -> i32 {
    // I'm not happy with doing this using a brute-force mechanism such as
    // this, but I've faffed around long enough. Here's the proper way to
    // solve it:
    //
    //     http://codercareer.blogspot.com/2011/12/no-27-area-of-rectangles.html
    //
    // There's code in the previous commit that would be needed to implement
    // that algorithm.

    // Plot the rectangles. 0 = empty; 1 = no overlap; 2 = overlap
    let mut sheet: [[u8; 1000]; 1000] = [[0; 1000]; 1000];
    for rect in rects.iter() {
        let (x1, y1) = rect.top_left();
        let (x2, y2) = rect.bottom_right();

        // This probably isn't idiomatic.
        for x in x1..x2 {
            for y in y1..y2 {
                if sheet[y][x] < 2 {
                    sheet[y][x] += 1;
                }
            }
        }
    }

    // Count the overlaps. Probably not idiomatic.
    let mut total = 0;
    for x in 0..1000 {
        for y in 0..1000 {
            if sheet[y][x] == 2 {
                total += 1;
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

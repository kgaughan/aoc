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

    let mut sheet: [[u8; 1000]; 1000] = [[0; 1000]; 1000];
    plot_rectangles(&rects, &mut sheet);

    println!("Total overlap: {}", count_overlaps(&sheet));
    println!("First without overlap: {}", find_first_without_overlap(&rects, &sheet).unwrap());

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

fn plot_rectangles(rects: &Vec<Rect>, sheet: &mut [[u8; 1000]; 1000]) -> () {
    // Plot the rectangles. 0 = empty; 1 = no overlap; 2 = overlap
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
}

fn count_overlaps(sheet: &[[u8; 1000]; 1000]) -> i32 {
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

fn find_first_without_overlap(rects: &Vec<Rect>, sheet: &[[u8; 1000]; 1000]) -> Option<String> {
    // Plot the rectangles. 0 = empty; 1 = no overlap; 2 = overlap
    for rect in rects.iter() {
        let (x1, y1) = rect.top_left();
        let (x2, y2) = rect.bottom_right();

        let mut overlapping = false;
        'outer: for x in x1..x2 {
            for y in y1..y2 {
                if sheet[y][x] == 2 {
                    overlapping = true;
                    break 'outer;
                }
            }
        }
        if !overlapping {
            return Some(rect.id.clone());
        }
    }
    None
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

    let mut sheet: [[u8; 1000]; 1000] = [[0; 1000]; 1000];
    plot_rectangles(&rects, &mut sheet);

    assert_eq!(count_overlaps(&sheet), 4)
}

#[test]
fn test_no_overlap() {
    let rects = vec![Rect::new("1", 1, 3, 4, 4),
                     Rect::new("2", 3, 1, 4, 4),
                     Rect::new("3", 5, 5, 2, 2)];

    let mut sheet: [[u8; 1000]; 1000] = [[0; 1000]; 1000];
    plot_rectangles(&rects, &mut sheet);

    assert_eq!(find_first_without_overlap(&rects, &sheet).unwrap(), "3")
}

use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;

struct Point(i32, i32);

struct Line {
    p1: Point,
    p2: Point,
}

enum Orientation {
    Horizontal,
    Vertical,
}

impl Line {

    fn orientation(&self) -> Orientation {
        if self.p1[0] == self.p2[0] {
            Horizontal
        } else if self.p1[1] == self.p2[1] {
            Vertical
        } else {
            panic!("WAT")
        }
    }

    fn intersects(&self, other: &Point) -> Option<Point> {
        // This is a simplified intersection algorithm that assume lines are
        // on a grid and do not overlap.

        // Lines on the same axis are assumed to not overlap:
        if self.orientation == other.orientation {
            return None;
        }
        if self.orientation == Vertical {

        } else {
            return other.intersects(self);
        }
    }
}

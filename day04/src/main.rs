extern crate pest;
#[macro_use]
extern crate pest_derive;

use pest::Parser;
use std::fs;

#[derive(Parser)]
#[grammar = "grammar.pest"]
pub struct GuardParser;

// The state of the guard.
enum State {
    StartShift(i32),
    FallAsleep,
    WakeUp,
}

struct Record {
    year: u16,
    month: u8,
    day: u8,
    hour: u8,
    minute: u8,
    state: State,
}

fn main() {
    let contents = fs::read_to_string("input.txt").expect("cannot read file");
    let file = GuardParser::parse(Rule::file, &contents)
        .expect("unsuccessful parse")
        .next().unwrap();

    let mut records: Vec<Record> = Vec::new();
    for line in file.into_inner() {
        match line.as_rule() {
            Rule::record => {
                let mut record = Record {
                    year: 0,
                    month: 0,
                    day: 0,
                    hour: 0,
                    minute: 0,
                    state: State::FallAsleep,
                };
                for record_field in line.into_inner() {
                    match record_field.as_rule() {
                        Rule::timestamp => {
                            for ts_field in record_field.into_inner() {
                                match ts_field.as_rule() {
                                    Rule::year => {
                                        record.year = ts_field.as_str().parse().unwrap();
                                    }
                                    Rule::month => {
                                        record.month = ts_field.as_str().parse().unwrap();
                                    }
                                    Rule::day => {
                                        record.day = ts_field.as_str().parse().unwrap();
                                    }
                                    Rule::hour => {
                                        record.hour = ts_field.as_str().parse().unwrap();
                                    }
                                    Rule::minute => {
                                        record.minute = ts_field.as_str().parse().unwrap();
                                    }
                                    _ => unreachable!(),
                                }
                            }
                        }
                        Rule::start_shift => {
                            for state_field in record_field.into_inner() {
                                match state_field.as_rule() {
                                    Rule::id => {
                                        let id: i32 = state_field.as_str().parse().unwrap();
                                        record.state = State::StartShift(id);
                                    }
                                    _ => unreachable!(),
                                }
                            }
                        }
                        Rule::fall_asleep => {
                            record.state = State::FallAsleep;
                        }
                        Rule::wake_up => {
                            record.state = State::WakeUp;
                        }
                        _ => unreachable!(),
                    }
                }
                records.push(record);
            }
            Rule::EOI => (),
            _ => unreachable!(),
        }
    }
}

use std::collections::HashSet;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;

fn load_intcode() -> io::Result<Vec<usize>> {
    let f = File::open("input.txt")?;

    let mut contents = String::new();
    BufReader::new(f).read_to_string(&mut contents).expect("couldn't read file");

    let intcode = contents
        .split(',')
        .map(|code| code.trim().parse::<usize>().unwrap())
        .collect();

    Ok(intcode)
}

fn execute(intcode: &mut Vec<usize>) {
    let mut pc = 0;
    loop {
        match intcode[pc] {
            1 | 2 => {
                let i_op1 = intcode[pc + 1];
                let i_op2 = intcode[pc + 2];
                let i_res = intcode[pc + 3];

                intcode[i_res] = match intcode[pc] {
                    1 => intcode[i_op1] + intcode[i_op2],
                    2 => intcode[i_op1] * intcode[i_op2],
                    _ => panic!("WAT"),
                };
                pc += 4;
            },
            99 => break,
            _ => panic!("Illegal opcode: {}", intcode[pc]),
        }
    }
}

const P2_TARGET: usize = 19690720;

fn main() -> io::Result<()> {
    let intcode = load_intcode().unwrap();

    // Part one:
    {
        let mut intcode_dup = intcode.clone();

        // Restore 1202 program alarm
        intcode_dup[1] = 12;
        intcode_dup[2] = 2;

        execute(&mut intcode_dup);

        println!("Part 1: {}", intcode_dup[0]);
    }

    // Part 2:
    // This bruteforces the solution, but, as an optimisation, will skip
    // anything where the target value of the noun and verb has been seen
    // before.
    let mut nouns_seen = HashSet::new();
    'outer: for noun in 0..100 {
        if nouns_seen.contains(&intcode[noun]) {
            continue;
        }
        nouns_seen.insert(intcode[noun]);

        let mut verbs_seen = HashSet::new();
        for verb in 0..100 {
            if verbs_seen.contains(&intcode[verb]) {
                continue;
            }
            verbs_seen.insert(intcode[verb]);

            let mut intcode_dup = intcode.clone();

            intcode_dup[1] = noun;
            intcode_dup[2] = verb;
            execute(&mut intcode_dup);

            if intcode_dup[0] == P2_TARGET {
                println!("Part 2: {}", 100 * intcode_dup[1] + intcode_dup[2]);
                break 'outer;
            }
        }
    }

    Ok(())
}

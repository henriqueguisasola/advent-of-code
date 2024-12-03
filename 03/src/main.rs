use std::fs::File;
use std::io::prelude::*;
use regex::Regex;

fn part1() {
    let mut sum = 0;
    let mut f = File::open("./src/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let parts = contents.split("mul(");

    for part in parts {
        let collection = part.split(")").collect::<Vec<&str>>()[0];

        let re = Regex::new(r"^\d{1,3},\d{1,3}$").unwrap();
        if re.is_match(collection) {
            let numbers = collection.split(",").collect::<Vec<&str>>();
            sum += numbers[0].parse::<i32>().unwrap() * numbers[1].parse::<i32>().unwrap();
        }
    }

    println!("{}", sum);
}

fn part2() {
    let mut sum = 0;
    let mut is_doing = true;
    let mut f = File::open("./src/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let parts = contents.split("mul(");

    for part in parts {
        if is_doing {
            let collection = part.split(")").collect::<Vec<&str>>()[0];

            let re = Regex::new(r"^\d{1,3},\d{1,3}$").unwrap();
            if re.is_match(collection) {
                let numbers = collection.split(",").collect::<Vec<&str>>();
                sum += numbers[0].parse::<i32>().unwrap() * numbers[1].parse::<i32>().unwrap();
            }
        }

        if part.contains("do()") {
            is_doing = true;
        } else if part.contains("don't()") {
            is_doing = false;
        }
    }

    println!("{}", sum);
}

fn main() {
    part1();
    part2();
}
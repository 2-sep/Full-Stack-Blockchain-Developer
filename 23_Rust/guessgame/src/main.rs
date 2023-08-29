use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    println!("The secret number is: {secret_number}");

    // 循环
    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        // 标准输入
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        // 类型转换 新值隐藏
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            // _ 通配符值
            Err(_) => continue,
        };

        println!("You guessed: {guess}");

        // match 与 分支
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}

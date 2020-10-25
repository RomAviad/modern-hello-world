#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
#[macro_use]extern crate serde;
extern crate rocket_contrib;
extern crate ngrams;

use rocket_contrib::json::Json;
use serde::{Deserialize};
use ngrams::Ngram;

const TEXT_LIMIT: u64 = 1024;

#[derive(Deserialize)]
struct NgramRequest<'a> {
    data: &'a str,
    n: usize,
}



#[get("/")]
fn hello() -> String {
    let my_result = "Hello World";
//    let tokens: Vec<String> = my_result.split(" ").map(|s| s.to_string()).collect();
    format!("{}", my_result)
}

#[post("/ngram", format = "application/json", data = "<body>")]
fn ngrams(body: Json<NgramRequest>) -> String {
    let tokens = body.data.split(" ");
    let n: usize = body.n;
    let result: Vec<_> = tokens.ngrams(n).collect();

    format!("{:?}", result)
}

fn main() {
    rocket::ignite().mount("/", routes![hello, ngrams]).launch();
}
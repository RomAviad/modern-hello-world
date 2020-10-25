#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
extern crate serde;
extern crate rocket_contrib;
extern crate ngrams;

use ngrams::Ngram;
use rocket::response::content;
use rocket_contrib::json::Json;
use serde::{Deserialize};

const TEXT_LIMIT: u64 = 1024;

#[derive(Deserialize)]
struct NgramRequest<'a> {
    data: &'a str,
    n: usize,
}



#[get("/")]
fn hello() -> String {
    format!("Hello World")
}

#[post("/ngram", format = "application/json", data = "<body>")]
fn ngrams(body: Json<NgramRequest>) -> content::Json<String> {
    let tokens = body.data.split(" ");
    let n: usize = body.n;
    let result: Vec<_> = tokens.ngrams(n).collect();
    let result_str = format!("{:?}", result);
    content::Json(result_str)
}

fn main() {
    rocket::ignite().mount("/", routes![hello, ngrams]).launch();
}
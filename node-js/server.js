const natural = require("natural");
const http = require("http");
const express = require("express");
const bodyParser = require("body-parser");

const app = express();
app.use(bodyParser());

app.get("/", function(req, res) {
	res.send("Hello World!");
});


app.post("/ngrams", function(req, res) {
	var data = req.body["data"];
	var n = parseInt(req.body["n"]);
	if (!n) {
		n = 2;
	}
	console.log(n);	
	var tokenizer = new natural.WordTokenizer();
	var NGrams = natural.NGrams;
	var tokens = tokenizer.tokenize(data);
        res.send(JSON.stringify(NGrams.ngrams(tokens, n)));
});

app.listen(8080);


natural = require("natural");
http = require("http");

var tokenizer = new natural.WordTokenizer();

var server = http.createServer(function (req, res) {
	res.end(JSON.stringify(tokenizer.tokenize("Your dog has fleas.")));
});

server.listen(8080);

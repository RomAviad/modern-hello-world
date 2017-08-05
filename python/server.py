from flask import Flask, request, jsonify


app = Flask(__name__)


def ngrams(input_text, n):
	tokens = input_text.split()
	return list(zip(*[tokens[i:] for i in range(n)]))


@app.route("/")
def hello_world():
	return "Hello World"


@app.route("/ngrams", methods=["POST"])
def hello_ngrams():
	body = request.json
	text = body.get("data")
	n = int(body.get("n", 2))
	return jsonify(ngrams(text, n))


if __name__ == '__main__':
	app.run("0.0.0.0", 8080)

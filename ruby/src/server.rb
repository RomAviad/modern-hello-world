require "sinatra"

get "/" do
  "Hello World!"
end

post "/ngrams" do
  payload = JSON.parse request.body.read
  input_text = payload["data"]
  n = payload["n"]
end
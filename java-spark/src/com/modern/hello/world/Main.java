package com.modern.hello.world;

import com.google.gson.*;
import spark.Request;
import spark.Response;

import static spark.Spark.*;

public class Main {
    public static void main(String[] args) {
        port(8080);
        get("/", (req, res) -> "Hello World!");
        post("/ngrams", "application/json", (Request req, Response res) -> {
            res.type("application/json");
            String body = req.body();
            JsonElement jsonElement = new JsonParser().parse(body);
            JsonObject jsonBody = jsonElement.getAsJsonObject();
            String data = jsonBody.get("data").getAsString();
            JsonElement nObj = jsonBody.get("n");
            int n = 2;
            if (nObj != null) {
                n = nObj.getAsInt();
            }

            String tokens[] = data.split(" ");
            JsonArray result = new JsonArray();
            for (int i = 0; i < tokens.length - n + 1; i++) {
                JsonArray ngram = new JsonArray();
                for (int j = i; j < i + n; j++) {
                    ngram.add(tokens[j]);
                }
                result.add(ngram);
            }

            // parse json body with GSON module
            // do the same trick I did with all other languages :)
            return result;
        });
    }
}

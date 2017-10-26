# Modern Hello-World
This repository will hold my conception of how a "hello world" should look.

For many years, our "introduction code" to new programming languages has been the `"Hello World!"` kind of program, in which you print a simple constant string to the screen.
This so-called "introduction code" never helped anyone to get a hold of how coding in a new language really feels, and is often skipped.

With all respect to tradition, I figured the first time you touch a new programming language should be when you implement something you might encounter in the real world. And so it goes with my coneption:
* Run an HTTP server
* Serve `Hello World` from the root route (`http://<server_address>:<port>/`)
* Serve a `POST` request that returns the N-grams of a given string (provided in JSON payload, along with N) 

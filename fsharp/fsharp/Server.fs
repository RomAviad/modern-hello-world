module HelloWorldServer = 
    open System

    open Nancy
    open Nancy.Hosting.Self
    open Nancy.Extensions
    open Newtonsoft.Json
    
    type NgramsPayload = {data: String; N: Int64}

    let StringSplit(text: String) = 
        let tokens = text.Split()
        (tokens)

    let rec MatchPairs tokensList =
        match tokensList with
        | [] | [_] -> []
        | x1::x2::rest -> String.Join(" ", (x1, x2)) :: (MatchPairs rest)

    type HelloWorldModule() as this = 
        inherit NancyModule()
        do
            this.Get.["/"] <- fun _ -> 
                "Hello World!" :> obj
            
            this.Post.["/ngrams"] <- fun _ ->
                let bodyString = this.Request.Body.AsString() 
                
                let bodyJson = JsonConvert.DeserializeObject<NgramsPayload>(bodyString)
                let text = bodyJson.data
                text |> StringSplit |> Array.toList |> MatchPairs :> obj
                //text :> obj
                //"Hello POST WORLD" :> obj

    [<EntryPoint>]
    let main args =
        let nancyHost = new NancyHost(new Uri("http://localhost:8080/"), new Uri("http://127.0.0.1:8080/"))
        nancyHost.Start()
        printfn "ready..."
        Console.ReadKey()
        nancyHost.Stop()
        0

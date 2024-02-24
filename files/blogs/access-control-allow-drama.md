First of all, just wanna point out the importance of reading a documentation before you start hacking around, and smacking your head against the wall.

So I was working on [Rex](https://github.com/mbaraa/rex) which is a GitHub action's server that deploys an application from a git repository after a push or merging a pull request (depending on your configuration of the action). So what I was doing is just fixing its CORS more specifically is the `Access-Control-Allow-Origin` header, so that only GitHub can make the request (to avoid excessive requests on my server) and here where the drama starts.

I have written a lot of Go backends in [Fiber](https://gofiber.io) and how I handled CORS is by providing a comma separated list of allowed origins, e.g. `http://localhost:8080,https://mbaraa.com` more [here](https://docs.gofiber.io/api/middleware/cors#config), so what I thought is that this is how the header works, and I kept thinking that since I was still using Fiber, but for Rex it was a ~100 lines script, so using Fiber is a bloat, and I went with Go's `net/http`, but whenever I make a request I get the CORS error, that the requesting origin is not allowed, I thought that it needed a space after the commas, or before, or even both, but non of them worked, so after some digging, a friend of mine pointed out that I should check [this](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) MDN post, more specifically the [allowed values part](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin#syntax).

After trying back and forth for like an hour, it turned out that this `http://localhost:8080,https://mbaraa.com` is one of Fiber's syntactic sugars, and that they have handled it in a way that checks for the origin if it's in the allow list on each request [here](https://github.com/gofiber/fiber/blob/6ecd607d9717b3312e3bd0c2da5194bdba78ff00/middleware/cors/cors.go#L126).

So that's exactly what I did, I had an environmental variable containing the allow list, and checked for the origin on each request.

I'll just slap Rex's `main.go` without the deployment stuff, and go from it

```go
package main

import (
	"flag"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	allowedOrigins    string
	allowedOriginsMap = map[string]bool{}
)

func main() {
	flag.StringVar(&allowedOrigins, "allowed-origins", os.Getenv("REX_ALLOWED_ORIGINS"), "give me a list of allowed origins")
	parseAllowedOringins()
	http.HandleFunc("/deploy/", handleDeployRepo)
	http.ListenAndServe(":8080", nil)
}

func parseAllowedOringins() {
	allowedOriginsList := strings.Split(
		// this regex is only to check if the comma has a trailing or a leading whitespace(s)
		regexp.MustCompile(`\s*,\s*`).ReplaceAllString(allowedOrigins, ","),
		",",
	)
	for _, allowedOrigin := range allowedOriginsList {
		allowedOriginsMap[allowedOrigin] = true
	}
}

func handleDeployRepo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if origin := req.Header.Get("Origin"); allowedOriginsMap[origin] || allowedOriginsMap["*"] {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Write([]byte("ok"))
}
```

Here what's going on is that I'm taking the comma separated list of allowed origins, put them into a check map, so that the accessing time is constant or logarithmic, setting the origin value on each request (if it exits in the allow list), and continuing with the request happily.

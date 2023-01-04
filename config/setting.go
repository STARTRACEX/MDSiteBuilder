package config

import "flag"

var Port = "8080"
var Domain = "localhost"
var Dev = false
const Cat=".yml"
const DefaultLang = "en"
const HtmlGlobPath = "views/**/*.html"
const StaticPath = "/public"
const PostsPath = "/posts"
const PostsURL = "/docs"
const ExpirationTime = 2592000

func init() {
	flag.StringVar(&Port, "p", Port, ":Port")
	flag.StringVar(&Domain, "d", Domain, "Domain")
	flag.BoolVar(&Dev, "dev", Dev, "Development mode")
	flag.Parse()
}

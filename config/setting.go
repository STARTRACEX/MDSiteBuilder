package config

import "flag"

var Port = "8080"
var Domain = "localhost"

const HtmlGlobPath = "views/**/*.html"
const StaticPath = "/public"
const PostsPath = "/posts"
const PostDefaultdPath = "/default"
const ExpirationTime = 2592000

func init() {
	flag.StringVar(&Port, "p", Port, ":Port")
	flag.StringVar(&Domain, "d", Domain, "Domain")
	flag.Parse()
}

package main

import (
	"fuzzywookie/foobot/agent"
	"fuzzywookie/foobot/proto"
	"fuzzywookie/foobot/conf"
)


func main() {
    conf.Init()
    conf.Set("irc.ident", "foobot")
    conf.Set("irc.name", "foobot")
    conf.Set("irc.version", "foobot 1.0")
    conf.Set("irc.pass", "baltycka")
    conf.Set("irc.server", "cube.mdevel.net:6697")

    irc := proto.NewIrcProto()
    a := agent.NewAgent()
    a.Attach(irc)
    a.Run()

}

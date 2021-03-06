package main

import (
    "fmt"
    "flag"
    "crypto/sha256"
    "github.com/VividCortex/godaemon"
    "github.com/mduszyk/foobot/log"
    "github.com/mduszyk/foobot/conf"
    "github.com/mduszyk/foobot/bot"
    "github.com/mduszyk/foobot/protoimpl"
)

var verbose *bool = flag.Bool("v", false, "Prints logs to stderr on trace level")
var pass *string = flag.String("P", "", "Set custom bot pass")
var ircServer *string = flag.String("s", "irc.example.com:6697", "irc server socket")
var ircPass *string = flag.String("p", "", "irc server password")

const VERSION = "foobot 1.0"

func main() {
    flag.Parse()

    if *verbose {
        log.EnableStderr()
        log.SetLevel(log.LEVEL_TRACE)
    } else {
        godaemon.MakeDaemon(&godaemon.DaemonAttr{})
    }

    if *pass != "" {
        *pass = fmt.Sprintf("%x", sha256.Sum256([]byte(*pass)))
    }

    conf.Init()
    conf.Set("irc.channel", "#bot")
    conf.Set("irc.ident", "foobot")
    conf.Set("irc.name", "foobot")
    conf.Set("irc.pass", *ircPass)
    conf.Set("irc.server", *ircServer)
    conf.Set("irc.version", VERSION)
    conf.Set("irc.ssl", "true")
    conf.Set("irc.ssl.noverify", "true")
    conf.Set("net.server.type", "tcp")
    conf.Set("net.server.socket", "localhost:6600")
    conf.Set("bot.cmd.buflen", "10")
    conf.Set("bot.worker.timeout", "120")
    conf.Set("bot.pass", *pass)
    conf.Set("bot.shell", "/bin/bash")

    ircProto := protoimpl.NewIrcProto()
    netServerProto := protoimpl.NewNetServerProto()

    a := bot.NewBot()

    a.AddModule(":conf", conf.GetConfModule())
    a.AddModule(":irc", ircProto)
    a.AddModule(":info", bot.NewInfoModule())
    a.AddModule(":log", log.GetLogModule())
    a.AddModule(":sh", bot.NewShellModule())
    a.AddModule(":auth", bot.NewAuthModule())
    a.AddModule(":run", bot.NewRunModule())

    a.AddProto("irc", ircProto)
    a.AddProto("net", netServerProto)

    a.StartProto("net")

    a.Run()
}

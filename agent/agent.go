package agent

import (
    "reflect"
	"fuzzywookie/foobot/log"
	"fuzzywookie/foobot/proto"
)

type Agent struct {
    protos map[string]proto.Proto
    modules map[string]proto.Interpreter
    proto proto.Proto
}

func NewAgent() *Agent {
	agent := &Agent{
        proto: nil,
        protos: make(map[string]proto.Proto),
        modules: make(map[string]proto.Interpreter),
    }
    return agent
}

func (agent *Agent) AddProto(name string, proto proto.Proto) {
    proto.Register(agent)
    agent.protos[name] = proto
    if agent.proto == nil {
        agent.proto = proto
    }
    log.INFO.Printf("Added proto, name: %s, type: %s", name, reflect.TypeOf(proto))
}

func (agent *Agent) AddModule(cmd string, module proto.Interpreter) {
    agent.modules[cmd] = module
    log.INFO.Printf("Added module, cmd: %s, type: %s", cmd, reflect.TypeOf(module))
}

func (agent *Agent) Run() {
    log.INFO.Printf("Starting agent")
    // run default proto
    agent.proto.Run()
}

func (agent *Agent) Handle(msg *proto.Msg) string {
    log.TRACE.Printf("Agent cmd, msg: %s", msg.Raw)

    rsp := ""

    module, ok := agent.modules[msg.Cmd]
    if ok {
        rsp = module.Handle(msg)
    }

    return rsp
}

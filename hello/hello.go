package main

import (
    "fmt"
    "log"

    "example.com/greetings"
    "github.com/hypebeast/go-osc/osc"

)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)

    client := osc.NewClient("localhost", 3030)



    addr := "127.0.0.1:9000"
    d := osc.NewStandardDispatcher()
    _ = d.AddMsgHandler("/track/1/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/1/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/2/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/2/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/3/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/3/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/4/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/4/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/5/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/5/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/6/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/6/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/7/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/7/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })
    _ = d.AddMsgHandler("/track/8/volume", func(msg *osc.Message) {
        //osc.PrintMessage(msg)

        replacementAddress := "/lx/mixer/channel/8/fader"
        incomingInteger := msg.Arguments[0].(int32)
        outgoingFloat := float64(incomingInteger)/128

        rmsg := osc.NewMessage(replacementAddress)
        rmsg.Append(outgoingFloat)
        _ = client.Send(rmsg)
    })

    server := &osc.Server{
        Addr: addr,
        Dispatcher:d,
    }

    for {
        err := server.ListenAndServe()
        if err != nil {
            log.Println("top::", err)
        }
    }
}

//Demonstrate Command Design Pattern in Golang with Unit Tests


package main

import (
    "fmt"
    "testing"
)

// Command interface
type Command interface {
    Execute() string
}

// Concrete command to turn on a device
type TurnOnCommand struct {
    device Device
}

func (c *TurnOnCommand) Execute() string {
    return c.device.On()
}

// Concrete command to turn off a device
type TurnOffCommand struct {
    device Device
}

func (c *TurnOffCommand) Execute() string {
    return c.device.Off()
}

// Device interface
type Device interface {
    On() string
    Off() string
}

// Concrete device (Light)
type Light struct {
}

func (l *Light) On() string {
    return "Light is turned on"
}

func (l *Light) Off() string {
    return "Light is turned off"
}

// RemoteControl acts as an invoker
type RemoteControl struct {
    command Command
}

func (r *RemoteControl) SetCommand(command Command) {
    r.command = command
}

func (r *RemoteControl) PressButton() string {
    return r.command.Execute()
}

// Unit tests
func TestTurnOnCommand(t *testing.T) {
    light := &Light{}
    turnOnCommand := &TurnOnCommand{device: light}
    remote := &RemoteControl{}

    remote.SetCommand(turnOnCommand)
    result := remote.PressButton()

    expected := "Light is turned on"
    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

func TestTurnOffCommand(t *testing.T) {
    light := &Light{}
    turnOffCommand := &TurnOffCommand{device: light}
    remote := &RemoteControl{}

    remote.SetCommand(turnOffCommand)
    result := remote.PressButton()

    expected := "Light is turned off"
    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

func main() {
    light := &Light{}
    turnOnCommand := &TurnOnCommand{device: light}
    turnOffCommand := &TurnOffCommand{device: light}
    remote := &RemoteControl{}

    remote.SetCommand(turnOnCommand)
    fmt.Println(remote.PressButton())

    remote.SetCommand(turnOffCommand)
    fmt.Println(remote.PressButton())
}

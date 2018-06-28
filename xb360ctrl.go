// Package joystick provides functions for monitoring and updating
// the state of an Xbox 360 controller.
package xb360ctrl

import "fmt"

/*
#include <linux/joystick.h>
#include <fcntl.h>
#include <unistd.h>
int jsOpen(char* f) {
  return open(f, O_RDONLY);
}
int jsClose(int fd) {
  return close(fd);
}
void jsRead(int js, struct js_event* e) {
  ssize_t r=read(js, e, sizeof(struct js_event));
}
*/
import "C"

var DEADZONE int16 = 200

// Js_even is a Go analog to the C Joystick struct
type Xbc_event struct {
	Time      uint32
	Value     int16
	EventType uint8
	Number    uint8
}

func DEBUG(a ...interface{}) {
	if debug {
		fmt.Println(a)
	}
}

func DebugModeOn() {
	debug = true
}

func DebugModeOff() {
	debug = false
}

var debug bool = false

// MarshalBinary encodes the event into binary and returns the result
func (e Xbc_event) MarshalBinary() (data []byte, err error) {
	// 32 + 16 + 8 + 8 = 64bits = 8 bytes
	a := make([]byte, 8)

	a[0] = byte(e.Time >> 24)
	a[1] = byte((e.Time >> 16) & 0xff)
	a[2] = byte((e.Time >> 8) & 0xffff)
	a[3] = byte(e.Time & 0xffffff)
	a[4] = byte(e.Value >> 8)
	a[5] = byte(e.Value & 0xff)
	a[6] = byte(e.EventType)
	a[7] = byte(e.Number)
	return a, nil
}

// MarshalBinary encodes the event into binary and returns the result
func (e *Xbc_event) UnMarshalBinary(data []byte) (err error) {
	e.Time = uint32(data[0])<<24 + uint32(data[1])<<16 + uint32(data[2])<<8 + uint32(data[3])
	e.Value = int16(data[4])<<8 + int16(data[5])
	e.EventType = data[6]
	e.Number = data[7]
	return nil
}

// Xbc_state represents the state of a joystick
type Xbc_state struct {
	A           bool
	B           bool
	X           bool
	Y           bool
	Back        bool
	Start       bool
	LBumper     bool
	RBumper     bool
	RStickPress bool
	LStickPress bool
	Guide       bool
	LStickX     int16
	LStickY     int16
	RStickX     int16
	RStickY     int16
	LTrigger    int16
	RTrigger    int16
	DPadX       int16
	DPadY       int16
}

// Init returns a file descriptor to an open joystick for
// a given file
func Init(f string) int {
	return int(C.jsOpen(C.CString(f)))
}

// Close closes an open joystick file descriptor
func Close(fd int) int {
	return int(C.jsClose(C.int(fd)))
}

// GetXbEvent reads a single event from the passed file descriptor
func GetXbEvent(fd int) *Xbc_event {
	var e C.struct_js_event
	C.jsRead(C.int(fd), &e)

	var event Xbc_event
	event.Time = uint32(e.time)
	event.Value = int16(e.value)
	event.EventType = uint8(e._type)
	event.Number = uint8(e.number)
	return &event
}

// UpdateState updates a given state with the contents of a given
// event
func UpdateState(event *Xbc_event, state *Xbc_state) {
	if event.EventType == 1 {
		var btn_state bool
		if event.Value == 0 {
			btn_state = false
		} else {
			btn_state = true
		}

		switch event.Number {
		case 0:
			state.A = btn_state
			DEBUG("A: ", btn_state)
		case 1:
			state.B = btn_state
			DEBUG("B: ", btn_state)
		case 2:
			state.X = btn_state
			DEBUG("X: ", btn_state)
		case 3:
			state.Y = btn_state
			DEBUG("Y: ", btn_state)
		case 4:
			state.LBumper = btn_state
			DEBUG("LBumper: ", btn_state)
		case 5:
			state.RBumper = btn_state
			DEBUG("RBumper: ", btn_state)
		case 6:
			state.Back = btn_state
			DEBUG("Back: ", btn_state)
		case 7:
			state.Start = btn_state
			DEBUG("Start: ", btn_state)
		case 8:
			state.Guide = btn_state
			DEBUG("Guide: ", btn_state)
		case 9:
			state.LStickPress = btn_state
			DEBUG("L3: ", btn_state)
		case 10:
			state.RStickPress = btn_state
			DEBUG("R3: ", btn_state)
		}
	} else {

		switch event.Number {
		case 0:
			if event.Value > DEADZONE || event.Value < -1*DEADZONE {
				state.LStickX = event.Value
				DEBUG("LStickX: ", event.Value)
			}
		case 1:
			if event.Value > DEADZONE || event.Value < -1*DEADZONE {
				state.LStickY = event.Value
				DEBUG("LStickY: ", event.Value)
			}
		case 2:
			if event.Value > DEADZONE || event.Value < -1*DEADZONE {
				state.LTrigger = event.Value
				DEBUG("LTrigger: ", event.Value)
			}
		case 3:
			if event.Value > DEADZONE || event.Value < -1*DEADZONE {
				state.RStickY = event.Value
				DEBUG("RStickY: ", event.Value)
			}
		case 4:
			state.RStickX = event.Value
			DEBUG("RStickX: ", event.Value)
		case 5:
			state.RTrigger = event.Value
			DEBUG("RTrigger: ", event.Value)
		case 6:
			state.DPadX = event.Value
			DEBUG("DPadX: ", event.Value)
		case 7:
			state.DPadY = event.Value
			DEBUG("DPadY: ", event.Value)
		}
	}
}

func PrepState(state *Xbc_state) {
	state.A = false
	state.B = false
	state.X = false
	state.Y = false
	state.Back = false
	state.Start = false
	state.LBumper = false
	state.RBumper = false
	state.RStickPress = false
	state.LStickPress = false
	state.Guide = false
	state.LStickX = 0
	state.LStickY = 0
	state.RStickX = 0
	state.RStickY = 0
	state.LTrigger = -32768
	state.RTrigger = -32768
	state.DPadX = 0
	state.DPadY = 0
}

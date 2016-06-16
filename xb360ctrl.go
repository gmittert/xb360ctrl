// Package joystick provides functions for monitoring and updating
// the state of an Xbox 360 controller.
package xb360ctrl

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
  read(js, e, sizeof(struct js_event));
}
*/
import "C"

// Js_even is a Go analog to the C Joystick struct
type Xbc_event struct {
  Time int
  Value int
  EventType int
  Number int
}

// Xbc_state represents the state of a joystick
type Xbc_state struct {
  A bool
  B bool
  X bool
  Y bool
  Back bool
  Start bool
  LBumper bool
  RBumper bool
  RStickPress bool
  LStickPress bool
  Guide bool
  LStickX int
  LStickY int
  RStickX int
  RStickY int
  LTrigger int
  RTrigger int
  DPadX int
  DPadY int
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

// GetJsEvent reads a single event from the passed file descriptor
func GetJsEvent(fd int)* Xbc_event {
  var e C.struct_js_event
  C.jsRead(C.int(fd), &e)

  var event Xbc_event
  event.Time = int(e.time)
  event.Value = int(e.value)
  event.EventType = int(e._type)
  event.Number = int(e.number)
  return &event
}

// UpdateState updates a given state with the contents of a given
// event
func UpdateState(event *Xbc_event, state *Xbc_state) {
  if event.EventType == 1 {
    var btn_state bool
    if event.Number == 0 {
      btn_state = false
    } else {
      btn_state = true
    }

    switch event.Value {
    case 0:
      state.A = btn_state
    case 1:
      state.B = btn_state
    case 2:
      state.X = btn_state
    case 3:
      state.Y = btn_state
    case 4:
      state.LBumper = btn_state
    case 5:
      state.RBumper = btn_state
    case 6:
      state.Back = btn_state
    case 7:
      state.Start = btn_state
    case 8:
      state.Guide = btn_state
    case 9:
      state.LStickPress = btn_state
    case 10:
      state.RStickPress = btn_state
    }
  } else {
    switch event.Value {
    case 0:
      state.LStickX = event.Number
    case 1:
      state.LStickY = event.Number
    case 2:
      state.RTrigger = event.Number
    case 3:
      state.RStickY = event.Number
    case 4:
      state.RStickX = event.Number
    case 5:
      state.RTrigger = event.Number
    case 6:
      state.DPadX = event.Number
    case 7:
      state.DPadY = event.Number
    }
  }
}

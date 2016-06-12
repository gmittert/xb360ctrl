package joystick

/*
  #include <linux/joystick.h>
  #include <fcntl.h>
  #include <unistd.h>
  struct js_event e;
  size_t size = sizeof(&e);
  int rd = O_RDONLY;
  int jsOpen() {
    return open("/dev/input/js0", O_RDONLY);
  }
  void jsRead(int js) {
    read(js, &e, sizeof(e));
  }
*/
import "C"

type Js_event struct {
  Time int
  Value int
  EventType int
  Number int
}

var js C.int

func Init() {
  js = C.jsOpen()
}

func GetJsEvent()* Js_event {
  C.jsRead(js)
  var event Js_event
  event.Time = int(C.e.time)
  event.Value = int(C.e.value)
  event.EventType = int(C.e._type)
  event.Number = int(C.e.number)
  return &event
}

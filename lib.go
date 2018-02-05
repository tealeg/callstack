package framesabove

import (
	"fmt"
	"runtime"
)

func FramesAbove(match string, maxDepth int) *runtime.Frames {
	var rpc []uintptr

	rpc = make([]uintptr, maxDepth)
	count := runtime.Callers(1, rpc)
	if count < 1 {
		fmt.Printf("No callers found")
		return nil
	}
	frames := runtime.CallersFrames(rpc)
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if frame.Func == nil {
			return nil
		}
		if frame.Function == match {
			if more {
				return frames
			}
			return nil
		}
	}
	return nil

}

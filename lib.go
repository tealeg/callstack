package callstack

import (
	"runtime"
)

// FramesAbove returns a *runtime.Frames instance that, when Next() is
// called upon it, will return the first frame on the callstack above
// the provided function reference, and up to a maximum call depth.
func FramesAbove(match string, maxDepth int) *runtime.Frames {
	var rpc []uintptr

	rpc = make([]uintptr, maxDepth)
	count := runtime.Callers(1, rpc)
	if count < 1 {
		// No callers found
		return nil
	}
	frames := runtime.CallersFrames(rpc)
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if frame.Func == nil {
			// An empty frame
			return nil
		}
		if frame.Function == match {
			if more {
				// if there are more frames, return them
				return frames
			}
			// if there are no more frames there's nothing to return
			return nil
		}
	}
	// We've run out of frames, oops!
	return nil

}

package framesabove

import (
	"runtime"
	"testing"
)

func intermediateB(match string) *runtime.Frames {
	return FramesAbove(match, 10)
}

func intermediateA(match string) *runtime.Frames {
	return intermediateB(match)
}

func assertFrame(t *testing.T, match string, frame runtime.Frame) {
	if frame.Func == nil {
		t.Fatal("Nil frame.Func returned")
	}
	if frame.Function != match {
		t.Fatalf("Expected %q, but got %q", match, frame.Function)
	}
}

func nextFrame(t *testing.T, frames *runtime.Frames) runtime.Frame {
	result, more := frames.Next()
	if !more {
		t.Fatal("Unexpected end of Frames")
	}
	return result
}

func TestFramesAbove(t *testing.T) {
	frames := intermediateA("github.com/tealeg/framesabove.FramesAbove")
	frame := nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.intermediateB", frame)
	frame = nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.intermediateA", frame)
	frame = nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.TestFramesAbove", frame)

	frames = intermediateA("github.com/tealeg/framesabove.intermediateB")
	frame = nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.intermediateA", frame)
	frame = nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.TestFramesAbove", frame)

	frames = intermediateA("github.com/tealeg/framesabove.intermediateA")
	frame = nextFrame(t, frames)
	assertFrame(t, "github.com/tealeg/framesabove.TestFramesAbove", frame)
}

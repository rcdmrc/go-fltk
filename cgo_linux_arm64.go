//go:build linux && arm64

package fltk

// #cgo linux,arm64 CXXFLAGS: -std=c++11
// #cgo linux,arm64 CPPFLAGS: -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo linux,arm64 LDFLAGS: -lfltk_images -lfltk_jpeg -lfltk_png -lfltk_z -lfltk_gl -lGLU -lGL -lfltk_forms -lfltk -lm -lX11 -lpthread -lm -ldl
import "C"

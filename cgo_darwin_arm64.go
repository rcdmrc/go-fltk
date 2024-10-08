//go:build darwin && arm64

package fltk

// #cgo darwin,arm64 CXXFLAGS: -std=c++11
// #cgo darwin,arm64 CPPFLAGS: -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo darwin,arm64 LDFLAGS: -lfltk_images -lfltk_jpeg -lfltk_png -lfltk_z -lfltk_gl -lfltk_forms -lfltk -lm -lpthread -framework OpenGL -framework Cocoa
import "C"

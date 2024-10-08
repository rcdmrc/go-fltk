//go:build linux && arm

package fltk

// #cgo linux,arm CXXFLAGS: -std=c++11
// #cgo linux,arm CPPFLAGS: -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo linux,arm LDFLAGS: -lfltk_images -lfltk_jpeg -lfltk_png -lfltk_z -lfltk_gl -lGLU -lGL -lfltk_forms -lfltk -lm -lX11 -lXext -lpthread -lXinerama -lXfixes -lXcursor -lXft -lXrender -lm -lfontconfig -ldl
import "C"

// go:build windows && amd64

package fltk

// #cgo windows,amd64 CXXFLAGS: -std=c++11
// #cgo windows,amd64 CPPFLAGS: -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64
// #cgo windows,amd64 LDFLAGS: -mwindows -lfltk_images -lfltk_jpeg -lfltk_png -lfltk_z -lfltk_gl -lglu32 -lopengl32 -lfltk_forms -lfltk -lgdiplus -lole32 -luuid -lcomctl32 -lws2_32
import "C"

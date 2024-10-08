//go:build openbsd && arm64

package fltk

// #cgo openbsd,arm64 CXXFLAGS: -std=c++11
// #cgo openbsd,arm64 CPPFLAGS: -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo openbsd,arm64 LDFLAGS: -L/usr/X11R6/lib -lfltk_images -lfltk_jpeg -lfltk_png -lfltk_z -lfltk_gl -lGLU -lGL -lfltk_forms -lfltk -lm -lX11 -lXext -lpthread -lXinerama -lXfixes -lXcursor -lXft -lXrender -lm -lfontconfig
import "C"

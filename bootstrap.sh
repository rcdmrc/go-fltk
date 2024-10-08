#!/usr/bin/env bash
set -e

fltk_cxxflags=()
fltk_cppflags=()
fltk_ldflags=()

ldflags=(\
    -lfltk_images\
    -lfltk_jpeg\
    -lfltk_png\
    -lfltk_z\
    -lfltk_gl\
    -lfltk_forms\
    -lfltk\
    -lm\
    -lpthread\
)

cxxflags=(\
    -std=c++11\
)

cppflags=(\
    -D_LARGEFILE_SOURCE\
    -D_LARGEFILE64_SOURCE\
    -D_FILE_OFFSET_BITS=64\
    -D_THREAD_SAFE\
    -D_REENTRANT\
)

darwin_ldflags=(\
    -framework\
    OpenGL\
    -framework\
    UniformTypeIdentifiers\
    -framework\
    ScreenCaptureKit\
    -framework\
    Cocoa\
)

###

case $1 in
  -h)
    echo "Create CGO flags for the host architecture."
    echo ""
    echo "Usage: "
    echo ""
    echo "  ./bootstrap.sh"
    echo ""
    echo "Environment variables:"
    echo ""
    echo "  FLTK_ROOT: fltk install path"
    echo ""
    exit 1
    ;;
  *)
    ;;
esac

if [ ! -z "$FLTK_ROOT" ]; then
    ldflags=("-L${FLTK_ROOT}/lib ${ldflags[@]}")
    cppflags=("-I${FLTK_ROOT}/include ${cppflags[@]}")
fi

###

os_name=""
os_arch=""

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    os_name="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    os_name="darwin"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
    os_name="unix"
else
    >&2 echo "unable to determine OS :("
    exit 1
fi
os_arch=$(uname -m)
out_file="cgo_${os_name}_${os_arch}.go"

echo "os_name : ${os_name}"
echo "os_arch : ${os_arch}"
echo "out_file: ${out_file}"

if [[ "${os_name}" = "darwin" ]]; then
    ldflags+=(${darwin_ldflags[@]})
fi

cat <<EOF > ${out_file}
//go:build ${os_name} && ${os_arch}

package fltk

// #cgo ${os_name},${os_arch} CXXFLAGS: ${cxxflags[@]}
// #cgo ${os_name},${os_arch} CPPFLAGS: ${cppflags[@]}
// #cgo ${os_name},${os_arch} LDFLAGS:  ${ldflags[@]}
import "C"
EOF
### LLVM Passes

This directory contains a set of simple passes of LLVM.

#### How to compile
```bash
# Get the LLVM and Clang source
svn co http://llvm.org/svn/llvm-project/llvm/trunk llvm-src
cd llvm/tools
svn co http://llvm.org/svn/llvm-project/cfe/trunk clang
# Get source code of this repo and copy them into your llvm source directory
# llvm-src/lib/Transform/HelloAgain
# llvm-src/lib/Transform/CMakeLists.txt
cd ../..
mkdir llvm-build
cd llvm-build
cmake ../llvm-src
make -j4
```

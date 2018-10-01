//===- SimplePasses.cpp - Example code from "Writing an LLVM Pass" --------===//
//
//                     The LLVM Compiler Infrastructure
//
// This file is distributed under the University of Illinois Open Source
// License. See LICENSE.TXT for details.
//
//===----------------------------------------------------------------------===//
//
// This file implements two versions of the LLVM "Hello World" pass described
// in docs/WritingAnLLVMPass.html
//
//===----------------------------------------------------------------------===//

#include "llvm/ADT/Statistic.h"
#include "llvm/IR/Module.h"
#include "llvm/IR/Function.h"
#include "llvm/Pass.h"
#include "llvm/Analysis/CallGraphSCCPass.h"
#include "llvm/Analysis/CallGraph.h"
#include "llvm/Support/raw_ostream.h"
using namespace llvm;

namespace {
  // ModulePass
  // -- Traverse a program and iterate over its Modules, Functions, BasicBlocks.
  struct SimpleModulePass : public ModulePass {
    static char ID;
    SimpleModulePass() : ModulePass(ID) {  }
  
    bool runOnModule(Module &M) override {
      errs() << "Enter Module: ";
      errs().write_escaped(M.getName()) << '\n';
      for (auto &F: M) {
        errs() << "Enter Function: ";
        errs().write_escaped(F.getName()) << '\n';
        for (auto &BB: F) {
          errs() << "Enter BasicBlock: ";
          errs().write_escaped(BB.getName()) << '\n';
          /* This will get nothing, because BasicBlock has no name by default */
          for (auto &I: BB) {
            errs() << "Instruction: ";
            errs() << I.getOpcodeName() << '\n';
          }
        }
      }
      return false; /* We only collect some information */
    }
  };

  // CallGraphSCCPass
  // -- Simple CallGraphSCCPass that dumps call graph info
  struct SimpleCallGraphSCCPass: public CallGraphSCCPass {
    static char ID;
    SimpleCallGraphSCCPass(): CallGraphSCCPass(ID) {  }
  
    bool runOnSCC(CallGraphSCC &SCC) override {
      errs() << " --- Enter Call Graph SCC ---\n";
      for (auto &G : SCC) {
        G->dump();
      }
      return false;
      errs() << " --- end of CallGraphSCC ---\n";
    }
  };

  struct OpsCounter : public FunctionPass {
    static char ID;
    std::map<std::string, int> opCounter;
    OpsCounter(): FunctionPass(ID) {  }
  
    bool runOnFunction(Function &F) override {
      for (auto &BB : F) {
        for (auto &I : BB) {
          auto opcode_it = opCounter.find(I.getOpcodeName());
          if (opcode_it != opCounter.end())
            // find one
            ++ opCounter[I.getOpcodeName()];
          else
            opCounter[I.getOpcodeName()] = 1;
        }
      }
      errs() << "opcode" << '\t' << "count" << '\n';
      for (auto &op : opCounter)
        errs() << op.first << '\t' << op.second << '\n';
      return false;
    }
  };
}

char SimpleModulePass::ID = 0;
static RegisterPass<SimpleModulePass> X("smp", "SimpleModulePass");

char SimpleCallGraphSCCPass::ID = 0;
static RegisterPass<SimpleCallGraphSCCPass> Y("scgsccp", "SimpleCallGraphSCCPass");

char OpsCounter::ID = 0;
static RegisterPass<OpsCounter> Z("opscnter", "OpsCounter");

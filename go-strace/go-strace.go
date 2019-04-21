package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"mysyscallCounter"
)

func main() {
	var ss mysyscallCounter.SyscallCounter
	
	fmt.Printf("Run %v\n", os.Args[1:])
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	var regs syscall.PtraceRegs
	exit := true
	
	for {
	if exit {
		err = syscall.PtraceGetRegs(pid, &regs)
		if err != nil {
		break
		}
		fmt.Printf("%s\n", ss.GetName(regs.Orig_rax))
   }
	err = syscall.PtraceSyscall(pid,0)
	  if err != nil {
		  panic(err)
	  }
	_,err = syscall.Wait4(pid,nil,0,nil)
	  if err != nil {
		  panic(err)
	  }
	  exit = !exit
   }
}

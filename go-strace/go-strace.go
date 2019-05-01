package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"Go-Practice/go-strace/mysyscallcounter"
)

func main() {
	var ss mysyscallcounter.SyscallCounter
	
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
	
	for{
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}

			 name := ss.GetName(regs.Orig_rax)
			 fmt.Printf("%s\n", name)
			
		

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}
	}
		ss.Print()
}
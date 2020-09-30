package runner

import (
	"os"
	"os/exec"
	"syscall"
)

func New(command []string) *Runner {
	return &Runner{
		chunks:    command,
		reloading: make(chan struct{}),

		terminationInit: make(chan struct{}),
		terminationDone: make(chan struct{}),
	}
}

type Runner struct {
	chunks    []string
	reloading chan struct{}

	terminationInit chan struct{}
	terminationDone chan struct{}
}

func (r *Runner) Run() {
	var cmd *exec.Cmd

	go func() {
		r.reloading <- struct{}{} // start loop workaround
	}()

	for {

		select {

		case <-r.terminationInit:
			if cmd != nil {
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			}
			r.terminationDone <- struct{}{}
			return

		case <-r.reloading:
			if cmd != nil {
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			}

			cmd = exec.Command(r.chunks[0], r.chunks[1:]...)
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			go func() {
				_ = cmd.Run()
			}()
		}

	}

}

func (r *Runner) Reload() {
	r.reloading <- struct{}{}
}

func (r *Runner) Terminate() {
	r.terminationInit <- struct{}{}
	<-r.terminationDone
}

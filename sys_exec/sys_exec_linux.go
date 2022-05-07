//go:build linux
// +build linux

package sys_exec

func BuildCmd(cmdStr string) *exec.Cmd {
	return exec.Command("sh", "-c", cmdStr)
}

func SetPgid(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	} else {
		cmd.SysProcAttr.Setpgid = true
	}
}

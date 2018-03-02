# Gitlab Runner

- repo: https://github.com/gitlabhq/gitlab-runner

## shell 

- set process group, and kill process group
  - process group and pipe https://www.usna.edu/Users/cs/aviv/classes/ic221/s16/lec/17/lec.html
  
helpers/process_group_unix.go

````go
func SetProcessGroup(cmd *exec.Cmd) {
	// Create process group
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}

// man 2 kill
func KillProcessGroup(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}

	process := cmd.Process
	if process != nil {
		if process.Pid > 0 {
			syscall.Kill(-process.Pid, syscall.SIGKILL)
		} else {
			// doing normal kill
			process.Kill()
		}
	}
}
````

## ssh

heplers/ssh/ssh_command.go

- pass environment variable over ssh by using multi reader for stdin

````go
  var envVariables bytes.Buffer
	for _, keyValue := range cmd.Environment {
		envVariables.WriteString("export " + helpers.ShellEscape(keyValue) + "\n")
	}

	session.Stdin = io.MultiReader(
		&envVariables,
		bytes.NewBufferString(cmd.Stdin),
	)
	session.Stdout = s.Stdout
	session.Stderr = s.Stderr
	err = session.Start(cmd.fullCommand())
	if err != nil {
		return err
	}
````
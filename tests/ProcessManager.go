package tests

import (
	"os/exec"
	"runtime"
	"strconv"
	"sync"
)

// ProcessManager 进程管理器
type ProcessManager struct {
	processes []*exec.Cmd
	mutex     sync.Mutex
}

// NewProcessManager 创建新的进程管理器
func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make([]*exec.Cmd, 0),
	}
}

// AddProcess 添加进程到管理器
func (pm *ProcessManager) AddProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.processes = append(pm.processes, cmd)
}

// CleanupAll 清理所有进程
func (pm *ProcessManager) CleanupAll() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			// Windows系统下使用更强制的方式终止进程
			if runtime.GOOS == "windows" {
				// 在Windows上，我们需要终止整个进程树
				cmd.Process.Kill()
				// 等待进程退出
				cmd.Wait()

				// 尝试查找并终止子进程
				pm.killChildProcesses(cmd.Process.Pid)
			} else {
				// Unix系统下使用进程组
				cmd.Process.Kill()
				cmd.Wait()
			}
		}
	}
	pm.processes = make([]*exec.Cmd, 0)
}

// killChildProcesses 在Windows上终止子进程
func (pm *ProcessManager) killChildProcesses(parentPid int) {
	// 在Windows上使用taskkill命令终止进程树
	killCmd := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(parentPid))
	killCmd.Run() // 忽略错误，因为进程可能已经退出
}

// GetPIDs 获取所有进程的PID
func (pm *ProcessManager) GetPIDs() []string {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	var pids []string
	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			pids = append(pids, strconv.Itoa(cmd.Process.Pid))
		}
	}
	return pids
}

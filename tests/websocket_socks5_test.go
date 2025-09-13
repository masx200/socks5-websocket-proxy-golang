package tests

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

// runWebSocketSocks5Proxy 测试WebSocket和SOCKS5级联代理功能
func runWebSocketSocks5Proxy(t *testing.T) {
	// 创建进程管理器
	processManager := NewProcessManager()
	defer processManager.CleanupAll()
	defer processManager.Close() // 确保日志文件被正确关闭

	// 创建缓冲区来捕获WebSocket服务器的输出
	var websocketOutput bytes.Buffer
	var websocketOutputMutex sync.Mutex

	// 创建缓冲区来捕获SOCKS5服务器的输出
	var socks5Output bytes.Buffer
	var socks5OutputMutex sync.Mutex

	// 创建一个多写入器，同时写入到标准输出和缓冲区
	websocketWriter := io.MultiWriter(os.Stdout, &websocketOutput)
	socks5Writer := io.MultiWriter(os.Stdout, &socks5Output)

	// 创建带缓冲的写入器来确保日志被及时刷新
	websocketBufWriter := bufio.NewWriter(websocketWriter)
	socks5BufWriter := bufio.NewWriter(socks5Writer)

	// 清理可能存在的旧的可执行文件
	if _, err := os.Stat("main.exe"); err == nil {
		os.Remove("main.exe")
	}

	// 添加测试超时检查
	timeoutTimer := time.AfterFunc(35*time.Second, func() {
		log.Println("\n⚠️ 测试即将超时，正在清理进程...")
		// 在超时前记录服务器日志
		var timeoutTestResults []string

		// 使用互斥锁保护对输出的访问
		websocketOutputMutex.Lock()
		websocketOutputLen := websocketOutput.Len()
		websocketOutputContent := websocketOutput.String()
		websocketOutputMutex.Unlock()

		socks5OutputMutex.Lock()
		socks5OutputLen := socks5Output.Len()
		socks5OutputContent := socks5Output.String()
		socks5OutputMutex.Unlock()

		timeoutTestResults = []string{
			"# WebSocket和SOCKS5级联测试记录（超时）",
			"",
			"## 测试时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
		}

		// 添加WebSocket服务器日志
		if websocketOutputLen > 0 {
			timeoutTestResults = append(timeoutTestResults, "## WebSocket服务器日志输出（超时前捕获）")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(websocketOutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					timeoutTestResults = append(timeoutTestResults, line)
				}
			}
			timeoutTestResults = append(timeoutTestResults, "```")
			timeoutTestResults = append(timeoutTestResults, "")
		} else {
			timeoutTestResults = append(timeoutTestResults, "## WebSocket服务器状态")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "⚠️ 没有捕获到WebSocket服务器日志")
			timeoutTestResults = append(timeoutTestResults, "")
		}

		// 添加SOCKS5服务器日志
		if socks5OutputLen > 0 {
			timeoutTestResults = append(timeoutTestResults, "## SOCKS5服务器日志输出（超时前捕获）")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(socks5OutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					timeoutTestResults = append(timeoutTestResults, line)
				}
			}
			timeoutTestResults = append(timeoutTestResults, "```")
			timeoutTestResults = append(timeoutTestResults, "")
		} else {
			timeoutTestResults = append(timeoutTestResults, "## SOCKS5服务器状态")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "⚠️ 没有捕获到SOCKS5服务器日志")
			timeoutTestResults = append(timeoutTestResults, "")
		}

		// 添加调试信息
		timeoutTestResults = append(timeoutTestResults, "## 调试信息")
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, fmt.Sprintf("[DEBUG] WebSocket输出长度: %d", websocketOutputLen))
		timeoutTestResults = append(timeoutTestResults, fmt.Sprintf("[DEBUG] SOCKS5输出长度: %d", socks5OutputLen))
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, "❌ 测试超时，但已捕获服务器日志")

		// 写入超时测试记录
		if err := writeTestResults3(timeoutTestResults); err != nil {
			log.Printf("写入超时测试记录失败: %v\n", err)
		}
		processManager.CleanupAll()
		processManager.Close()
		// 强制退出测试
		t.Fatal("测试超时")
	})
	defer timeoutTimer.Stop()

	// 测试结果记录
	var testResults []string
	testResults = append(testResults, "# WebSocket和SOCKS5级联测试记录")
	testResults = append(testResults, "")
	testResults = append(testResults, "## 测试时间")
	testResults = append(testResults, time.Now().Format("2006-01-02 15:04:05"))
	testResults = append(testResults, "")

	// 检查端口是否被占用
	if isPortOccupied3(8080) {
		t.Fatal("端口8080已被占用，请先停止占用该端口的进程")
	}
	if isPortOccupied3(10810) {
		t.Fatal("端口10810已被占用，请先停止占用该端口的进程")
	}

	// 编译代理服务器
	testResults = append(testResults, "## 1. 编译代理服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `go build -o main.exe ../cmd/main.go`")
	testResults = append(testResults, "")

	// 先编译代理服务器
	testResults = append(testResults, "编译代理服务器...")
	buildCmd := exec.Command("go", "build", "-o", "main.exe", "../cmd/main.go")
	buildCmd.Stdout = websocketBufWriter
	buildCmd.Stderr = websocketBufWriter

	// 记录编译命令
	processManager.LogCommand(buildCmd, "BUILD")

	if err := buildCmd.Run(); err != nil {
		// 刷新缓冲区确保所有日志都被捕获
		websocketBufWriter.Flush()
		processManager.LogCommandResult(buildCmd, err, "")
		t.Fatalf("编译代理服务器失败: %v", err)
	}
	// 刷新缓冲区确保所有日志都被捕获
	websocketBufWriter.Flush()
	processManager.LogCommandResult(buildCmd, nil, "")
	testResults = append(testResults, "✅ 代理服务器编译成功")
	testResults = append(testResults, "")

	// 启动WebSocket服务器
	testResults = append(testResults, "## 2. 启动WebSocket服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `./main.exe -mode server -protocol websocket -addr :8080`")
	testResults = append(testResults, "")

	websocketCmd := exec.Command("./main.exe", "-mode", "server", "-protocol", "websocket", "-addr", ":8080")
	websocketCmd.Stdout = websocketBufWriter
	websocketCmd.Stderr = websocketBufWriter

	// 设置进程属性，确保能终止所有子进程（跨平台兼容）
	if runtime.GOOS == "windows" {
		// Windows特定的进程组设置
		websocketCmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		}
	}
	// Unix-like系统不需要特殊设置，go会自动处理

	err := websocketCmd.Start()
	if err != nil {
		// 刷新缓冲区确保错误日志被捕获
		websocketBufWriter.Flush()
		t.Fatalf("启动WebSocket服务器失败: %v", err)
	}

	// 将WebSocket服务器进程添加到管理器
	processManager.AddProcess(websocketCmd)
	log.Printf("WebSocket服务器已启动，PID: %d\n", websocketCmd.Process.Pid)

	// 等待服务器启动并刷新缓冲区
	time.Sleep(2 * time.Second)
	websocketBufWriter.Flush()
	log.Printf("WebSocket服务器启动检查完成，当前输出长度: %d", websocketOutput.Len())

	// 记录启动命令
	processManager.LogCommand(websocketCmd, "SERVER")

	// 确保进程能正确退出
	go func() {
		websocketCmd.Wait()
		log.Println("WebSocket服务器进程已退出")
	}()

	// 记录WebSocket服务器PID
	testResults = append(testResults, fmt.Sprintf("📋 WebSocket服务器进程PID: %d", websocketCmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待服务器启动
	testResults = append(testResults, "等待WebSocket服务器启动...")
	time.Sleep(2 * time.Second)

	// 启动SOCKS5服务器
	testResults = append(testResults, "## 3. 启动SOCKS5服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`")
	testResults = append(testResults, "")

	socks5Cmd := exec.Command("./main.exe", "-mode", "server", "-protocol", "socks5", "-addr", ":10810",
		"-upstream-type", "websocket", "-upstream-address", "ws://localhost:8080")
	socks5Cmd.Stdout = socks5BufWriter
	socks5Cmd.Stderr = socks5BufWriter

	// 设置进程属性，确保能终止所有子进程（跨平台兼容）
	if runtime.GOOS == "windows" {
		// Windows特定的进程组设置
		socks5Cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		}
	}
	// Unix-like系统不需要特殊设置，go会自动处理

	err = socks5Cmd.Start()
	if err != nil {
		// 刷新缓冲区确保错误日志被捕获
		socks5BufWriter.Flush()
		t.Fatalf("启动SOCKS5服务器失败: %v", err)
	}

	// 将SOCKS5服务器进程添加到管理器
	processManager.AddProcess(socks5Cmd)
	log.Printf("SOCKS5服务器已启动，PID: %d\n", socks5Cmd.Process.Pid)

	// 等待服务器启动并刷新缓冲区
	time.Sleep(2 * time.Second)
	socks5BufWriter.Flush()
	log.Printf("SOCKS5服务器启动检查完成，当前输出长度: %d", socks5Output.Len())

	// 记录启动命令
	processManager.LogCommand(socks5Cmd, "SERVER")

	// 确保进程能正确退出
	go func() {
		socks5Cmd.Wait()
		log.Println("SOCKS5服务器进程已退出")
	}()

	// 记录SOCKS5服务器PID
	testResults = append(testResults, fmt.Sprintf("📋 SOCKS5服务器进程PID: %d", socks5Cmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待服务器启动
	testResults = append(testResults, "等待SOCKS5服务器启动...")
	time.Sleep(2 * time.Second)

	// 添加启动成功的日志输出提示
	log.Println("WebSocket和SOCKS5服务器启动成功，开始执行测试...")

	// 测试SOCKS5代理功能
	testResults = append(testResults, "## 4. 测试SOCKS5代理功能")
	testResults = append(testResults, "")

	// 第一个curl测试
	testResults = append(testResults, "### 测试1: HTTP代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`")
	testResults = append(testResults, "")

	// 创建curl进程
	curlCmd1 := exec.Command("curl", "-v", "-I", "http://www.baidu.com", "-x", "socks5://localhost:10810")
	// 创建缓冲区来捕获curl输出
	var curlOutput1 bytes.Buffer
	curlCmd1.Stdout = &curlOutput1
	curlCmd1.Stderr = &curlOutput1

	// 记录测试命令
	processManager.LogCommand(curlCmd1, "TEST")

	// 启动curl进程
	err1 := curlCmd1.Run()
	output1 := curlOutput1.Bytes()

	// 检查进程退出状态码
	exitCode1 := 0
	if curlCmd1.ProcessState != nil {
		exitCode1 = curlCmd1.ProcessState.ExitCode()
	}

	// 将curl进程添加到管理器
	processManager.AddProcess(curlCmd1)

	// 记录curl进程PID和退出状态码
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试1进程PID: %d, 退出状态码: %d", curlCmd1.Process.Pid, exitCode1))
	testResults = append(testResults, "")
	if err1 != nil || exitCode1 != 0 {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err1))
		testResults = append(testResults, fmt.Sprintf("退出状态码: %d", exitCode1))
		testResults = append(testResults, fmt.Sprintf("错误输出: %s", string(output1)))
	} else {
		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, "输出结果:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(output1))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 记录测试结果
	processManager.LogCommandResult(curlCmd1, err1, string(output1))

	// 第二个curl测试
	testResults = append(testResults, "### 测试2: HTTPS代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`")
	testResults = append(testResults, "")

	// 创建curl进程
	curlCmd2 := exec.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "socks5://localhost:10810")
	// 创建缓冲区来捕获curl输出
	var curlOutput2 bytes.Buffer
	curlCmd2.Stdout = &curlOutput2
	curlCmd2.Stderr = &curlOutput2

	// 记录测试命令
	processManager.LogCommand(curlCmd2, "TEST")

	// 启动curl进程
	err2 := curlCmd2.Run()
	output2 := curlOutput2.Bytes()

	// 检查进程退出状态码
	exitCode2 := 0
	if curlCmd2.ProcessState != nil {
		exitCode2 = curlCmd2.ProcessState.ExitCode()
	}

	// 将curl进程添加到管理器
	processManager.AddProcess(curlCmd2)

	// 记录curl进程PID和退出状态码
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试2进程PID: %d, 退出状态码: %d", curlCmd2.Process.Pid, exitCode2))
	testResults = append(testResults, "")
	if err2 != nil || exitCode2 != 0 {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err2))
		testResults = append(testResults, fmt.Sprintf("退出状态码: %d", exitCode2))
		testResults = append(testResults, fmt.Sprintf("错误输出: %s", string(output2)))
	} else {
		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, "输出结果:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(output2))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 记录测试结果
	processManager.LogCommandResult(curlCmd2, err2, string(output2))

	// 记录所有进程PID信息
	testResults = append(testResults, "### 📋 所有进程PID记录")
	testResults = append(testResults, "")
	allPIDs := processManager.GetPIDs()
	testResults = append(testResults, fmt.Sprintf("所有进程PID: %s", strings.Join(allPIDs, ", ")))
	testResults = append(testResults, "")

	// 写入测试记录到文件
	err = writeTestResults3(testResults)
	if err != nil {
		t.Errorf("写入测试记录失败: %v", err)
	}

	// 验证测试结果
	if err1 != nil {
		t.Errorf("第一个curl测试失败: %v", err1)
	}
	if err2 != nil {
		t.Errorf("第二个curl测试失败: %v", err2)
	}

	// 如果curl命令运行成功，关闭服务器进程
	if err1 == nil && err2 == nil {
		testResults = append(testResults, "## 5. 关闭服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "✅ 所有curl测试成功，正在关闭服务器进程...")
		testResults = append(testResults, "")

		// 停止超时计时器
		timeoutTimer.Stop()

		// 终止SOCKS5服务器
		testResults = append(testResults, "🛑 正在终止SOCKS5服务器进程...")
		if socks5Cmd.Process != nil {
			log.Printf("正在终止SOCKS5服务器进程 PID: %d\n", socks5Cmd.Process.Pid)

			if err := socks5Cmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止SOCKS5服务器进程失败: %v", err))
				log.Printf("终止SOCKS5服务器进程失败: %v\n", err)
			} else {
				socks5Cmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ SOCKS5服务器进程已终止")
				log.Println("SOCKS5服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 终止WebSocket服务器
		testResults = append(testResults, "🛑 正在终止WebSocket服务器进程...")
		if websocketCmd.Process != nil {
			log.Printf("正在终止WebSocket服务器进程 PID: %d\n", websocketCmd.Process.Pid)

			if err := websocketCmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止WebSocket服务器进程失败: %v", err))
				log.Printf("终止WebSocket服务器进程失败: %v\n", err)
			} else {
				websocketCmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ WebSocket服务器进程已终止")
				log.Println("WebSocket服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 清理所有进程
		testResults = append(testResults, "🧹 正在清理所有子进程...")
		testResults = append(testResults, "")
		processManager.CleanupAll()
		testResults = append(testResults, "✅ 所有子进程已清理完成")
		testResults = append(testResults, "")

		// 等待进程完全关闭并释放资源
		time.Sleep(2 * time.Second)

		// 等待进程完全退出
		time.Sleep(2 * time.Second)

		// 清理编译的可执行文件
		if _, err := os.Stat("main.exe"); err == nil {
			os.Remove("main.exe")
			testResults = append(testResults, "🧹 已清理编译的可执行文件")
		}

		// 将WebSocket服务器输出添加到测试记录
		log.Println("正在记录WebSocket服务器日志...")

		// 刷新缓冲区确保所有日志都被捕获
		websocketBufWriter.Flush()

		// 使用互斥锁保护对websocketOutput的访问
		websocketOutputMutex.Lock()
		websocketOutputLen := websocketOutput.Len()
		websocketOutputContent := websocketOutput.String()
		websocketOutputMutex.Unlock()

		if websocketOutputLen > 0 {
			testResults = append(testResults, "### WebSocket服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(websocketOutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
					log.Println("[WebSocket日志]", line) // 同时打印到控制台
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		} else {
			testResults = append(testResults, "### WebSocket服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "⚠️ 没有捕获到WebSocket服务器日志")
			testResults = append(testResults, "")
			log.Println("⚠️ 没有捕获到WebSocket服务器日志")

			// 添加调试信息
			testResults = append(testResults, "### 调试信息")
			testResults = append(testResults, "")
			testResults = append(testResults, fmt.Sprintf("WebSocket服务器输出缓冲区长度: %d", websocketOutputLen))
			testResults = append(testResults, "")
			testResults = append(testResults, "可能的原因:")
			testResults = append(testResults, "- WebSocket服务器程序没有输出日志")
			testResults = append(testResults, "- 日志输出被重定向到其他地方")
			testResults = append(testResults, "- 缓冲区没有正确捕获输出")
			testResults = append(testResults, "")
		}

		// 将SOCKS5服务器输出添加到测试记录
		log.Println("正在记录SOCKS5服务器日志...")

		// 刷新缓冲区确保所有日志都被捕获
		socks5BufWriter.Flush()

		// 使用互斥锁保护对socks5Output的访问
		socks5OutputMutex.Lock()
		socks5OutputLen := socks5Output.Len()
		socks5OutputContent := socks5Output.String()
		socks5OutputMutex.Unlock()

		if socks5OutputLen > 0 {
			testResults = append(testResults, "### SOCKS5服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(socks5OutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
					log.Println("[SOCKS5日志]", line) // 同时打印到控制台
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		} else {
			testResults = append(testResults, "### SOCKS5服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "⚠️ 没有捕获到SOCKS5服务器日志")
			testResults = append(testResults, "")
			log.Println("⚠️ 没有捕获到SOCKS5服务器日志")

			// 添加调试信息
			testResults = append(testResults, "### 调试信息")
			testResults = append(testResults, "")
			testResults = append(testResults, fmt.Sprintf("SOCKS5服务器输出缓冲区长度: %d", socks5OutputLen))
			testResults = append(testResults, "")
			testResults = append(testResults, "可能的原因:")
			testResults = append(testResults, "- SOCKS5服务器程序没有输出日志")
			testResults = append(testResults, "- 日志输出被重定向到其他地方")
			testResults = append(testResults, "- 缓冲区没有正确捕获输出")
			testResults = append(testResults, "")
		}

		// 将curl进程输出添加到测试记录
		testResults = append(testResults, "### 所有子进程日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "```")

		// 添加curl1输出
		if curlOutput1.Len() > 0 {
			testResults = append(testResults, "### Curl测试1输出 ###")
			curl1Lines := strings.Split(curlOutput1.String(), "\n")
			for _, line := range curl1Lines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
		}

		// 添加curl2输出
		if curlOutput2.Len() > 0 {
			testResults = append(testResults, "### Curl测试2输出 ###")
			curl2Lines := strings.Split(curlOutput2.String(), "\n")
			for _, line := range curl2Lines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
		}

		testResults = append(testResults, "```")
		testResults = append(testResults, "")

		// 验证端口是否已释放
		if !isPortOccupied3(8080) && !isPortOccupied3(10810) {
			testResults = append(testResults, "✅ 端口8080和10810已成功释放")
		} else {
			testResults = append(testResults, "❌ 端口8080或10810仍被占用")
		}

		// 重新写入测试记录
		err = writeTestResults3(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}

	} else {
		// 如果有测试失败，也记录关闭进程的信息
		testResults = append(testResults, "## 5. 关闭服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "⚠️ 部分测试失败，但仍需关闭服务器进程...")
		testResults = append(testResults, "")

		// 终止SOCKS5服务器
		testResults = append(testResults, "🛑 正在终止SOCKS5服务器进程...")
		if socks5Cmd.Process != nil {
			log.Printf("正在终止SOCKS5服务器进程 PID: %d\n", socks5Cmd.Process.Pid)

			if err := socks5Cmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止SOCKS5服务器进程失败: %v", err))
				log.Printf("终止SOCKS5服务器进程失败: %v\n", err)
			} else {
				socks5Cmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ SOCKS5服务器进程已终止")
				log.Println("SOCKS5服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 终止WebSocket服务器
		testResults = append(testResults, "🛑 正在终止WebSocket服务器进程...")
		if websocketCmd.Process != nil {
			log.Printf("正在终止WebSocket服务器进程 PID: %d\n", websocketCmd.Process.Pid)

			if err := websocketCmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止WebSocket服务器进程失败: %v", err))
				log.Printf("终止WebSocket服务器进程失败: %v\n", err)
			} else {
				websocketCmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ WebSocket服务器进程已终止")
				log.Println("WebSocket服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 清理所有进程
		testResults = append(testResults, "🧹 正在清理所有子进程...")
		testResults = append(testResults, "")
		processManager.CleanupAll()
		testResults = append(testResults, "✅ 所有子进程已清理完成")
		testResults = append(testResults, "")

		// 等待进程完全关闭并释放资源
		time.Sleep(2 * time.Second)

		// 等待进程完全退出
		time.Sleep(2 * time.Second)

		// 清理编译的可执行文件
		if _, err := os.Stat("main.exe"); err == nil {
			os.Remove("main.exe")
			testResults = append(testResults, "🧹 已清理编译的可执行文件")
		}

		// 将WebSocket服务器输出添加到测试记录

		// 使用互斥锁保护对websocketOutput的访问
		websocketOutputMutex.Lock()
		websocketOutputLen := websocketOutput.Len()
		websocketOutputContent := websocketOutput.String()
		websocketOutputMutex.Unlock()

		if websocketOutputLen > 0 {
			testResults = append(testResults, "### WebSocket服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(websocketOutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		} else {
			testResults = append(testResults, "### WebSocket服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "⚠️ 没有捕获到WebSocket服务器日志")
			testResults = append(testResults, "")

			// 添加调试信息
			testResults = append(testResults, "### 调试信息")
			testResults = append(testResults, "")
			testResults = append(testResults, fmt.Sprintf("WebSocket服务器输出缓冲区长度: %d", websocketOutputLen))
			testResults = append(testResults, "")
			testResults = append(testResults, "可能的原因:")
			testResults = append(testResults, "- WebSocket服务器程序没有输出日志")
			testResults = append(testResults, "- 日志输出被重定向到其他地方")
			testResults = append(testResults, "- 缓冲区没有正确捕获输出")
			testResults = append(testResults, "")
		}

		// 将SOCKS5服务器输出添加到测试记录

		// 使用互斥锁保护对socks5Output的访问
		socks5OutputMutex.Lock()
		socks5OutputLen := socks5Output.Len()
		socks5OutputContent := socks5Output.String()
		socks5OutputMutex.Unlock()

		if socks5OutputLen > 0 {
			testResults = append(testResults, "### SOCKS5服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(socks5OutputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		} else {
			testResults = append(testResults, "### SOCKS5服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "⚠️ 没有捕获到SOCKS5服务器日志")
			testResults = append(testResults, "")

			// 添加调试信息
			testResults = append(testResults, "### 调试信息")
			testResults = append(testResults, "")
			testResults = append(testResults, fmt.Sprintf("SOCKS5服务器输出缓冲区长度: %d", socks5OutputLen))
			testResults = append(testResults, "")
			testResults = append(testResults, "可能的原因:")
			testResults = append(testResults, "- SOCKS5服务器程序没有输出日志")
			testResults = append(testResults, "- 日志输出被重定向到其他地方")
			testResults = append(testResults, "- 缓冲区没有正确捕获输出")
			testResults = append(testResults, "")
		}

		// 重新写入测试记录
		err = writeTestResults3(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}
	}
}

// isSocks5ProxyRunning 检查SOCKS5代理服务器是否正在运行
func isSocks5ProxyRunning() bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 创建一个测试请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return false
	}

	// 设置代理
	proxyURL, err := url.Parse("socks5://localhost:10810")
	if err != nil {
		return false
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client.Transport = transport

	// 发送测试请求
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

// isPortOccupied3 检查端口是否被占用
func isPortOccupied3(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return true
	}
	listener.Close()
	return false
}

// writeTestResults3 写入测试结果到文件
func writeTestResults3(results []string) error {
	// 写入到测试记录.md
	file, err := os.OpenFile("测试记录.md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 移动到文件末尾
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// 写入分隔符
	_, err = writer.WriteString("\n\n###\n\n")
	if err != nil {
		return err
	}

	// 写入测试结果
	for _, line := range results {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// TestMainWebSocket 主测试函数
func TestMainWebSocket(t *testing.T) {
	// 创建带有35秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan bool, 1)

	// 在goroutine中运行测试
	go func() {
		// 运行测试
		runWebSocketSocks5Proxy(t)
		resultChan <- true
	}()

	// 等待测试完成或超时
	select {
	case <-resultChan:
		// 测试正常完成
		return
	case <-ctx.Done():
		// 超时或取消
		log.Println("\n⏰ 测试超时（35秒），强制退出...")

		// 在Windows上强制终止所有go进程
		if runtime.GOOS == "windows" {
			killCmd := exec.Command("taskkill", "/F", "/IM", "go.exe")
			killCmd.Run()
		}

		// 记录超时信息
		timeoutMessage := []string{
			"# WebSocket和SOCKS5级联测试超时记录",
			"",
			"## 超时时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"❌ 测试执行超过35秒超时限制，强制退出",
			"",
			"可能的原因:",
			"- 服务器进程未正常退出",
			"- curl命令卡住",
			"- 网络连接问题",
			"- 级联代理配置问题",
			"",
			"已尝试终止所有相关进程",
			"",
		}

		// 写入超时记录
		if err := writeTestResults3(timeoutMessage); err != nil {
			log.Printf("写入超时记录失败: %v\n", err)
		}

		// 强制退出
		t.Fatal("测试超时")
	}
}

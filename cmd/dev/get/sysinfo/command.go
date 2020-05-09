package sysinfo

import (
	"errors"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.SystemInfoCanonicalNoun,
		Aliases: constants.SystemInfoAliases,
		Short:   "retrieves information about the current host system",
		Run: func(command *cobra.Command, args []string) {
			var info syscall.Sysinfo_t
			syscall.Sysinfo(&info)
			workingDirectory, getwdError := os.Getwd()
			if getwdError != nil {
				workingDirectory = "(unknown)"
			}
			hostname, getHostnameError := os.Hostname()
			if getHostnameError != nil {
				hostname = "(unknown)"
			}
			userHome, getUserHomeDirError := os.UserHomeDir()
			if getUserHomeDirError != nil {
				userHome = "(unknown)"
			}
			log.Printf("timestamp         : %s\n", time.Now().Format(constants.DevHumanTimeFormat))
			log.Printf("home directory    : %s\n", userHome)
			log.Printf("current directory : %s\n", workingDirectory)
			log.Printf("current group id  : %v\n", os.Getegid())
			log.Printf("current user id   : %v\n", os.Getuid())
			log.Printf("hostname          : %s\n", hostname)
			log.Printf("architecture      : %s\n", runtime.GOARCH)
			log.Printf("operating system  : %s\n", runtime.GOOS)
			log.Printf("number of cpus    : %v\n", runtime.NumCPU())

			var totalMemory, swapTotal uint64
			memory, getMemoryError := mem.VirtualMemory()
			if getMemoryError == nil {
				totalMemory = memory.Total
				swapTotal = memory.SwapTotal
				getMemoryError = errors.New("")
			}
			log.Printf("total ram         : %v %s\n", humanize.Bytes(totalMemory), getMemoryError)
			log.Printf("total swap        : %v %s\n", humanize.Bytes(swapTotal), getMemoryError)
		},
	}
	return &cmd
}

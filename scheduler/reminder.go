package scheduler

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/robfig/cron/v3"
)

func StartReminderScheduler() {
	c := cron.New()

	// Schedule: Every day at 3 PM (15:00)
	if _, err := c.AddFunc("0 15 * * *", func() {
		sendReminder("💰 Reminder: Don't forget to log your expenses!")
	}); err != nil {
		fmt.Println("Failed to schedule 3 PM reminder:", err)
	}

	// Schedule: Every day at 9 PM (21:00)
	if _, err := c.AddFunc("0 21 * * *", func() {
		sendReminder("💰 Reminder: Don't forget to log your income/expenses!")
	}); err != nil {
		fmt.Println("Failed to schedule 9 PM reminder:", err)
	}

	c.Start()

	// Keep it running if needed (optional if used in CLI mode)
	go func() {
		select {} // Block forever
	}()
}

func sendReminder(message string) {
	fmt.Println(time.Now().Format("15:04"), "-", message)

	switch runtime.GOOS {
	case "darwin":
		// macOS: AppleScript notification
		exec.Command("osascript", "-e", `display notification "`+message+`" with title "Cash Log"`).Run()
	case "linux":
		// Linux: notify-send
		exec.Command("notify-send", "Cash Log", message).Run()
	case "windows":
		// Windows: fallback
		fmt.Println("🔔 (Reminder - Windows):", message)
	default:
		fmt.Println("Unsupported OS notification. Just printing the message.")
	}
}

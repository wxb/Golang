package storage

import (
	"fmt"
)

func bytesInUse(username string) int64 { return 990000000 /* ... */ }

// Email sender configuration.
// NOTE: never put passwords in source code!
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`

var notifyUser = func(username, msg string) {
	// auth := smtp.PlainAuth("", sender, password, hostname)
	// err := smtp.SendMail(hostname+":587", auth, sender,
	// 	[]string{username}, []byte(msg))
	// if err != nil {
	// 	log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	// }
	fmt.Println("----------")
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		fmt.Println("111111", percent)
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

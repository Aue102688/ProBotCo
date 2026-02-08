package main

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// เริ่ม playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	// เปิด Chromium
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		SlowMo:   playwright.Float(150),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	// context = profile ผู้ใช้
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Viewport: &playwright.Size{
			Width:  1280,
			Height: 800,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	// เว็บ demo login
	_, err = page.Goto("https://www.thaiticketmajor.com/")
	if err != nil {
		log.Fatal(err)
	}

	page.WaitForTimeout(3000)

	log.Println("✅ Automation login สำเร็จ")
}

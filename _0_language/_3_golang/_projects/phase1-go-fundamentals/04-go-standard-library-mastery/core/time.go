package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 Go time Package Mastery Examples")
	fmt.Println("====================================")

	// 1. Current Time
	fmt.Println("\n1. Current Time:")
	now := time.Now()
	fmt.Printf("Current time: %s\n", now)
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	fmt.Printf("Unix micro timestamp: %d\n", now.UnixMicro())
	fmt.Printf("Unix nano timestamp: %d\n", now.UnixNano())

	// 2. Time Creation
	fmt.Println("\n2. Time Creation:")
	
	// Create time from components
	specificTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	fmt.Printf("Specific time: %s\n", specificTime)
	
	// Create time from Unix timestamp
	unixTime := time.Unix(1703514645, 0)
	fmt.Printf("Unix time: %s\n", unixTime)
	
	// Create time from Unix micro timestamp
	unixMicroTime := time.UnixMicro(1703514645000000)
	fmt.Printf("Unix micro time: %s\n", unixMicroTime)

	// 3. Time Formatting
	fmt.Println("\n3. Time Formatting:")
	
	// Standard formats
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("RFC3339Nano: %s\n", now.Format(time.RFC3339Nano))
	fmt.Printf("RFC822: %s\n", now.Format(time.RFC822))
	fmt.Printf("RFC822Z: %s\n", now.Format(time.RFC822Z))
	fmt.Printf("RFC850: %s\n", now.Format(time.RFC850))
	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
	fmt.Printf("RFC1123Z: %s\n", now.Format(time.RFC1123Z))
	
	// Custom formats
	fmt.Printf("Custom format: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Date only: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Time only: %s\n", now.Format("15:04:05"))
	fmt.Printf("12-hour format: %s\n", now.Format("3:04:05 PM"))
	fmt.Printf("With timezone: %s\n", now.Format("2006-01-02 15:04:05 MST"))

	// 4. Time Parsing
	fmt.Println("\n4. Time Parsing:")
	
	// Parse RFC3339
	parsedTime, err := time.Parse(time.RFC3339, "2023-12-25T15:30:45Z")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed RFC3339: %s\n", parsedTime)
	}
	
	// Parse custom format
	customTime, err := time.Parse("2006-01-02 15:04:05", "2023-12-25 15:30:45")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed custom: %s\n", customTime)
	}
	
	// Parse in specific location
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("Location error: %v\n", err)
	} else {
		localTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-12-25 15:30:45", loc)
		if err != nil {
			fmt.Printf("Parse error: %v\n", err)
		} else {
			fmt.Printf("Parsed in NY timezone: %s\n", localTime)
		}
	}

	// 5. Time Zones
	fmt.Println("\n5. Time Zones:")
	
	// UTC time
	utcTime := now.UTC()
	fmt.Printf("UTC time: %s\n", utcTime)
	
	// Local time
	localTime := now.Local()
	fmt.Printf("Local time: %s\n", localTime)
	
	// Specific timezone
	nyLoc, _ := time.LoadLocation("America/New_York")
	nyTime := now.In(nyLoc)
	fmt.Printf("NY time: %s\n", nyTime)
	
	// Tokyo timezone
	tokyoLoc, _ := time.LoadLocation("Asia/Tokyo")
	tokyoTime := now.In(tokyoLoc)
	fmt.Printf("Tokyo time: %s\n", tokyoTime)
	
	// List available timezones
	fmt.Println("Available timezones:")
	zones := []string{"UTC", "America/New_York", "Europe/London", "Asia/Tokyo", "Australia/Sydney"}
	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err == nil {
			fmt.Printf("  %s: %s\n", zone, now.In(loc).Format("2006-01-02 15:04:05 MST"))
		}
	}

	// 6. Duration
	fmt.Println("\n6. Duration:")
	
	// Create durations
	duration1 := 5 * time.Minute
	duration2 := 2 * time.Hour
	duration3 := 30 * time.Second
	duration4 := 500 * time.Millisecond
	
	fmt.Printf("5 minutes: %s\n", duration1)
	fmt.Printf("2 hours: %s\n", duration2)
	fmt.Printf("30 seconds: %s\n", duration3)
	fmt.Printf("500 milliseconds: %s\n", duration4)
	
	// Duration arithmetic
	totalDuration := duration1 + duration2 + duration3 + duration4
	fmt.Printf("Total duration: %s\n", totalDuration)
	
	// Duration components
	fmt.Printf("Hours: %.2f\n", totalDuration.Hours())
	fmt.Printf("Minutes: %.2f\n", totalDuration.Minutes())
	fmt.Printf("Seconds: %.2f\n", totalDuration.Seconds())
	fmt.Printf("Milliseconds: %.2f\n", float64(totalDuration.Nanoseconds())/1e6)
	fmt.Printf("Nanoseconds: %d\n", totalDuration.Nanoseconds())

	// 7. Time Operations
	fmt.Println("\n7. Time Operations:")
	
	// Add duration to time
	futureTime := now.Add(24 * time.Hour)
	fmt.Printf("Now: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("24 hours later: %s\n", futureTime.Format("2006-01-02 15:04:05"))
	
	// Subtract time from time
	pastTime := now.Add(-7 * 24 * time.Hour)
	fmt.Printf("7 days ago: %s\n", pastTime.Format("2006-01-02 15:04:05"))
	
	// Time comparison
	fmt.Printf("Is future time after now? %t\n", futureTime.After(now))
	fmt.Printf("Is past time before now? %t\n", pastTime.Before(now))
	fmt.Printf("Are times equal? %t\n", now.Equal(now))
	
	// Time difference
	diff := futureTime.Sub(now)
	fmt.Printf("Time difference: %s\n", diff)

	// 8. Timer
	fmt.Println("\n8. Timer:")
	
	// One-time timer
	fmt.Println("Starting 2-second timer...")
	timer := time.NewTimer(2 * time.Second)
	
	// Wait for timer
	<-timer.C
	fmt.Println("Timer fired!")
	
	// Timer with timeout
	fmt.Println("Starting timer with timeout...")
	timer2 := time.NewTimer(1 * time.Second)
	select {
	case <-timer2.C:
		fmt.Println("Timer completed")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout occurred")
		timer2.Stop()
	}

	// 9. Ticker
	fmt.Println("\n9. Ticker:")
	
	// Create ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	// Run ticker for 3 seconds
	timeout := time.After(3 * time.Second)
	count := 0
	
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Ticker %d: %s\n", count, time.Now().Format("15:04:05.000"))
		case <-timeout:
			fmt.Println("Ticker stopped")
			goto done
		}
	}
done:

	// 10. Sleep
	fmt.Println("\n10. Sleep:")
	fmt.Println("Sleeping for 1 second...")
	start := time.Now()
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("Slept for: %s\n", elapsed)

	// 11. AfterFunc
	fmt.Println("\n11. AfterFunc:")
	
	// Execute function after delay
	timer3 := time.AfterFunc(1*time.Second, func() {
		fmt.Println("AfterFunc executed!")
	})
	
	// Wait for execution
	time.Sleep(2 * time.Second)
	
	// Stop timer if not fired
	if !timer3.Stop() {
		fmt.Println("Timer was already fired")
	}

	// 12. Time Truncation
	fmt.Println("\n12. Time Truncation:")
	
	// Truncate to different units
	fmt.Printf("Original: %s\n", now.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Truncated to hour: %s\n", now.Truncate(time.Hour).Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Truncated to minute: %s\n", now.Truncate(time.Minute).Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Truncated to second: %s\n", now.Truncate(time.Second).Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Truncated to millisecond: %s\n", now.Truncate(time.Millisecond).Format("2006-01-02 15:04:05.000000000"))

	// 13. Time Rounding
	fmt.Println("\n13. Time Rounding:")
	
	// Round to different units
	fmt.Printf("Original: %s\n", now.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Rounded to hour: %s\n", now.Round(time.Hour).Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Rounded to minute: %s\n", now.Round(time.Minute).Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Rounded to second: %s\n", now.Round(time.Second).Format("2006-01-02 15:04:05.000000000"))

	// 14. Duration Parsing
	fmt.Println("\n14. Duration Parsing:")
	
	// Parse duration strings
	durations := []string{"1h30m", "45m", "30s", "500ms", "1.5h", "2h45m30s"}
	
	for _, durStr := range durations {
		duration, err := time.ParseDuration(durStr)
		if err != nil {
			fmt.Printf("Parse error for %s: %v\n", durStr, err)
		} else {
			fmt.Printf("%s = %s (%.2f hours)\n", durStr, duration, duration.Hours())
		}
	}

	// 15. Time Constants
	fmt.Println("\n15. Time Constants:")
	
	// Common time constants
	fmt.Printf("Nanosecond: %s\n", time.Nanosecond)
	fmt.Printf("Microsecond: %s\n", time.Microsecond)
	fmt.Printf("Millisecond: %s\n", time.Millisecond)
	fmt.Printf("Second: %s\n", time.Second)
	fmt.Printf("Minute: %s\n", time.Minute)
	fmt.Printf("Hour: %s\n", time.Hour)
	fmt.Printf("Day: %s\n", 24*time.Hour)
	fmt.Printf("Week: %s\n", 7*24*time.Hour)
	fmt.Printf("Year: %s\n", 365*24*time.Hour)

	// 16. Time Marshaling
	fmt.Println("\n16. Time Marshaling:")
	
	// Marshal to JSON
	jsonData, err := now.MarshalJSON()
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
	} else {
		fmt.Printf("JSON: %s\n", string(jsonData))
	}
	
	// Marshal to text
	textData, err := now.MarshalText()
	if err != nil {
		fmt.Printf("Text marshal error: %v\n", err)
	} else {
		fmt.Printf("Text: %s\n", string(textData))
	}

	// 17. Time Unmarshaling
	fmt.Println("\n17. Time Unmarshaling:")
	
	// Unmarshal from JSON
	var unmarshaledTime time.Time
	err = unmarshaledTime.UnmarshalJSON(jsonData)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
	} else {
		fmt.Printf("Unmarshaled from JSON: %s\n", unmarshaledTime.Format("2006-01-02 15:04:05"))
	}

	// 18. Time Components
	fmt.Println("\n18. Time Components:")
	
	// Extract time components
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %d (%s)\n", now.Month(), now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	fmt.Printf("Hour: %d\n", now.Hour())
	fmt.Printf("Minute: %d\n", now.Minute())
	fmt.Printf("Second: %d\n", now.Second())
	fmt.Printf("Nanosecond: %d\n", now.Nanosecond())
	fmt.Printf("Weekday: %d (%s)\n", now.Weekday(), now.Weekday())
	fmt.Printf("YearDay: %d\n", now.YearDay())
	year, week := now.ISOWeek()
	fmt.Printf("ISOWeek: %d, %d\n", year, week)

	// 19. Time Comparison with Different Units
	fmt.Println("\n19. Time Comparison with Different Units:")
	
	// Compare times with different precision
	time1 := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)
	time2 := time.Date(2023, 12, 25, 15, 30, 45, 987654321, time.UTC)
	
	fmt.Printf("Time1: %s\n", time1.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Time2: %s\n", time2.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("Equal (nanosecond precision): %t\n", time1.Equal(time2))
	fmt.Printf("Equal (second precision): %t\n", time1.Truncate(time.Second).Equal(time2.Truncate(time.Second)))
	fmt.Printf("Equal (minute precision): %t\n", time1.Truncate(time.Minute).Equal(time2.Truncate(time.Minute)))

	// 20. Practical Examples
	fmt.Println("\n20. Practical Examples:")
	
	// Stopwatch
	fmt.Println("Stopwatch example:")
	start = time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed = time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
	
	// Timeout for operation
	fmt.Println("Timeout example:")
	done := make(chan bool)
	go func() {
		time.Sleep(500 * time.Millisecond)
		done <- true
	}()
	
	select {
	case <-done:
		fmt.Println("Operation completed")
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
	
	// Rate limiting
	fmt.Println("Rate limiting example:")
	rateLimiter := time.NewTicker(100 * time.Millisecond)
	defer rateLimiter.Stop()
	
	for i := 0; i < 5; i++ {
		<-rateLimiter.C
		fmt.Printf("Rate limited operation %d\n", i+1)
	}

	fmt.Println("\n🎉 time Package Mastery Complete!")
}

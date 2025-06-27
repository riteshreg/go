package main

import (
	"fmt"
	"time"
)

// blocking mean un-buffered

func XD1() {
	fmt.Printf("=== Understanding 'Blocks Until Receiver is Ready' ===\n")

	// Create an unbuffered channel
	ch := make(chan string)

	fmt.Println("1. Creating unbuffered channel...")

	// This goroutine will send a message
	go func() {
		fmt.Println("2. Goroutine: About to send message...")
		ch <- "Hello from goroutine!" // This BLOCKS until someone receives
		fmt.Println("3. Goroutine: Message sent successfully!")
	}()

	fmt.Println("4. Main: Sleeping for 2 seconds to show blocking...")
	time.Sleep(2 * time.Second)

	fmt.Println("5. Main: About to receive message...")
	message := <-ch // Now the send can complete
	fmt.Printf("6. Main: Received: %s\n", message)

	fmt.Println("\n=== What Happened ===")
	fmt.Println("• Step 2: Goroutine tries to send, but blocks")
	fmt.Println("• Step 4: Main sleeps while goroutine is still blocked")
	fmt.Println("• Step 5: Main receives, which unblocks the goroutine")
	fmt.Println("• Step 3: Goroutine can now continue and print 'Message sent'")

	fmt.Println("\n=== Try This: Comment out the receive to see deadlock ===")
	fmt.Println("If you comment out line 25 (the receive), the program will deadlock!")
	fmt.Println("Because the goroutine will be forever blocked waiting for a receiver.")
}

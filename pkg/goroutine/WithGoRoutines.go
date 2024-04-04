package goroutine

import (
	"fmt"
	"time"
)

func countToOneBillion() {
    start := time.Now()
    for i := 0; i < 1000000000; i++ {
        // Do nothing inside the loop, just count
    }
    elapsed := time.Since(start)
    fmt.Printf("Counting to 1 billion took %s\n", elapsed)
}
func countToNineBillion() {
    start := time.Now()
    for i := 0; i < 900000000000; i++ {
        // Do nothing inside the loop, just count
    }
    elapsed := time.Since(start)
    fmt.Printf("Counting to 9 billion took %s\n", elapsed)
}


func TestGoRoutine() {

    start := time.Now()

    // Run 3 instances of countToOneBillion in parallel using goroutines
    for i := 0; i < 3; i++ {
        go countToOneBillion()
        // go countToTenBillion()
    }

    // Wait for all goroutines to finish
    time.Sleep(1 * time.Second) // Adjust sleep time if needed based on system performance

    totalTime := time.Since(start)
    fmt.Printf("\nTotal time for all 3 instances: %s\n", totalTime)
}
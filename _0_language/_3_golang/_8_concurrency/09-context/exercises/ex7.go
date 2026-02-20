package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 7: Context Value Propagation
func Exercise7() {
	fmt.Println("\nExercise 7: Context Value Propagation")
	fmt.Println("=====================================")
	
	// TODO: Create context with values and propagate through functions
	// 1. Create context with initial values
	// 2. Pass through multiple function calls
	// 3. Add values at each level
	// 4. Extract values at the end
	
	ctx := context.Background()
	ctx = addValue(ctx, "level", "root")
	ctx = addValue(ctx, "requestID", "req-001")
	
	ctx = processLevel1(ctx)
	ctx = processLevel2(ctx)
	ctx = processLevel3(ctx)
	
	// Extract all values
	fmt.Printf("  Exercise 7: Final values: %v\n", getAllValues(ctx))
	fmt.Println("Exercise 7 completed")
}

func addValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func processLevel1(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level1", "processed")
	ctx = addValue(ctx, "timestamp1", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 1 processed")
	return ctx
}

func processLevel2(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level2", "processed")
	ctx = addValue(ctx, "timestamp2", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 2 processed")
	return ctx
}

func processLevel3(ctx context.Context) context.Context {
	ctx = addValue(ctx, "level3", "processed")
	ctx = addValue(ctx, "timestamp3", time.Now().Format("15:04:05"))
	fmt.Println("  Exercise 7: Level 3 processed")
	return ctx
}

func getAllValues(ctx context.Context) map[string]interface{} {
	values := make(map[string]interface{})
	
	// Extract known values
	if level := ctx.Value("level"); level != nil {
		values["level"] = level
	}
	if requestID := ctx.Value("requestID"); requestID != nil {
		values["requestID"] = requestID
	}
	if level1 := ctx.Value("level1"); level1 != nil {
		values["level1"] = level1
	}
	if level2 := ctx.Value("level2"); level2 != nil {
		values["level2"] = level2
	}
	if level3 := ctx.Value("level3"); level3 != nil {
		values["level3"] = level3
	}
	
	return values
}
package main

import (
	"fmt"
	"time"
)

// Example structs for demonstration
type SmallStruct struct {
	ID   int
	Name string
}

type LargeStruct struct {
	ID          int
	Name        string
	Description string
	Data        [1000]int // Large array
	Timestamp   time.Time
	// ... imagine many more fields
}

// Method with pointer receiver (can modify the struct)
func (s *SmallStruct) UpdateName(newName string) {
	s.Name = newName // This modifies the original struct
}

// Method with value receiver (cannot modify the struct)
func (s SmallStruct) GetInfo() string {
	return fmt.Sprintf("ID: %d, Name: %s", s.ID, s.Name)
}

// Function that takes a pointer (can modify the original)
func updateStructByPointer(s *SmallStruct, newName string) {
	s.Name = newName
}

// Function that takes a value (works on a copy)
func updateStructByValue(s SmallStruct, newName string) {
	s.Name = newName // This only modifies the copy!
}

func demonstratePointerRules() {
	fmt.Println("🎯 When to Use Pointers vs Values in Go")
	fmt.Println("=====================================")
	
	// Rule 1: Use pointers when you need to MODIFY the original
	fmt.Println("\n1. MODIFICATION - Use pointers to modify originals:")
	
	original := SmallStruct{ID: 1, Name: "Original"}
	fmt.Printf("   Before: %+v\n", original)
	
	// Using pointer - modifies original
	updateStructByPointer(&original, "Modified by pointer")
	fmt.Printf("   After pointer update: %+v\n", original)
	
	// Using value - doesn't modify original
	updateStructByValue(original, "Modified by value")
	fmt.Printf("   After value update: %+v (unchanged!)\n", original)
	
	// Rule 2: Large structs should use pointers for efficiency
	fmt.Println("\n2. EFFICIENCY - Use pointers for large structs:")
	
	large := LargeStruct{ID: 1, Name: "Large"}
	fmt.Printf("   Large struct created: %s\n", large.Name)
	fmt.Printf("   Size of LargeStruct: ~%d bytes\n", 8000) // Approximate
	fmt.Printf("   Size of pointer: 8 bytes\n")
	fmt.Println("   ✅ Always pass large structs by pointer!")
	
	// Rule 3: Small structs can be passed by value
	fmt.Println("\n3. SMALL STRUCTS - Values are often fine:")
	
	small := SmallStruct{ID: 1, Name: "Small"}
	info := small.GetInfo() // Value receiver is fine for small, read-only operations
	fmt.Printf("   %s\n", info)
	fmt.Println("   ✅ Small structs can use value receivers for read-only methods")
	
	// Rule 4: Method receivers
	fmt.Println("\n4. METHOD RECEIVERS:")
	fmt.Println("   • Use pointer receivers (*T) when:")
	fmt.Println("     - Method needs to modify the struct")
	fmt.Println("     - Struct is large")
	fmt.Println("     - You want consistency (if any method uses pointer, use pointer everywhere)")
	fmt.Println("   • Use value receivers (T) when:")
	fmt.Println("     - Method only reads data")
	fmt.Println("     - Struct is small")
	fmt.Println("     - Struct is immutable by design")
}

func showCommonPatterns() {
	fmt.Println("\n🔧 Common Go Patterns:")
	fmt.Println("=====================")
	
	fmt.Println("\n✅ USE POINTERS for:")
	fmt.Println("   • Structs you need to modify")
	fmt.Println("   • Large structs (>100 bytes)")
	fmt.Println("   • Database models")
	fmt.Println("   • HTTP handlers (http.Server, http.Request)")
	fmt.Println("   • Slices, maps, channels (already reference types)")
	
	fmt.Println("\n✅ USE VALUES for:")
	fmt.Println("   • Small structs (<100 bytes)")
	fmt.Println("   • Immutable data")
	fmt.Println("   • Configuration structs (read-only)")
	fmt.Println("   • Basic types (int, string, bool)")
	fmt.Println("   • When you want a copy, not the original")
	
	fmt.Println("\n⚠️  SPECIAL CASES:")
	fmt.Println("   • Slices: Already reference types, don't need *[]T")
	fmt.Println("   • Maps: Already reference types, don't need *map[K]V")
	fmt.Println("   • Channels: Already reference types, don't need *chan T")
	fmt.Println("   • Interfaces: Usually passed by value")
}

func showRealWorldExamples() {
	fmt.Println("\n🌍 Real-World Examples:")
	fmt.Println("======================")
	
	fmt.Println("\n// Database model - USE POINTER")
	fmt.Println("type User struct {")
	fmt.Println("    ID       int")
	fmt.Println("    Name     string")
	fmt.Println("    Email    string")
	fmt.Println("    // ... many fields")
	fmt.Println("}")
	fmt.Println("func (u *User) Save() error { /* modify u */ }")
	
	fmt.Println("\n// Configuration - USE VALUE")
	fmt.Println("type Config struct {")
	fmt.Println("    Port    int")
	fmt.Println("    Debug   bool")
	fmt.Println("}")
	fmt.Println("func (c Config) GetPort() int { return c.Port }")
	
	fmt.Println("\n// HTTP Server - USE POINTER (required)")
	fmt.Println("server := &http.Server{")
	fmt.Println("    Addr: \":8080\",")
	fmt.Println("}")
	fmt.Println("// ListenAndServe() needs to modify server state")
	
	fmt.Println("\n// Small coordinate - USE VALUE")
	fmt.Println("type Point struct {")
	fmt.Println("    X, Y int")
	fmt.Println("}")
	fmt.Println("func (p Point) Distance() float64 { /* read-only */ }")
}

// RunPointerGuide runs the pointer demonstration
// Call this function from main() if you want to see the guide
// Example: go run . --guide
func RunPointerGuide() {
	demonstratePointerRules()
	showCommonPatterns()
	showRealWorldExamples()
}

// Uncomment the lines below and comment out the main() in main.go to run this guide:
//
// func main() {
// 	RunPointerGuide()
// }

// 代码生成时间: 2025-07-31 08:12:49
package main

import (
    "fmt"
    "sort"
)

// SortingService is a service that provides sorting functionality.
type SortingService struct{}

// SortInts sorts a slice of integers in ascending order.
func (s *SortingService) SortInts(ints []int) ([]int, error) {
    // Check if the input slice is valid
    if ints == nil {
        return nil, fmt.Errorf("input slice cannot be nil")
    }
    
    // Create a copy of the slice to avoid modifying the original slice
    copyInts := make([]int, len(ints))
    copy(copyInts, ints)
    
    // Sort the slice in ascending order
    sort.Ints(copyInts)
    
    return copyInts, nil
}

// SortFloats sorts a slice of floats in ascending order.
func (s *SortingService) SortFloats(floats []float64) ([]float64, error) {
    // Check if the input slice is valid
    if floats == nil {
        return nil, fmt.Errorf("input slice cannot be nil")
    }
    
    // Create a copy of the slice to avoid modifying the original slice
    copyFloats := make([]float64, len(floats))
    copy(copyFloats, floats)
    
    // Sort the slice in ascending order
    sort.Float64s(copyFloats)
    
    return copyFloats, nil
}

func main() {
    // Create an instance of SortingService
    service := SortingService{}
    
    // Example usage of SortInts
    ints := []int{5, 2, 9, 1, 5, 6}
    sortedInts, err := service.SortInts(ints)
    if err != nil {
        fmt.Println("Error sorting integers: ", err)
    } else {
        fmt.Println("Sorted integers: ", sortedInts)
    }
    
    // Example usage of SortFloats
    floats := []float64{3.5, 2.1, 4.8, 1.9}
    sortedFloats, err := service.SortFloats(floats)
    if err != nil {
        fmt.Println("Error sorting floats: ", err)
    } else {
        fmt.Println("Sorted floats: ", sortedFloats)
    }
}
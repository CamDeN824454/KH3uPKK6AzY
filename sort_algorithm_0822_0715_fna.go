// 代码生成时间: 2025-08-22 07:15:58
package main
# 增强安全性

import (
    "fmt"
    "sort"
)

// BubbleSort is a function that performs bubble sort on a slice of integers
func BubbleSort(nums []int) ([]int, error) {
    // Check if the input is nil
    if nums == nil {
        return nil, fmt.Errorf("input slice is nil")
    }
# 增强安全性

    n := len(nums)
# 增强安全性
    for i := 0; i < n-1; i++ {
# 扩展功能模块
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                // Swap the elements
                nums[j], nums[j+1] = nums[j+1], nums[j]
# NOTE: 重要实现细节
            }
        }
# TODO: 优化性能
    }
# 增强安全性
    return nums, nil
}

// SelectionSort is a function that performs selection sort on a slice of integers
func SelectionSort(nums []int) ([]int, error) {
    // Check if the input is nil
    if nums == nil {
        return nil, fmt.Errorf("input slice is nil")
    }

    n := len(nums)
    for i := 0; i < n-1; i++ {
        minIndex := i
# 扩展功能模块
        for j := i+1; j < n; j++ {
            if nums[j] < nums[minIndex] {
                minIndex = j
            }
        }
        // Swap the found minimum element with the first element
        nums[i], nums[minIndex] = nums[minIndex], nums[i]
    }
    return nums, nil
}

// InsertionSort is a function that performs insertion sort on a slice of integers
func InsertionSort(nums []int) ([]int, error) {
    // Check if the input is nil
    if nums == nil {
        return nil, fmt.Errorf("input slice is nil")
    }

    for i := 1; i < len(nums); i++ {
# 改进用户体验
        key := nums[i]
        j := i - 1
        // Move elements of nums[0..i-1], that are greater than key, to one position ahead
        for j >= 0 && nums[j] > key {
            nums[j+1] = nums[j]
            j--
# 改进用户体验
        }
        nums[j+1] = key
    }
# 添加错误处理
    return nums, nil
}

func main() {
# NOTE: 重要实现细节
    // Example usage of sorting algorithms
    nums := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Original array: ", nums)

    sortedNums, err := BubbleSort(nums)
# 改进用户体验
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
# 增强安全性
        fmt.Println("Sorted array (Bubble Sort): ", sortedNums)
    }

    nums = []int{64, 34, 25, 12, 22, 11, 90}
    sortedNums, err = SelectionSort(nums)
# NOTE: 重要实现细节
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Sorted array (Selection Sort): ", sortedNums)
    }

    nums = []int{64, 34, 25, 12, 22, 11, 90}
    sortedNums, err = InsertionSort(nums)
    if err != nil {
        fmt.Println("Error: ", err)
# 改进用户体验
    } else {
        fmt.Println("Sorted array (Insertion Sort): ", sortedNums)
    }
}
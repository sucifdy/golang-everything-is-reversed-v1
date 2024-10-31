package main

import "fmt"

// ReverseData mengubah urutan dan isi dari array menjadi terbalik.
func ReverseData(arr [5]int) [5]int {
    var result [5]int

    // Membalik urutan array
    for i := 0; i < 5; i++ {
        // Membalik isi setiap angka
        reversedNumber := reverseInt(arr[4-i])
        result[i] = reversedNumber
    }

    return result
}

// reverseInt membalik angka menjadi integer.
func reverseInt(num int) int {
    reversed := 0
    for num > 0 {
        digit := num % 10
        reversed = reversed*10 + digit
        num /= 10
    }
    return reversed
}

func main() {
    // Contoh penggunaan
    fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))         // Output: [2, 1, 11, 654, 321]
    fmt.Println(ReverseData([5]int{456789, 44332, 2221, 12, 10})) // Output: [1, 21, 1222, 23344, 987654]
    fmt.Println(ReverseData([5]int{10, 10, 10, 10, 10}))         // Output: [1, 1, 1, 1, 1]
    fmt.Println(ReverseData([5]int{23456, 789, 123, 456, 500}))  // Output: [5, 654, 321, 987, 65432]
    fmt.Println(ReverseData([5]int{0, 0, 0, 0, 0}))              // Output: [0, 0, 0, 0, 0]
}

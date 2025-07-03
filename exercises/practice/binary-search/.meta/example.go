package binarysearch

// SearchInts performs a binary search for the given key in the sorted slice.
// Returns the index of the key if found, otherwise returns -1.
// The input slice must be sorted in ascending order.
func SearchInts(slice []int, key int) int {
    if len(slice) == 0 {
        return -1
    }

    left := 0
    right := len(slice) - 1

    for left <= right {
        mid := (left + right) / 2

        if slice[mid] == key {
            return mid
        }

        if slice[mid] < key {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}


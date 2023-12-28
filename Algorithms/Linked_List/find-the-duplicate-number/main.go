package main

func findDuplicate(nums []int) int {
    // Step 1: Find the intersection point of the two pointers
    slow := nums[0]
    fast := nums[0]

    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }
    }

    // Step 2: Find the entrance to the cycle (duplicate number)
    slow = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}

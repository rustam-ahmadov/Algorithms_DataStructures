#include <iostream>
#include <vector>

int search(std::vector<int> &nums, int target)
{
    const int n = nums.size();

    int l = 0, r = n - 1;
    int m = (r + l) / 2;

    while (l <= r)
    {
        if (nums[m] == target)
            return m;

        if (nums[l] <= nums[m]) // left sorted portion
        {
            if (target > nums[m] || target < nums[l]) // [2, 3, 4, 5m, 6, 7t, 1] || [3, 4, 5, 6m, 7, 1t, 2]
                l = m + 1;
            else // [2, 3, 4t, 5m, 6, 7, 1]
                r = m - 1;
        }
        else                                          // right sorted portion
            if (target < nums[m] || target > nums[r]) // [8, 2, 3t, 4m, 5, 6, 7] || // [8m, 2, 3, 4m, 5, 6, 7]
                r = m - 1;
            else // [8, 2, 3, 4m, 5, 6t, 7]
                l = m + 1;
        m = (r + l) / 2;
    }
    return -1;
}

int main()
{
    std::vector<int> nums1 = {2, 3, 4, 5, 6, 1};
    std::vector<int> nums2 = {5, 1, 2, 3, 4};
    std::vector<int> nums3 = {4, 5, 6, 7, 0, 1, 2};
    std::vector<int> nums4 = {3, 1};
    std::vector<int> nums5 = {7, 8, 1, 2, 3, 4, 5, 6};

    const int target = 1;

    int t = search(nums2, target);
    std::cout << t;

    return 0;
}
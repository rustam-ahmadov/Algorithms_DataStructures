#include <iostream>
#include <vector>

int findMin(std::vector<int> &nums)
{
    int n = nums.size();

    if (n == 1)
        return nums[0];

    int firstI = 0;
    int midI = n / 2;
    int lastI = n - 1;

    while (firstI != lastI)
    {
        if (nums[firstI] < nums[lastI])
            return nums[firstI];

        if (firstI == lastI - 1)
            return nums[firstI] < nums[lastI] ? nums[firstI] : nums[lastI];

        if (nums[midI] < nums[firstI])
        {
            firstI++;
            lastI = midI;
            midI = (lastI + firstI) / 2;
        }
        else
        {
            firstI = midI;
            midI = (lastI + midI) / 2;
        }
    }
    return nums[firstI];
}

int main()
{
    std::vector<int> nums = { 2,3,1};
    int min = findMin(nums);
    std::cout << min;

    return 0;
}
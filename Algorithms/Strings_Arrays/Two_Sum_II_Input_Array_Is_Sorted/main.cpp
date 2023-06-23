#include <iostream>
#include <vector>

void print(std::vector<int> &nums)
{
    const int n = nums.size();
    for (int i = 0; i < n; i++)
        std::cout << nums[i] << ' ';
    std::cout << std::endl;
}

std::vector<int> twoSum(std::vector<int> &nums, int target)
{
    const int n = nums.size();
    std::vector<int> res;
    for (int l = 0, r = n - 1; l < r;)
    {
        int lNum = nums[l];
        int rNum = nums[r];
        int sum = lNum + rNum;

        if (sum == target)
        {
            res.push_back(l + 1);
             res.push_back(r + 1);
            break;
        }
        else if (sum > target)
            r--;
        else
            l++;
    }
    return res;
}

int main()
{
    std::vector<int> nums = {2, 7, 11, 15};
    std::vector<int> nums1 = {-1, 0};
    const int target = -1;

    std::vector<int> result = twoSum(nums1, target);
    print(result);
}
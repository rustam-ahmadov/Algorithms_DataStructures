#include <iostream>
#include <vector>

void print(std::vector<int> &nums)
{
    int n = nums.size();
    for (int i = 0; i < n; i++)
        std::cout << nums[i] << ' ';
    std::cout << std::endl;
}
std::vector<int> sortedSquares(std::vector<int> &nums)
{
    const int n = nums.size();
    std::vector<int> squaredNums(n);
    if (n == 0)
        return squaredNums;

    for (int i = n - 1, l = 0, r = i; i >= 0; i--)
    {
        int lNum = nums[l];
        int rNum = nums[r];

        int squaredLeft = lNum * lNum;
        int squaredRight = rNum * rNum;
        if (squaredLeft > squaredRight)
        {
            squaredNums[i] = squaredLeft;
            l++;
        }
        else
        {
            squaredNums[i] = squaredRight;
            r--;
        }        
    }
    return squaredNums;
}

int main()
{
    std::vector<int> nums = {-4, -1, 0, 3, 10}; // output: [0,1,9,16,100]
    std::vector<int> squaredNums = sortedSquares(nums);

    print(squaredNums);


    return 0;
}
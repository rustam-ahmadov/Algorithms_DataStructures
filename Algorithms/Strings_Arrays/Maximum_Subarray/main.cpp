#include <iostream>
#include <vector>

void print(std::vector<int> arr, int i)
{
    if (i < 0)
        return;
    else
        print(arr, i - 1);
    std::cout << arr[i] << ' ';
}

void printReverse(std::vector<int> arr, int i)
{
    if (i < 0)
        return;
    std::cout << arr[i] << ' ';
    printReverse(arr, i - 1);
}

// Bruteforce O(n^3)
int maxSubArray(std::vector<int> &nums)
{
    int n = nums.size();
    if (n == 1)
        return nums[1];
    int max = 0;
    for (int i = 0; i < n; i++)
        for (int j = i; j < n; j++)
        {
            int maxPerCycle = 0;
            for (int k = i; k <= j; k++)
            {
                maxPerCycle += nums[k];
            }
            max = max < maxPerCycle ? maxPerCycle : max;
        }
    return max;
}

// Bruteforce O(n^2)
int maxSubArrayN2(std::vector<int> &nums)
{
    int n = nums.size();
    if (n == 1)
        return nums[1];
    int max = 0;
    for (int i = 0; i < n; i++)
    {
        int currentMax = 0;
        for (int j = i; j < n; j++)
        {
            currentMax += nums[j];
        }
        if (max < currentMax)
            max = currentMax;
    }

    return max;
}

int maxSubArrayKadaneS(std::vector<int> &nums)
{
    int n = nums.size();
    if (n == 1)
        return nums[0];

    int max = INT_MIN, curSum = 0;
    for (int i = 0; i < n; i++)
    {
        if (curSum < 0)
            curSum = 0;
        curSum += nums[i];
        max = max < curSum ? curSum : max;
    }
    return max;
}

int main()
{
    std::vector<int> arr = {-2, -1};

    print(arr, arr.size() - 1);
    std::cout << '\n';

    printReverse(arr, arr.size() - 1);
    std::cout << '\n';

    int max = maxSubArrayKadaneS(arr);
    std::cout << "Max: " << max;
}
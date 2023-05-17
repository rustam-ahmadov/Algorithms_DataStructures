#include <iostream>
#include <vector>

using namespace std;

// O(n2)
vector<int> twoSum(vector<int> &nums, int target)
{
    int sum;
    vector<int> result;
    for (int i = 0; i < nums.size(); i++)
    {
        for (int j = i + 1; j < nums.size(); j++)
        {
            sum = nums[i] + nums[j];
            if (sum == target)
            {

                result.push_back(i);
                result.push_back(j);
                break;
            }
        }
    }
    return result;
}

int main()
{
    std::vector<int> vector{1, 3, 7, 11, 20, 30};
    int target = 12;

    std::vector<int> result = twoSum(vector, target);

    return 0;
}
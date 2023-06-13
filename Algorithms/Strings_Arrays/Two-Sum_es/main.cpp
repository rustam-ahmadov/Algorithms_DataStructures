#include <iostream>
#include <vector>
#include <map>

using namespace std;

// O(n2)
vector<int> twoSum(vector<int> &nums, int target)
{
    int sum;
    vector<int> result;
    const int n = nums.size();
    for (int i = 0; i < n - 1; i++)
    {
        for (int j = i + 1; j < n; j++)
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

// O(n)
vector<int> twoSumN(const std::vector<int> &nums, const int target)
{
    int sum = 0;
    const int n = nums.size();
    std::vector<int> result;
    std::map<int, int> table;

    for (int i = 0; i < n; i++)
    {
        int num = target - nums[i];
        if (table[nums[i]] == 0)
            table[num] = i+1;
        else
        {
            result.push_back(table[nums[i]] - 1);
            result.push_back(i);
            break;
        }
    }
    return result;
}

void print(std::vector<int> array)
{
    const int n = array.size();
    for(int i = 0; i < n; i++)
        std::cout<<array[i]<<' ';
    std::cout<<'\n';
}

int main()
{
    std::vector<int> vector{1, 3, 7, 11, 20, 30};
    int target = 12;

    std::vector<int> result = twoSumN(vector, target);
    print(result);
    return 0;
}
#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

bool containsNearbyDuplicate(vector<int> &nums, int k)
{
  unordered_map<int, int> mp;
  int n = nums.size();

  for (int i = 0; i < n; i++)
  {
    if (mp.count(nums[i]))
    { 
      if (i -  mp[nums[i]] <= k)
        return true;
    }
    mp[nums[i]] = i ;
  }
  return false;
}

int main()
{

  vector<int> nums = {1, 0, 1, 1};
  int k = 1;
  std::cout << containsNearbyDuplicate(nums, k);

  return 0;
}
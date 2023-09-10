#include <iostream>
#include <vector>
#include <unordered_map>

// O(n^2)
// bool containsDuplicate(std::vector<int> &nums)
// {
//   int n = nums.size();
//   for (int i = 0; i < n - 1; i++)
//     for (int j = i + 1; j < n; j++)
//       if (nums[i] == nums[j])
//         return true;

//   return false;
// }

//O(n) using hashTable
bool containsDuplicate(std::vector<int> & nums)
{
    std::unordered_map<int,int> table;
    // storing elements 
    int n = nums.size();
    for(int i = 0; i < n; i++ )
      table[nums[i]]++;
    
    for(int i = 0; i < n; i++)
      if(table[nums[i]] > 1)
        return true;
    return false;
}


int main()
{

  std::vector<int> v = {1, 2, 3, 3, 4, 5};
  bool isContains = containsDuplicate(v);
  std::cout << isContains;

  return 0;
}
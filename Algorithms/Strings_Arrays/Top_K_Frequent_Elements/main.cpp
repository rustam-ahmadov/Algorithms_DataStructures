#include <iostream>
#include <unordered_map>
#include <vector>
#include <algorithm>
#include <queue>

using namespace std;

vector<int> topKFrequent(vector<int> &nums, int k)
{

  const int n = nums.size();
  vector<int> res;
  if (!n)
    return res;

  unordered_map<int, int> table(n);
  for (int i = 0; i < n; i++)
    table[nums[i]]++;

  priority_queue<pair<int, int>> max_heap;
  for (auto [num, freq] : table)
    max_heap.push(pair(freq, num));

  for(int i = 0; i < k; i++)
  {
    res.push_back(max_heap.top().second);
    max_heap.pop();
  }
  
  return res;
}

int main()
{
  vector<int> nums = {1, 1, 1, 2, 2, 3333};

  nums = topKFrequent(nums, 2);

  for (int i = 0; i < nums.size(); i++)
    std::cout << nums[i] << ' ';

  return 0;
}
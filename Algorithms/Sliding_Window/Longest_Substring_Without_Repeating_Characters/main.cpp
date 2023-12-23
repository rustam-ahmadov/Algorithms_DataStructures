#include <iostream>
#include <unordered_map>

int lengthOfLongestSubstring(std::string s)
{
  const int n = s.size();
  if (!n)
    return 0;

  std::unordered_map<char, int> table;
  int i = 0, j = 0, max = 0, win_size = 0;
  while (j < n)
  {
    table[s[j]]++;
    j++;
    win_size = j - i;

    if (table.size() == win_size)
      max = max < win_size ? win_size : max;

    while (table.size() < win_size)
    {      
      table[s[i]]--;
      if (table[s[i]] == 0)
        table.erase(s[i]);

      i++;
      win_size--;
    }
  }
  return max;
}

int main()
{
  std::string s = {"pwwkew"};
  std::cout << lengthOfLongestSubstring(s);

  return 0;
}
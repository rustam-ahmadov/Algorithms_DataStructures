#include <iostream>
#include <unordered_map>

using namespace std;

int maxSym(int table[])
{
  int max = 0;
  const int n = 26;
  for (int i = 0; i < n; i++)
    if (max < table[i])
      max = table[i];

  return max;
}

int characterReplacement(std::string s, int k)
{
  const int n = s.size();
  int i = 0, max = 0, max_s = 0;
  int table[26]{0};
  for (int j = 0; j < n; j++)
  {
    table[s[j] - 'A']++;
    max_s = maxSym(table);
    if (j - i + 1 - max_s > k)
    {
      table[s[i] - 'A']--;
      i++;
    }
    max = max < j - i + 1 ? j - i + 1 : max;
  }
  return max;
}

int main()
{

  std::string s = {"RKOPRZWNXMUMEVBNJYK"};
  std::cout << characterReplacement(s, 10);
}
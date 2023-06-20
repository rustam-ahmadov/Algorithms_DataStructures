#include <iostream>
#include <unordered_map>
#include <string.h>

bool isAnagram(std::string s, std::string t)
{
    const int n = s.size();
    if (n != t.size())
        return false;

    std::unordered_map<char, int> table;

    for (int i = 0; i < n; i++)
    {
        table[s[i]]++;
    }

    for (int i = 0; i < n; i++)
    {
        if (table[t[i]] == 0)
            return false;
        else
            table[t[i]]--;
    }
    return true;
}

int main()
{

    std::string s = {"aacc"};
    std::string t = {"ccac"};

    bool is_anagram = isAnagram(s, t);
    std::cout << is_anagram;

    return 0;
}
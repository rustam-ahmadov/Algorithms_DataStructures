#include <iostream>
#include <string.h>
#include <vector>
#include <algorithm>
#include <unordered_map>

using namespace std;

void print(vector<vector<string>> &strs)
{
    std::cout << '[';
    for (int i = 0; i < strs.size(); i++)
    {
        std::cout << '[';
        
        for (int j = 0; j < strs[i].size(); j++)
        {
            std::cout << strs[i][j];
        }
        std::cout << ']';
        if(i < strs.size() - 1)
            std::cout<<',';
    }
    std::cout << ']';
}

vector<vector<string>> groupAnagrams(vector<string> &strs)
{
    const int n = strs.size();
    vector<vector<string>> result;
    if (!n)
        return result;

    unordered_map<string, vector<string>> table;
    for (int i = 0; i < n; i++)
    {
        string sorted = strs[i];
        sort(sorted.begin(), sorted.end());
        table[sorted].push_back(strs[i]);
    }

    for (auto kv : table)
        result.push_back(kv.second);

    return result;
}

int main()
{

    vector<string> strs1{"abc", "cba", "aab", "aac", "cab", "aba", "aca", "qwe", "wett", "azq"};
    vector<string> strs2{"eat", "tea", "tan", "ate", "nat", "bat"};
    vector<string> strs3{""};
    vector<string> strs4{"a"};

    vector<vector<string>> result = groupAnagrams(strs2);
    print(result);

    return 0;
}
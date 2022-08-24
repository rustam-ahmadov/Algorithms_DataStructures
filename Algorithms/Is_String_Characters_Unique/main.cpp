#include <iostream>
#include <unordered_map>

// O(n) by using hashtable
bool isUnique(std::string &str)
{
    std::unordered_map<char, int> map;
    const int n = str.size();

    for (int i = 0; i < n; i++)
    {
        if (map[str[i]] == NULL)
            map[str[i]] = '1';
        else
            return false;
    }
    return true;
}

// brute force O(n^2) 
bool isUnique(std::string &str, bool brute_force)
{
    const int n = str.size();
    for (int i = 0; i < n; i++)
        for (int j = i + 1; j < n; j++)
            if (str[i] == str[j])
                return false;

    return true;
}

// O(n) CCI p198
bool isUniqueChars(std::string &str)
{
    int n = str.size();
    if (n > 128)
        return false;

    bool *char_set = new bool[128];
    for (int i = 0; i < n; i++)
    {
        int val = str[i];
        if (char_set[val])
            return false;
        char_set[val] = true;
    }
    return true;
}

// O(n) CCI p193
bool isUniqueChars(std::string &str, bool with_bit)
{
    int n = str.size();
    int checker = 0;

    for (int i = 0; i < n; i++)
    {
        int val = str[i] - 'a';   

        if ((checker & (1 << val)) > 0)
            return false;

        checker = checker | (1 << val);
    }
}

int main(int argc, char **argv)
{

    std::string str = "ruar";
    std::cout << isUniqueChars(str);

    int num = 3;
    num << 1;
    std::cout << num;

    return 0;
}
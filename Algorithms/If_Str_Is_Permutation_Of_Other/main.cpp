#include <iostream>
#include <bits/stdc++.h>
#include <unordered_map>

// O(n log n) sort then check
bool isPermutation(std::string &s1, std::string &s2)
{
    const int s1_size = s1.size();
    const int s2_size = s2.size();

    if (s1_size != s2_size)
        return false;

    // O(n log n)
    sort(s1.begin(), s1.end());
    sort(s2.begin(), s2.end());

    // O(n)
    for (int i = 0; i < s1.size(); i++)
        if (s1[i] != s2[i])
            return false;

    return true;
}

// O(n)
bool isPermutation(std::string &s1, std::string &s2, bool array)
{
    const int s1_size = s1.size();
    const int s2_size = s2.size();

    if (s1_size != s2_size)
        return false;

    int letters[128];
    for(int i = 0; i < 128; i++)
        letters[i] = 0;

    for (int i = 0; i < s1_size; i++)
        letters[s1[i]]++;
    for (int i = 0; i < s2_size; i++)
    {
        letters[s2[i]]--;

        if (letters[s2[i]] == -1)
            return false;
    }
    return true;
}

int main(int argc, char **argv)
{
    // Given two strings, write a method to decide if one is a permutation of the
    //  other.

    std::string s1 = "rustam";
    std::string s2 = "tamsua";

    std::cout << isPermutation(s1, s2, true);

    return 0;
}

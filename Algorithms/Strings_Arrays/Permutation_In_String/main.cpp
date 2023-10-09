#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

// O(26 * n)
bool checkInclusion(string s1, string s2)
{
    unordered_map<char, int> mp, window;
    int n1 = s1.size(), n2 = s2.size();

    if (n1 > n2)
        return false;
    for (int i = 0; i < n1; i++)
    {
        mp[s1[i]]++;
        window[s2[i]]++;
    }

    if (mp == window)
        return true;

    for (int i = 0, j = n1; j < n2; j++)
    {
        window[s2[i]]--;
        if (!window[s2[i]])
            window.erase(s2[i]);
        window[s2[j]]++;

        if (mp == window)
            return true;
        i++;
    }
    return false;
}

// O(n)
bool checkInclusionN(string s1, string s2)
{
    const int n1 = s1.size();
    const int n2 = s2.size();
    if (n1 > n2)
        return false;

    const int CHAR_COUNT = 26;
    vector<int> s1Count(CHAR_COUNT, 0);
    vector<int> s2Count(CHAR_COUNT, 0);
    for (int i = 0; i < n1; ++i)
    {
        s1Count[s1[i] - 'a']++; // abc       -- 0, 1, 2, - - - - - - - - - - - - - - - - -
        s2Count[s2[i] - 'a']++; // bdawbacq  -- 0, 1, - ,3
    }

    // compare two arrays
    int matches = 0;
    for (int i = 0; i < CHAR_COUNT; ++i)
        if (s1Count[i] == s2Count[i])
            matches++;

    int l = 0;
    for (int r = n1; r < n2; ++r)
    {
        if (matches == CHAR_COUNT)
            return true;

        int i = s2[r] - 'a';  // bdawbacq
        s2Count[i]++;

        if (s2Count[i] == s1Count[i]) // if left is equal now
            matches++;
        else if (s2Count[i] - 1 == s1Count[i]) // if left was equal before
            matches--;

        i = s2[l] - 'a';
        s2Count[i]--;

        if (s2Count[i] == s1Count[i]) // if right is equal now 
            matches++;
        else if (s2Count[i] + 1 == s1Count[i]) // if right was equal now 
            matches--;

        l++;
    }

    return matches == 26;
}

int main()
{
    string s1 = "abc", s2 = "bdawbacq";
    std::cout << checkInclusionN(s1, s2);

    return 0;
}

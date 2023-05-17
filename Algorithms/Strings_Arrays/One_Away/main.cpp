#include <iostream>
#include <string.h>
// for abs();
#include <cstdlib>

bool replacement(std::string &s1, std::string &s2);
bool insertionOrRemoval(std::string &s1, std::string &s2);

// O(n)
bool oneEditAway(std::string &s1, std::string &s2)
{
    const int N1 = s1.size();
    const int N2 = s2.size();

    // if difference in length is greater than 1
    if (N1 != N2 && (N1 - N2 > 1 || N2 - N1 > 1))
        return false;

    if (N1 == N2)
        return replacement(s1, s2);
    else
        return insertionOrRemoval(s1, s2);
}

// O(n)
bool insertionOrRemoval(std::string &s1, std::string &s2)
{
    const int N1 = s1.size();
    bool one_away = false;

    for (int i = 0, j = 0; i < N1; i++)
    {
        if (s1[i] == s2[j])
            j++;
        else
        {
            if (one_away)
                return false;
            one_away = true;
        }
    }
    return one_away;
}

// O(n)
bool replacement(std::string &s1, std::string &s2)
{
    const int N = s1.size();
    bool one_away = false;

    for (int i = 0; i < N; i++)
    {
        if (s1[i] != s2[i])
        {
            if (one_away)
                return false;
            one_away = true;
        }
    }
    return one_away;
}

// We can put this two algorithms together
// CCI p200
bool oneEditAway(std::string &s1, std::string &s2, bool one_edit_away)
{
    const int N1 = s1.size();
    const int N2 = s2.size();

    if (abs(N1 - N2) > 1)
        return false;

    // Get longer and shorter strings
    std::string l = N1 > N2 ? s1 : s2;
    std::string s = N1 < N2 ? s1 : s2;

    int i = 0, j = 0;
    bool one_away = false;

    while (i < N1 && j < N2)
    {
        if (s[i] != s[j])
        {
            if (one_away)
                return false;
            one_away = true;

            if (N1 == N2)
                i++;
        }
        else
            i++;
        j++;
    }
}

int main(int argc, char **argv)
{
    std::string s1 = {"rustam"};
    std::string s2 = {"rustlm"};

    std::cout << oneEditAway(s1, s2, true);
    return 0;
}
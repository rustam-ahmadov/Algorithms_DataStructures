#include <iostream>
#include <stack>
#include <vector>
#include <bits/stdc++.h>
#include <algorithm>

// O(n)
std::string removeNonAlphanumeric(std::string &str);

// O(n)
void allCharactersToLowerCase(std::string &str);

// O(n)
bool palindrom(const std::string &str);

// Method 1 //Brute Force
// O(n^3)
int longestPalSubstr(const std::string &str);

// Method 2: Dynamic Programming.
// Approach: The time complexity can be reduced by storing results of sub-problems.
// O(n^2)
int longestPalSubstr_dynamic(const std::string &str);

// O(n)
void printSubStr(const std::string &str, const int &low, const int &high);

int main()
{
    longestPalSubstr_dynamic("1218765678");

    std::cout << std::endl;
    std::cout << std::endl;
    std::cout << std::endl;

    std::cout << std::endl;
    return 0;
}

std::string removeNonAlphanumeric(std::string &str)
{
    int n = str.size();

    std::string newStr;

    for (int i = 0; i < n; i++)
    {
        if (str[i] >= 97 && str[i] <= 122)
            newStr.push_back(str[i]);
        else if (str[i] >= 48 && str[i] <= 57)
            newStr.push_back(str[i]);
    }
    return newStr;
}

void allCharactersToLowerCase(std::string &str)
{
    int n = str.size();
    for (int i = 0; i < n; i++)
    {
        if (str[i] >= 65 && str[i] <= 90)
        {
            str[i] = str[i] + 32;
        }
    }
}

bool palindrom(const std::string &str)
{
    int n = str.size();

    if (n == 0)
        return false;
    if (n == 1)
        return true;

    int left = 0, right = 0;
    if (n % 2 == 0)
    {
        right = n / 2;
        left = n / 2 - 1;
    }
    else
        left = n / 2 - 1, right = n / 2 + 1;

    int pal_substrings_lenght = 0;
    while (right < n && left >= 0)
    {
        if (str[left] != str[right])
            return false;
        left--;
        right++;
    }
    return true;
}

void printSubStr(const std::string &str, const int &low, const int &high)
{
    for (int i = low; i <= high; i++)
    {
        std::cout << str[i];
    }
    std::cout << '\n';
}

int longestPalSubstr(const std::string &str)
{
    int n = str.size();
    int maxLenght = 1, start = 0;
    bool isPalindrom = true;

    for (int i = 0; i < n; i++)
    {
        for (int j = i + 1; j < n; j++)
        {
            isPalindrom = true;

            for (int k = 0; k < (j - i + 1) / 2; k++)
            {
                if (str[i + k] != str[j - k])
                {
                    isPalindrom = false;
                    break;
                }
            }

            if (isPalindrom && (j - i + 1) > maxLenght)
            {
                start = i;
                maxLenght = j - i + 1;
            }
        }
    }

    std::cout << "Longest palindrome substring is:";
    printSubStr(str, start, start + maxLenght - 1);

    return maxLenght;
}

int longestPalSubstr_dynamic(const std::string &str)
{
    int n = str.size();

    // table[i][j] will be false if substring
    // str[i..j] is not palindrome.
    // Else table[i][j] will be true
    bool table[n][n];
    // fill a table with zeros
    memset(table, false, sizeof(table));

    int max_lenght = 1;

    //'1 0 0 0'
    //'0 1 0 0'
    //'0 0 1 0'
    //'0 0 0 1'
    for (int i = 0; i < n; i++)
        table[i][i] = true;

    // check for sub-string of lenght 2.
    int start = 0;
    for (int i = 0; i < n - 1; i++)
    {
        if (str[i] == str[i + 1])
        {
            table[i][i + 1] = true;
            start = i;
            max_lenght = 2;
        }
    }

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
            std::cout << table[i][j] << ' ';
        std::cout << '\n';
    }

    // check for lenghts greater than 2.
    // k is lenght of substring
    for (int k = 3; k <= n; k++)
    {
        int sub_str_lenght = n - k + 1;
        for (int i = 0; i < sub_str_lenght; i++)
        {
            int j = i + k - 1;

            if (table[i + 1][j - 1] == true && str[i] == str[j])
            {
                table[i][j] = true;

                if (k > max_lenght)
                {
                    start = i;
                    max_lenght = k;
                }
            }
        }
    }

    printSubStr(str, start, start + max_lenght - 1);
    return 0;
}
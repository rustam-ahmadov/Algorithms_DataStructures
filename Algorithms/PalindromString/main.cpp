#include<iostream>


//O(n)
int isPalindrom(const std::string &str)
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

int main(int argc, int **argv)
{

    return 0;
}
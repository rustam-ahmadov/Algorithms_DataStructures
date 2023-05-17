#include <iostream>


//O(n) //pg 90 CCI
std::string whiteSpaceReplace(std::string &str, int true_length)
{
    const int n = str.size();
    int index = n;

    for (int i = true_length - 1; i >= 0; i--)
    {
        if (str[i] == ' ')
        {
            str[index - 1] = '0';
            str[index - 2] = '2';
            str[index - 3] = '%';
            index -= 3;
        }
        else
        {
            str[index - 1] = str[i];
            index--;
        }
    }
    return str;
}

int main(int argc, char **argv)
{
    std::string s = "Mr John Smith    ";
    std::cout << whiteSpaceReplace(s, 13);

    return 0;
}
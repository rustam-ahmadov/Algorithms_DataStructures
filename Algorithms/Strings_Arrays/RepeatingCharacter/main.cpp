#include <iostream>
#include <string.h>
#include <stack>

// O(n^2) time complexity
std::string getRepeatingCharacters(const std::string &str)
{
    std::string rep_characters;
    std::stack<char> s;
    int str_lenght = str.length();
    int rep_char_index = 0;

    for (int i = 1; i < str_lenght; i++)
    {
        for (int j = i - 1; j >= 0; j--)
        {
            if ((str[i] == str[j] )&& (rep_characters.find(str[i]) == -1) )
            {
                int a = rep_characters.find(str[i]);
                rep_characters.push_back(str[i]);
                break;
            }
        }
    }
    return rep_characters;
}
int main(int argv, char **args)
{
    std::string rep_char = getRepeatingCharacters("heloleaaa");
    std::cout << rep_char;
    return 0;
}
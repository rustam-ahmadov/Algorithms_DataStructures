#include <iostream>
#include <string.h>
#include <vector>


//prototypes
int getCharNum(char);
int *getFrequencyTable(int *table, int table_size, std::string &str);
bool chechMaxOneOdd(int *table, int table_size, std::string &str);


bool isPermutationOfPalindrome(std::string &str)
{
    int *table = new int[32];
    table = getFrequencyTable(table, 32, str);
    bool is_permutation_of_palindrome = chechMaxOneOdd(table, 32, str);
    delete table;
    return is_permutation_of_palindrome;
}


int *getFrequencyTable(int *table, int table_size, std::string &str)
{
    for (int i = 0; i < table_size; i++)
        table[i] = 0;

    int n = str.size();
    int character_val = 0;

    for (int i = 0; i < n; i++)
    {
        // if it's not space
        character_val = getCharNum(str[i]);
        if (character_val != -1)
            table[character_val]++;
    }
    return table;
}

bool chechMaxOneOdd(int *table, int table_size, std::string &str)
{
    bool foundOdd = false;
    for (int i = 0; i < table_size; i++)
    {
        // if there is an Odd character
        if (table[i] % 2)
        {
            if (foundOdd)
                return false;

            foundOdd = true;
        }
    }
    return true;
}

int getCharNum(char c)
{
    int val = c;

    // from 'a' to 'z'
    if (97 <= val && val <= 122)
        return val - 97;

    // from 'A' to 'Z' return 'a' - 'z'
    if (65 <= val && val <= 90)
        return val - 65;

    return -1;
}

int main(int argc, char **argv)
{

    std::string s = "Rusik i kisur";
    std::cout << isPermutationOfPalindrome(s);

    return 0;
}
#include <iostream>
using namespace std;

//O(n)
void printLongestCommonPrefix(string &s1, string &s2, string &s3)
{
    std::string common_prefix = "";
    int i = 0;
    while (s1[i] != '\0' && s2[i] != '\0' && s3[i] != '\0')
    {
        if (s1[i] != s2[i])
            break;
        common_prefix.push_back(s1[i]);
        i++;
    }
    i = 0;
    while (common_prefix[i] != 0 && s3[i] != 0)
    {
        if (common_prefix[i] == s3[i])
            i++;
        else
            common_prefix.pop_back();
    }
    std::cout << common_prefix;
}

int main(int argc, char **argv)
{
    std::string str1 = "abcdRustam";
    std::string str2 = "abcdamil";
    std::string str3 = "abcmiko";

    printLongestCommonPrefix(str1, str2, str3);

    return 0;
}
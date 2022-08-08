#include <iostream>

std::string commonPrefixUtil(std::string &s1, std::string &s2)
{
    int s_length = std::min(s1.size(), s2.size());
    std::string common_prefix = "";

    for (int i = 0; i < s_length; i++)
    {
        if (s1[i] != s2[i])
            break;
        common_prefix.push_back(s1[i]);
    }
    return common_prefix;
}

std::string commonPrefix(std::string arr[], int n)
{
    std::string prefix = arr[0];

    for (int i = 1; i < n; i++)
        prefix = commonPrefixUtil(prefix, arr[i]);

    return (prefix);
}
int main(int argc, char **argv)
{

    std::string s1 = "123Rustam";
    std::string s2 = "123Ruslan";
    std::string s3 = "123Namiq";

    std::string strs[3] = {s1, s2, s3};

    std::string common_prefix = commonPrefix(strs, 3);
    std::cout << common_prefix;

    return 0;
}
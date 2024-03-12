#include <iostream>

std::string commonPrefix(std::string &s1, std::string &s2)
{
    int s1_size = s1.size(), s2_size = s2.size();
    std::string com_pref = "";
    int n = s1_size <= s2_size ? s1_size : s2_size;

    if (!n)
        return com_pref;

    for (int i = 0; i < n; i++)
    {
        if (s1[i] != s2[i])
            break;
        com_pref.push_back(s1[i]);
    }
    return com_pref;
}

std::string getLongestCommonPrefix(std::string *str_arr, int n)
{
    std::string longest_com_pref = " ";
    if(n<2)
        return longest_com_pref;

    longest_com_pref = commonPrefix(str_arr[0],str_arr[1]);
    if(n==2)
        return longest_com_pref;
    
    for(int i = 2; i < n; i++)
    {
        longest_com_pref = commonPrefix(longest_com_pref,str_arr[i]);
    }
    return longest_com_pref;
}
int main(int argc, char **argv)
{

    std::string *str_arr = new std::string[2]; 
    str_arr[0] = "123Rustam";
    str_arr[1] = "123Ruslan";
    str_arr[2] = "123Namiq";

    std::cout<<getLongestCommonPrefix(str_arr,3);

    return 0;
}
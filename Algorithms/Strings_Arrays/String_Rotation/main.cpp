#include <iostream>
#include <string.h>

bool isRotation(std::string &s1, std::string &s2)
{
    bool is_rotation = false;

    const int N1 = s1.size();
    const int N2 = s2.size();

    if (N1 != N2)
        return is_rotation;

    std::string s1s1 = s1 + s1;

    const int N3 = s1s1.size();

    for (int i = 0, j = 0; i < N3, j < N2; i++, j++)
    {
        //?
    }
    return true;
}

int main(int argc, char **argv)
{
    std::string s1 = "rustam";
    std::string s2 = "tamrus";

    std::cout << isRotation(s1s1, s2);

    return 0;
}
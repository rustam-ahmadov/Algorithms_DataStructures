#include <iostream>
#include <string.h>

// O(n)
std::string getCompressed(const std::string &s)
{
    const int N = s.size();
    std::string compressed = "";

    int rep_size = 0;
    for (int i = 0; i < N; i++)
    {
        rep_size++;

        if (s[i] != s[i + 1] || i == N - 1)
        {
            compressed.push_back(s[i]);
            compressed.push_back(rep_size + 48);
            rep_size = 1;
        }
    }

    const int COMPRESSED_SIZE = compressed.size();
    if (COMPRESSED_SIZE < N)
        return compressed;
    return s;
}

int main(int argc, char **argv)
{
    std::cout << getCompressed("hhhhhhhhheello");
    return 0;
}
#include <iostream>

int reverse(int n)
{
    int res = 0;
    while (n)
    {
        res = res * 10 + n % 10;
        n /= 10;
    }
    return res;
}


int main()
{
    std::cout<< reverse(1)<<'\n';
    std::cout<< reverse(2)<<'\n';;
    std::cout<< reverse(32)<<'\n';;
    std::cout<< reverse(456)<<'\n';;
    return 0;
}
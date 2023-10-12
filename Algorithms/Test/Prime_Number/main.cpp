#include <iostream>
#include <cmath>

bool prime(int n)
{
    if (n == 0 || n == 1)
        return true;

    int root = sqrt(n);
    for (int i = 2; i <= root; i++)
    {
        int l = n % i;
        if (!l)
            return false;
    }

    return true;
}

int main()
{

    std::cout << prime(0);
    std::cout << '\n';
    std::cout << prime(1);
    std::cout << '\n';
    std::cout << prime(2);
    std::cout << '\n';
    std::cout << prime(3);
    std::cout << '\n';
    std::cout << prime(19);
    std::cout << '\n';
    std::cout << prime(41);
    std::cout << '\n';
    std::cout << prime(105);
}
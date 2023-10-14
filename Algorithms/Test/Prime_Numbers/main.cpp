#include <iostream>
#include <cmath>

void primeNumbers(int n)
{
    if (n == 1)
        return;

    bool isPrime = true;
    for (int i = 2; i <= n; i++)
    {
        int root = sqrt(i);
        for (int j = 2; j <= root; j++)
        {
            if (!(i % j))
            {
                isPrime = false;
                break;
            }
        }
        if (isPrime)
            std::cout << i << ' ';
        isPrime = true;
    }
}

int main()
{
    primeNumbers(100);
    return 0;
}
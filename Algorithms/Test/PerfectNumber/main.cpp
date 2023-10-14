#include <iostream>

bool isPerfect(int n)
{
    int sum = 1;
    for (int i = 2; i < n; i++)
    {
        int l = n % i;
        if (!l)
            sum += i;
    }
    if (sum == n)
        return true;
    return false;
}

void printPerfectNumbers(int n)
{
    if (n < 2)
        return;
    for (int i = 2; i <= n; i++)
        if (isPerfect(i))
            std::cout << i << ' ';
}

int main()
{

    printPerfectNumbers(1000);

    return 0;
}
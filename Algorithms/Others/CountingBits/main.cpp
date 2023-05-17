#include <iostream>
#include <vector>

// O(log n)
int getCountOfOne(int num)
{
    int ones_sum = 0;
    while (num)
    {
        ones_sum += num % 2;
        num /= 2;
    }
    return ones_sum;
}

// O(n)
std::vector<int> countBits(int num)
{
    std::vector<int> vector;

    for (int i = 0; i <= num; i++)
    {
        vector.push_back(getCountOfOne(i));
    }
    return vector;
}

int main(int argc, char **argv)
{
    std::vector<int> vec = countBits(5);
    int n = vec.size();

    for (int i = 0; i < n; i++)
    {
        std::cout << vec[i];
    }

    return 0;
}
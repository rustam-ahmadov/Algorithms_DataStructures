#include <iostream>
#include <vector>

void print(std::vector<int> nums)
{
    int n = nums.size();
    for (int i = 0; i < n; i++)
        std::cout << nums[i] << ' ';
    std::cout << '\n';
}

// O(n^2) Brute force
std::vector<int> productExceptSelf(std::vector<int> &nums)
{
    const int n = nums.size();
    std::vector<int> products;

    if (n == 0)
        return products;

    for (int i = 0; i < n; i++)
    {
        int product = 1;
        for (int j = 0; j < n; j++)
            if (j != i)
                product *= nums[j];

        products.push_back(product);
    }
    return products;
}

std::vector<int> productExceptSelfN(std::vector<int> &nums)
{
    const int n = nums.size();
    std::vector<int> products(n);

    // prefix
    int product = 1;
    for (int i = 0; i < n; i++)
    {
        products[i] = product;
        product *= nums[i];
    }

    product = 1;
    for (int i = n - 1; i >= 0; i--)
    {
        products[i] *= product;
        product *= nums[i];
    }
    return products;
}

int main()
{
    std::vector<int> nums = {2, 3, 4, 5};
    std::vector<int> products = productExceptSelfN(nums);
    print(products);

    return 0;
}
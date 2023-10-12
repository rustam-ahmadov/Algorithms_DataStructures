#include <iostream>
#include <vector>

using namespace std;

// Brute force
int maxAreaON2(vector<int> &height)
{
    const int n = height.size();
    int res = 0;
    int temp = 0;
    int min = 0;
    for (int i = 0; i < n - 1; i++)
    {
        min = height[i];
        for (int j = i + 1; j < n; j++)
        {
            if (height[j] < height[i])
                min = height[j];
            else
                min = height[i];

            temp = min * (j - i);
            if (res < temp)
                res = temp;
        }
    }
    return res;
}

int maxArea(vector<int> &height)
{
    const int n = height.size();
    int l = 0, r = n - 1;
    int l_n = 0, r_n = 0;
    int min = 0, res = 0, water = 0, i = 0;
    while (l != r)
    {
        l_n = height[l];
        r_n = height[r];

        if (l_n < r_n)
        {
            min = l_n;
            l++;
        }
        else
        {
            min = r_n;
            r--;
        }

        i++;
        water = min * (n - i);
        if(water > res)
            res = water;
    }
    return res;
}

int main()
{
    vector<int> height = {1,8,6,2,5,4,8,3,7};

    std::cout << maxArea(height);
    return 0;
}
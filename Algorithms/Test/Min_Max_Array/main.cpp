#include <iostream>

void printMinMax(int *arr, int n)
{
    int min = arr[0];
    int max = arr[0];

    for (int i = 0; i < n; i++)
    {
        if (max < arr[i])
            max = arr[i];
        if (min > arr[i])
            min = arr[i];
    }
    std::cout << "Max: " << max << std::endl;
    std::cout << "Min: " << min << std::endl;
}

int main()
{

    int *arr = new int[10]{3, 2, 5, 1, 63, 25, 76, 33, -1, 6};
    printMinMax(arr, 10);

    return 0;
}
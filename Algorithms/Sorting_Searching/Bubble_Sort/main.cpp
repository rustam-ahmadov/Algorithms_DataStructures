#include <iostream>

void print(int *arr, int n)
{
    std::cout << "Arr: ";
    for (int i = 0; i < n; i++)
        std::cout << arr[i] << ' ';
    std::cout << std::endl;
}

//best & worst O(n^2)
void bubble_sort(int *arr, int n)
{
    for (int i = 0; i < n - 1; i++)
    {
        for (int j = i + 1; j < n; j++)
            if (arr[i] > arr[j])
            {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        print(arr, n);
    }
}

int main()
{
    const int size = 7;
    int arr[size]{5, 3, 2, 10, 0, 65, 1};

    bubble_sort(arr, size);
}
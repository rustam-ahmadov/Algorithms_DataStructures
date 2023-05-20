#include <iostream>

void print(int *arr, int n)
{
    std::cout << "Arr: ";

    for (int i = 0; i < n; i++)
        std::cout << arr[i] << ' ';
    std::cout << '\n';
}

//Worst and Best case O(n^2)
void selection_sort(int *arr, int n)
{
    int min_ind;
    for (int i = 0; i < n; i++)
    {
        min_ind = i;

        // Find the minimum element in
        // unsorted array
        for (int j = i + 1; j < n; j++)
        {
            if (arr[j] < arr[min_ind])
                min_ind = j;
        }

        // Swap the found minimum element
        // with the first element
        if (min_ind != i)
        {
            int temp = arr[min_ind];
            arr[min_ind] = arr[i];
            arr[i] = temp;
        }
    }
}

int main()
{

    const int size = 7;
    int arr[size]{5, 3, 6, 2, 1, 9, 8};
    selection_sort(arr, size);
    print(arr, size);
}
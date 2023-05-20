#include <iostream>



void print(int *arr, int size);


//best case O(n) worst average O(n^2)
void insertion_sort(int *arr, int n)
{
    int j = 0;
    int key = 0;
    for (int i = 1; i < n ; i++)
    {
        key = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > key)
        {
            arr[j + 1] = arr[j];
            j--;
        }
        arr[j + 1] = key;
        print(arr,n);
    }
}

void print(int *arr, int size)
{
    std::cout << "Arr: ";
    for (int i = 0; i < size; i++)
        std::cout << arr[i] << ' ';
    std::cout<<'\n';
}

int main()
{
    const int size = 7;
    int arr[size] = {4, 7, 6, 5, 10, 2, 1};
    insertion_sort(arr, size);
    //print(arr, size);
}
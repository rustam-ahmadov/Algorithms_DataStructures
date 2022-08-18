#include <iostream>

void merge(int *arr, int l, int m, int r);

void fill(const int arr_size, int *arr)
{
    for (int i = 0; i < arr_size; i++)
        arr[i] = rand() % 20 + 1;
}

void print(const int arr_size, int *arr)
{
    for (int i = 0; i < arr_size; i++)
        std::cout << arr[i] << ' ';
    std::cout << '\n';
}

void mergeSort(int *arr, int l, int r)
{
    if (l < r)
    {
        // finding mid element
        int m = l + (r - l) / 2;
        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);

        merge(arr, l, m, r);
    }
}

void merge(int *arr, int l, int m, int r)
{
    int i, j, k;
    int n1 = m - l + 1;
    int n2 = r - m;

    // create temp arrays
    int L[n1], R[n2];

    // copy data to temp arr
    for (i = 0; i < n1; i++)
        L[i] = arr[l + i];

    for (j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    // merge the temp arrays
    i = 0;
    j = 0;
    k = l;
    while (i < n1 && j < n2)
    {
        if (L[i] <= R[j])
        {
            arr[k] = L[i];
            i++;
        }
        else
        {
            arr[k] = R[j];
            j++;
        }
        k++;
    }
    // Copy the remaining elements of L[]
    while (i < n1)
    {
        arr[k] = L[i];
        i++;
        k++;
    }
    // Copy the remaining elements of R[]
    while (j < n2)
    {
        arr[k] = R[j];
        j++;
        k++;
    }
}

int main(int argc, char **argv)
{

    srand(time(NULL));

    const int SIZE = 8;
    int arr[SIZE]{7, 2, 0, 10, 5, 8, 6, 4};

    // fill(SIZE, arr);
    print(SIZE, arr);

    mergeSort(arr, 0, SIZE - 1);
    print(SIZE, arr);

    return 0;
}
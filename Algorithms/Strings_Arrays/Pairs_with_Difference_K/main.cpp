#include <iostream>
#include <vector>
#include <bits/stdc++.h>
#include <unordered_map>

// O(log n)
bool binarySearch(std::vector<int> &, int);

// O(n^2)
void printPairsWithDifference(std::vector<int> &vec, int difference)
{
    int n = vec.size();
    for (int i = 0; i < n; i++)
    {
        for (int j = i + 1; j < n; j++)
        {
            if ((vec[j] - vec[i]) > difference)
                break;

            if ((vec[j] - vec[i] == difference))
            {
                std::cout << vec[i] << ' ' << vec[j] << "    ";
                break;
            }
        }
    }
}

// O(n log n)
void printPairsWithDifferenceBinary(std::vector<int> &vec, int difference)
{
    int n = vec.size();
    for (int i = 0; i < n; i++)
    {
        if (binarySearch(vec, vec[i] + difference))
            std::cout << vec[i] << ' ' << vec[i] + difference << "   ";
    }
}

// O(log n)
bool binarySearch(std::vector<int> &vec, int value)
{
    int first_index = 0;
    int last_index = vec.size() - 1;
    int mid_index;

    while (first_index <= last_index)
    {
        mid_index = (first_index + last_index) / 2;

        if (vec[mid_index] == value)
            return true;
        if (vec[mid_index] > value)
            last_index = mid_index - 1;
        else
            first_index = mid_index + 1;
    }
    return false;
}

int main(int argc, char **argv)
{

    std::vector<int> vec = {1, 7, 5, 9, 2, 12, 3};

    // O(n log n)
    // sort(vec.begin(), vec.end());
    // after sorting we do O(n^2) what why O(n log n + n^2) ==> O(n^2)
    // printPairsWithDifference(vec, 2);

    // We now have a two - step algorithm, where both steps take O(N log N) time.
    // Now, sorting is the new bottleneck.
    // Optimizing the second step won't help because the first step is slowing us down anyway.
    // sort(vec.begin(), vec.end());
    // printPairsWithDifferenceBinary(vec, 2);

    // but we can use hashTable and then search in hashTable for pairs O(n);

    return 0;
}
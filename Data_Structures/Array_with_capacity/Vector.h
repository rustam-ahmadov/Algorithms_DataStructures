#pragma once
#include <iostream>

class Vector
{
    int size;
    int capacity;
    int *arr;

public:
    Vector()
    {
        size = capacity = 0;
        arr = nullptr;
    }

    int &operator[](int index)
    {
        return arr[index];
    }
    void push_back(const int &value)
    {

        // for first element
        if (capacity == 0 && size == 0)
        {
            capacity++;
            arr = new int[capacity];
            arr[capacity - 1] = value;
        }
        else if (size < capacity)
        {
            // O(1)
            arr[size] = value;
        }
        else if (size == capacity)
        {
            capacity *= 2;
            int *new_arr = new int[capacity];

            // O(n)
            // copy old arr and push_back new value
            for (int i = 0; i < size; i++)
                new_arr[i] = arr[i];
            new_arr[size] = value;

            delete[] arr;
            arr = new_arr;
        }
        size++;
    }

    int getSize()
    {
        return size;
    }
    int getCapacity()
    {
        return capacity;
    }
    void print()
    {
        for (int i = 0; i < size; i++)
            std::cout << arr[i] << ' ';
    }

    ~Vector()
    {
        delete[] arr;
    }
};
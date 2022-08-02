#pragma once
#include <iostream>

template <typename T>
class Queue
{
	T *arr;
	int size;

public:
	Queue() : size(0), arr(nullptr){};
	Queue(const Queue &other) : size(other.size)
	{
		arr = new T[size];
		for (size_t i = 0; i < size; i++)
		{
			arr[i] = other.arr[i];
		}
	}
	Queue &operator=(const Queue &other)
	{
		if (this == &other)
		{
			return *this;
		}
		else
		{
			delete[] arr;
			arr = new T[other.size];
			size = other.size;
			for (int i = 0; i < size; i++)
			{
				arr[i] = other.arr[i];
			}
			return *this;
		}
	}
	Queue(Queue &&other)
	{
		arr = other.arr;
		size = other.size;
		other.arr = nullptr;
		other.size = 0;
	};

	Queue &operator=(Queue &&other)
	{
		if (this == &other)
		{
			return *this;
		}
		delete[] arr;
		arr = other.arr;
		size = other.size;
		other.arr = nullptr;
		other.size = 0;
	};

	T &front()
	{
		return arr[0];
	}

	T &back()
	{
		return arr[size - 1];
	}

	void pop()
	{
		T *new_arr = new T[--size];
		for (int i = 0; i < size; i++)
		{
			new_arr[i] = arr[i + 1];
		}
		delete[] arr;
		arr = new_arr;
	};
	
	void push(const T &element)
	{
		T *n_arr = new T[++size];
		for (int i = 0; i < size - 1; i++)
		{
			n_arr[i] = arr[i];
		}
		delete[] arr;
		n_arr[size - 1] = element;
		arr = n_arr;
	};
	
	int rsize() { return size; };
	
	void swap(Queue &other)
	{
		T *temp = arr;
		arr = other.arr;
		other.arr = temp;
		int tempovik = size;
		size = other.size;
		other.size = tempovik;
	}
	
	bool empty()
	{
		if (!size)
		{
			return true;
		}
		return false;
	};
	template <typename T1>
	friend std::ostream &operator<<(std::ostream &os, const Queue<T1> &a);
	
	~Queue()
	{
		delete[] arr;
	}
};

template <typename T1>
std::ostream &operator<<(std::ostream &os, const Queue<T1> &a)
{
	for (int i = 0; i < a.size; i++)
	{
		os << a.arr[i] << ' ';
	}
	return os;
};
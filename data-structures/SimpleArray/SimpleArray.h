#pragma once
#include<iostream>
#include<limits>
#include<string>



class SimpleArray {
private:
	int size=0;
	int* m_data = nullptr;

	
public:		
	void pushBack(int element) {
		int* arr = new int[size+1];	
	
		for (int i = 0; i < size; i++)
		{
			arr[i] = m_data[i];
		}

		arr[size] = element;
		delete[]m_data;
		m_data = arr;
		size++;
	}

	void check() {		
		if (size==0)
		{
			std::cout << "Array is empty";
		}
		else
		{
			std::cout << "Array isn't empty";
			
		}
	}

	int getsize() {
		return size;
	}
	int get(int index) {		
		return m_data[index];
	}

	int *getAddr() {
		return m_data;
	}

	~IntArray()
	{ 
		delete[]m_data;
	}

	

	
};
#pragma once
#include<iostream>;

template<typename T>
class Stack {
	T* arr;
	int size;
public:
	//pochemu ya ne mogu napisat eti functions v otdelnom cpp ?
	Stack() :size(0), arr(nullptr) {};
	//Constructor of copy
	Stack(const Stack& other) :size(other.size) {
		arr = new T[size];
		for (size_t i = 0; i < size; i++)
		{
			arr[i] = other.arr[i];
		}
	}
	//operator prisvaivaniya
	Stack& operator=(const Stack& other) {
		if (this == &other)
		{
			return *this;
		}
		else
		{
			delete[]arr;
			arr = new T[other.size];
			size = other.size;
			for (int i = 0; i < size; i++)
			{
				arr[i] = other.arr[i];
			}
			return *this;
		}
	}
	//Constructor peremesheniya
	Stack(Stack&& other) {
		arr = other.arr;
		size = other.size;
		other.arr = nullptr;
		other.size = 0;
	};
	//Operator prisvaivaniya peremesheniyem // xota ne znayu tochno kak on naz)
	Stack& operator=(Stack&& other) {
		if (this == &other)
		{
			return *this;
		}
		delete[]arr;
		arr = other.arr;
		size = other.size;
		other.arr = nullptr;
		other.size = 0;
	};
	// - top() возвращает ссылку на элемнет из вершины стека to est last one 
	T top() {
		return arr[size - 1];
	}
	//- pop() удаляет элемент из вершины стека  to est last one 
	void pop() {
		//I tried it as you said
		size--;
	};
	//- push(element) добавляет элемент на вершину стека  to est last one 
	void push(const T& element) {
		T* n_arr = new T[++size];
		for (int i = 0; i < size - 1; i++)
		{
			n_arr[i] = arr[i];
		}
		delete[]arr;
		n_arr[size - 1] = element;
		arr = n_arr;
	};
	//-size() возвращает количество элементов в стеке
	int rSize() { return size; };
	//- swap(other_stack) обменивает содержимым два стека
	void swap(Stack& other) {
		T* temp = arr;
		arr = other.arr;
		other.arr = temp;
		int tempovik = size;
		size = other.size;
		other.size = tempovik;
	};
	//empty() возвращает true если стек пуст и false в противном случае
	bool empty() {
		if (!size)
			return true;

		return false;
	};

	//just for std::cout<<;
	template<typename T1>
	friend std::ostream& operator<<(std::ostream& os, const Stack<T1>& a);

	//  Kuda je bez nego)))
	~Stack()
	{
		delete[]arr;
	}
};

template<typename T1>
std::ostream& operator<<(std::ostream& os, const Stack<T1>& a) {
	for (int i = 0; i < a.size; i++)
	{
		os << a.arr[i] << ' ';
	}
	return os;
};




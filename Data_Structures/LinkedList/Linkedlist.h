#pragma once
#include<iostream>



template<typename T>
class Node {
	Node* next;
	T data;
	template<typename T>
	friend class List;
};

template<typename T>
class List {
	Node<T>* head;
	Node<T>* tail;
	int size;
public:
	//Default
	List();
	//Copy constructor
	List(const List<T>& other);
	//Move constructor
	List(List<T>&& other);

	//Assign operator
	List& operator=(const List<T>& other);
	//Assign operator for temporary object(rvalue)
	List&operator=(List<T>&& other);
	
	List operator+(const List<T>&);
	
	void push_back(const T& value);
	
	void insert_after(const T& value, int position);
	
	void pop_front();

	void pop_back();

	bool empty() { return head; }
	
	void swap(List& other);

	void print_list();
	
	void remove(int);

	int list_size() { return size; }

	~List();
};

template<typename T>
class HashContainer {
	int size;
	List<T>* arr;
public:
	HashContainer() { size = 0, arr = nullptr; }
	HashContainer operator[](const T& value) {

	}
};

template<typename T>
List<T>::List()
{
	size = 0;
	head = nullptr;
	tail = nullptr;
}

template<typename T>
List<T>::List(const List<T>& other)
{
	size = other.size;
	Node<T>* temp = new Node<T>;
	temp->data = other.head->data;
	temp->next = NULL;
	head = temp;
	tail = temp;
	Node<T>* tempOther = other.head;
	for (int i = 0; i < size - 1; i++)
	{
		Node<T>* temp = new Node<T>;
		tempOther = tempOther->next;
		temp->data = tempOther->data;
		tail->next = temp;
		tail = temp;
	}
	tail->next = NULL;
}

template<typename T>
List<T>::List(List<T>&& other)
{
	size = other.size;
	head = other.head;
	tail = other.tail;	
	other.tail = NULL;
	other.head = NULL;

}

template<typename T>
List<T>&::List<T>::operator=(const List<T>& other)
{
	while (head)
	{
		Node<T>* temp = head;
		head = head->next;
		delete temp;
	}
	size = other.size;
	Node<T>* tempOther = other.head;
	for (int i = 0; i < size; i++)
	{
		Node<T>* temp = new Node<T>;
		temp->data = tempOther->data;
		if (i == 0)
		{
			head = temp;
			tail = temp;
		}
		else
		{
			tail->next = temp;
			tail = temp;
		}
		tempOther = tempOther->next;
	}
	tail->next = NULL;
	return *this;
}

template<typename T>
List<T>& List<T>::operator=(List<T>&& other)
{
	while (head)
	{
		Node<T>* temp = head;
		head = head->next;
		delete temp;
	}
	size = other.size;
	size = other.size;
	head = other.head;
	tail = other.tail;
	other.tail = NULL;
	other.head = NULL;
	return *this;
}

template<typename T>
List<T> List<T>::operator+(const List<T>& other)
{
	List<T> temp;
	temp.size = other.size + size;

	//temp for other
	Node<T>* tempOther = other.head;
	//temp for current
	Node<T>* tempCurrent = head;
	for (int i = 0; i < temp.size; i++)
	{
		Node<T>* tempNode = new Node<T>;
		if (i < other.size)
		{
			tempNode->data = tempOther->data;
			tempOther = tempOther->next;
			if (i == 0)
			{
				temp.head = tempNode;
				temp.tail = tempNode;
			}
			else
			{
				temp.tail->next = tempNode;
				temp.tail = tempNode;
			}
		}
		else if(i>=other.size)
		{
			tempNode->data = tempCurrent->data;
			tempCurrent = tempCurrent->next;
			temp.tail->next = tempNode;
			temp.tail = tempNode;
		}		
	}
	temp.tail->next = NULL;
	return temp;
}

template<typename T>
void List<T>::push_back(const T& value)
{
	size++;
	Node<T>* temp = new Node<T>;
	temp->data = value;
	temp->next = NULL;
	if (head == NULL)
	{
		head = temp;
		tail = temp;
		temp = NULL;
	}
	else
	{
		tail->next = temp;
		tail = temp;
	}
}

template<typename T>
void List<T>::pop_front()
{
	Node<T>* temp = head->next;
	delete	head;
	head = temp;
	size--;
}

template<typename T>
inline void List<T>::swap(List& other)
{
	Node<T>* temp = other.head;
	other.head = head;
	head = temp;
	int sizetemp = other.size;
	other.size = size;
	size = sizetemp;
}

template<typename T>
void List<T>::print_list()
{
	Node<T>* temp = head;

	int a = 0;
	std::cout << "Element " << ++a << " adress-";
	std::cout << head << ':';
	while (temp != NULL)
	{
		if (temp->next == NULL)
		{
			std::cout << "Value-";
			std::cout << temp->data << " \n";
			std::cout << "Next adress-" << temp->next;
			break;
		}
		std::cout << "Value-";
		std::cout << temp->data << " ";
		std::cout << '\n';
		std::cout << "Element " << ++a << " adress-";
		std::cout << temp->next << ':';
		temp = temp->next;
	}
	std::cout << '\n';
	std::cout << '\n';
}

template<typename T>
inline void List<T>::insert_after(const T& value, int position)
{
	if (position <= size)
	{
		Node<T>* temp = head;
		int a = 0;
		if (a == position)
		{
			Node<T>* newNode = new Node<T>;
			newNode->data = value;
			head = newNode;
			head->next = temp;
		}
		else
		{
			while (a < position - 1)
			{
				a++;
				temp = temp->next;
			}
			Node<T>* newNode = new Node<T>;
			Node<T>* temp_add = temp->next;
			temp->next = newNode;
			newNode->data = value;
			newNode->next = temp_add;
		}
		size++;
	}
}

template<typename T>
inline void List<T>::remove(int position)
{
	if (position <= size)
	{
		if (position == 1)
		{
			Node<T>* temp = head;
			head = head->next;
			delete temp;
		}
		else
		{
			Node<T>* temp = head;
			Node<T>* temptemp = temp;
			for (int i = 0; i < position - 1; i++)
			{
				temptemp = temp;
				temp = temp->next;
			}
			temptemp->next = temp->next;
			delete temp;
		}
		size--;
	}
}


template<typename T>
void List<T>::pop_back()
{
	Node<T>* temp = head;
	while (temp->next != tail)
	{
		temp = temp->next;
	}
	delete tail;
	tail = temp;
	tail->next = NULL;
	size--;
}

template<typename T>
List<T>::~List()
{
	while (head)
	{
		Node<T>* temp = head;
		head = head->next;
		delete temp;
	}
	size = 0;
}

#pragma once
#include <iostream>
#include <functional>

// prototype
template <typename Tkey, typename Tvalue>
class LinkedList;
template <typename Tkey, typename Tvalue>
class HashTable;

template <typename Tkey, typename Tvalue>
class Element
{
	Tkey key;
	Tvalue value;
	Element<Tkey, Tvalue> *next;
	friend LinkedList<Tkey, Tvalue>;
	friend HashTable<Tkey, Tvalue>;

public:
	Element(const Tkey &key, const Tvalue &value)
	{
		this->key = key;
		this->value = value;
	}

	Element(const Tkey &key)
	{
		this->key = key;
	}

	Tkey &getKey()
	{
		return key;
	}

	Tvalue &getValue()
	{
		return value;
	}
};

template <typename Tkey, typename Tvalue>
class LinkedList
{
	Element<Tkey, Tvalue> *head;
	Element<Tkey, Tvalue> *last;
	int size;
	friend HashTable<Tkey, Tvalue>;

public:
	LinkedList() { head = NULL, last = NULL, size = 0; }
	void pushback(const Tkey &key)
	{
		Element<Tkey, Tvalue> *temp = new Element<Tkey, Tvalue>(key);
		temp->next = NULL;
		if (!head)
		{
			head = temp;
			last = temp;
		}
		else
		{
			last->next = temp;
			last = temp;
		}
		last->next = NULL;
		size++;
	}
	~LinkedList()
	{
		while (head)
		{
			Element<Tkey, Tvalue> *temp = head;
			head = head->next;
			delete temp;
		}
	}
};
template <typename Tkey, typename Tvalue>
class HashTable
{
	LinkedList<Tkey, Tvalue> *arr = nullptr;
	int bucket;

public:
	HashTable(const int &bucket)
	{
		this->bucket = bucket, arr = new LinkedList<Tkey, Tvalue>[bucket];
	};

	Tvalue &operator[](const char *key)
	{
		std::hash<std::string> hashFunction;
		int index = hashFunction(key) % bucket;
		Element<Tkey, Tvalue> *temp = arr[index].head;

		if (arr[index].head)
		{
			while (temp)
			{
				if (temp->key == key)
				{
					return temp->value;
				}
				temp = temp->next;
			}
			arr[index].pushback(key);
			return arr[index].last->value;
		}
		else
		{
			arr[index].pushback(key);
			return arr[index].head->value;
		}
	}
	void print()
	{
		for (int i = 0; i < bucket; i++)
		{
			if (arr[i].head)
			{
				Element<Tkey, Tvalue> *temp = arr[i].head;
				std::cout << i + 1 << ":";
                
				while (temp)
				{
					std::cout << temp->value << ' ';
					temp = temp->next;
				}
			}
			else
			{
				std::cout << i + 1 << ":";
			}
			std::cout << '\n';
		}
	}
	~HashTable()
	{
		delete[] arr;
	}
};
#include <iostream>
template <typename T>
class Stack
{
private:
    T *_data = nullptr;
    int _capacity = 1;
    int _size = 0;
    
public:
    Stack(){};
    void push(T data)
    {
        
        if (_size < _capacity - 1)
        {
            _data[_size] = data;
            _size++;
        }
        else
        {
            _capacity *= 2;
            T *temp = new T[_capacity];

            for (int i = 0; i < _size; i++)
                temp[i] = _data[i];

            temp[_size] = data;

            delete[] _data;
            _data = temp;

            _size++;
        }
    }

    void pop()
    {
        if (_size > 0){
            _size--;
            if (_size < _capacity / 2)
            {
                _capacity /= 2;
                T *temp = new T[_capacity];
                for (int i = 0; i < _capacity; i++)
                    temp[i] = _data[i];
                delete[] _data;
                _data = temp;
            }            
        }
    }

    T& pick(){
        return _data[_size-1];
    }

    void print()
    {
        for (int i = 0; i < _size; i++)
            std::cout << _data[i] << '\n';
    }

    int get_capacity()
    {
        return _capacity;
    }
};

#include <stack>
#include <iostream>

class MinStack
{
    std::stack<int> _min;
    std::stack<int> _stack;

public:
    MinStack() {}

    void push(int val)
    {
        if (getMin() > val || _min.empty())
            _min.push(val);

        _stack.push(val);
    }

    void pop()
    {
        if (_stack.top() == getMin())
            _min.pop();
        _stack.pop();
    }

    int top()
    {
        return _stack.top();
    }

    int getMin()
    {
        return _min.top();
    }
};
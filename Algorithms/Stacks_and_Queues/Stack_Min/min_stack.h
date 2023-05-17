#include <stack>
#include <iostream>

class Min_Stack
{
    std::stack<int> min_s;
    std::stack<int> s;

public:
    void push(int val)
    {
        s.push(val);
        if (min_s.empty() || min_s.top() >= val)
            min_s.push(val);
    }
    void pop()
    {
        if (s.top() == min_s.top())
            min_s.pop();
        s.pop();
    }
    int top()
    {
        return s.top();
    }

    int getMin()
    {
        return min_s.top();
    }
};
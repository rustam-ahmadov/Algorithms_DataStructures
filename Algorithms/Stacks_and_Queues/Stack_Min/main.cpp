#include "min_stack.h"
#include <iostream>

int main(int argc, char **argv)
{
    MinStack m;
    m.push(0);
    m.push(1);
    m.push(-10);
    m.push(-9);

    std::cout << m.getMin();
    std::cout<<'\n';
    m.pop();
    std::cout << m.getMin();
    std::cout<<'\n';
    m.pop();
    std::cout << m.getMin();
    std::cout<<'\n';
    return 0;
}
#include"min_stack.h"

int main(int argc, char **argv)
{
    Min_Stack m;
    m.push(4);
    m.push(5);
    m.push(2);
    m.push(3);
    m.push(1);
    m.push(0);
    m.push(-1);
    std::cout<<m.getMin();
    return 0;
}
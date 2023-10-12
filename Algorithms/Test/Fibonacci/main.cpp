#include <iostream>

int fibonacci(int n)
{
    if(n == 1 || n == 2)
        return 1;
    
    int f = 1, s = 1;
    int res = 0;
    for(int i = 3; i <= n; i++)
    {
        res = f + s;
        f = s;
        s = res;
    }
    return res;
}


int main(){

    std::cout<<fibonacci(1)<<std::endl;
    std::cout<<fibonacci(2)<<std::endl;
    std::cout<<fibonacci(3)<<std::endl;
    std::cout<<fibonacci(4)<<std::endl;
    std::cout<<fibonacci(5)<<std::endl;
    std::cout<<fibonacci(6)<<std::endl;
    std::cout<<fibonacci(7)<<std::endl;
    std::cout<<fibonacci(8)<<std::endl;
    std::cout<<fibonacci(9)<<std::endl;


    return 0;
}
#include<iostream>

void fizzBuzz(int n)
{
    std::cout<<n<<':';
    if( !(n % 5) && !(n % 3))
        std::cout<<"FizzBuzz";
    else if(!(n % 3))
        std::cout<<"Fizz";
    else if(!(n % 5))
        std::cout<<"Buzz";
    std::cout<<'\n';
}

void fizzBuzzTill(int n)
{
    for(int i = 3; i <= n; i++ )
        fizzBuzz(i);
}


int main(){

    fizzBuzzTill(100);


    return 0;
}
#include<iostream>

//O(n)
int getFibNum(int n)
{
    if(n == 1||n==2)
        return 1;

    int t1 = 1, t2 = 1, nextTerm = 0;

    for (int i = 2; i < n; ++i) {
        
        nextTerm = t1 + t2;
        t1 = t2;
        t2 = nextTerm;
    }
    return nextTerm;
}

int main(){

    std::cout<<getFibNum(6);


    return 0;
}
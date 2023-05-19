#include<iostream>


int count_of_consecutive_digit_one(int*bits,int N)
{
    int one_d_per_round = 0;
    int one_d_max = 0;
    for(int i = 0; i < N; i++)
    {
        if(bits[i] == 1){
            one_d_per_round++;
            if(one_d_per_round > one_d_max)
                one_d_max = one_d_per_round;
        }
        else
            one_d_per_round = 0;
    }
    return one_d_max;
}

int main(){


    int bits[25]{1,1,1,1,1,0,0,0,1,1,1,1,1,1,0,0,0,0,1,1,1,1,1,1,1};
    int one_d_max = count_of_consecutive_digit_one(bits, 25);
    std::cout<<one_d_max;


}
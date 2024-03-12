#include "Vector.h"

int main(int argc, char **argv)
{

    Vector vector;

    int num = 1;

    while (num <= 100)
    {
        vector.push_back(num);
        std::cout << vector[num - 1] << ' ';
        std::cout << "Capacity:" << vector.getCapacity() << ' ';
        std::cout << "Size:" << vector.getSize();
        std::cout<<'\n';
        num++;
    }

    return 0;
}
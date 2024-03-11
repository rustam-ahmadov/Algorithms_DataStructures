#include "BinarySearchTree.h"

int main(int argc, char **argv)
{
    // BinarySearchTree bst;
    // int result = bst.insert(100);
    // result = bst.insert(50);
    // result = bst.insert(150);
    // result = bst.insert(25);
    // result = bst.insert(36);
    // result = bst.insert(110);
    // result = bst.insert(105);
    // result = bst.insert(106);
    // //bst.display();
    // //result = bst.remove(100);
    // //bst.display();
    // bst.makeEmpty();
    // bst.display();

    int *a = new int{1};
    int *b = new int{2};

    std::cout<<"address of pointer a:"<<&a;
    std::cout<<'\n';
    std::cout<<"address of memory where a points:"<<&*a;
    std::cout<<'\n';
    std::cout<<"value:"<<*a;
    std::cout<<'\n';
    delete a;
    std::cout<<"address of memory where points a after delete a:"<<&*a;
    std::cout<<'\n';
    std::cout<<"value of a after delete:"<<*a;

    // std::cout<<'\n';

    // a = b;

    // std::cout<<a[0];
    // std::cout<<'\n';
    // std::cout<<b[0];

    return 0;
}
#include <iostream>
#include "../Linked_List/list.h"

void removeMidNode(List::node *head)
{
    if (head == nullptr)
        return;

    List::node *temp = head;
    int size = 0;
    List::node *cur = head;
    List::node *bef_cur = nullptr;

    //finding the size of list
    while (temp)
    {
        size++;
        temp = temp->next;
    }

    // put cur on the mid node of list
    int mid = size / 2;
    for (int i = 0; i < mid; i++)
    {
        bef_cur = cur;
        cur = cur->next;
    }

    //removing
    bef_cur->next = cur->next;
    delete cur;
}

int main(int argc, char **argv)
{
    List list;
    list.pushBack(1);
    list.pushBack(2);
    list.pushBack(3);
    list.pushBack(4);
    list.pushBack(5);
    list.pushBack(6);
    list.pushBack(7);
    list.pushBack(8);
    list.pushBack(9);

    removeMidNode(list.head);
    list.print();

    return 0;
}
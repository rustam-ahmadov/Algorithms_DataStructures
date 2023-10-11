#include <iostream>
#include "../Linked_List/list.h"

List::node *reverseList(List::node *head)
{
    List::node *prev = NULL, *next = NULL;
    while (head)
    {
        next = head->next;
        head->next = prev;
        prev = head;
        head = next;
    }
    return head;
}

int main()
{

    List *l = new List();
    l->pushBack(1);
    l->pushBack(2);
    l->pushBack(3);
    l->pushBack(4);
    l->pushBack(5);

    l->print();
    return 0;
}
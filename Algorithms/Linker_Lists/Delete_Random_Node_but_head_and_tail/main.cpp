#include <iostream>
#include "../Linked_List/list.h"

// O(1)
void deleteRandomNode(List::node *&node)
{
    List::node *temp = node;
    node = node->next;
    delete temp;
}

// O(1)
void deleteNode(List::node *node)
{
    if (node == nullptr)
        return;

    if (node->next == nullptr)
    {
        delete node;
        return;
    }
    List::node *temp = node->next;
    node->val = temp->val;
    node->next = temp->next;
    delete temp;
}

int main(int argc, char **argv)
{
    List list;
    list.pushBack(1);
    list.pushBack(2);
    list.pushBack(3);
    list.pushBack(4);
    deleteNode(list.head->next->next->next);
    list.print();

    return 0;
}
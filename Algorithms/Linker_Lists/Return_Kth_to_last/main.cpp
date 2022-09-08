#include "..\Linked_List/list.h"

// from k to the last element
void printFromKthToLast(List::node *head, int k)
{
    List::node *current = head;
    int element_num = 0;
    while (current)
    {
        element_num++;
        if (element_num == k)
            break;
        current = current->next;
    }

    while (current)
    {
        std::cout << current->val << ' ';
        current = current->next;
    }
}

int printKthToLast(List::node *head, int k)
{
    if (!head)
        return 0;
    int index = printKthToLast(head->next, k) + 1;
    if (index == k)
        std::cout << k << "th to last node is " << head->val;
    return index;
}
int main(int argc, char **argv)
{
    List l;
    l.pushBack(0);
    l.pushBack(1);
    l.pushBack(2);
    l.pushBack(3);
    l.pushBack(4);
    l.pushBack(5);
    l.pushBack(6);
    l.pushBack(7);
    l.pushBack(8);
    l.pushBack(9);
    l.print();

    printKthToLast(l.head, 4);

    return 0;
}

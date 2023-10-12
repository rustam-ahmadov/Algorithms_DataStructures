#include "../Linked_List/list.h"

//O(n)
List::node *partition(List::node *head, int x)
{
    if (!head)
        return head;

    List less;
    List more;

    while (head)
    {
        if (head->val >= x)
            more.pushBack(head->val);
        else
            less.pushBack(head->val);

        head = head->next;
    }

    List::node *temp = less.head;

    while (temp->next)
        temp = temp->next;
    temp->next = more.head;

    return less.head;
}

int main(int argc, char **argv)
{
    List l;
    l.pushBack(3);
    l.pushBack(5);
    l.pushBack(8);
    l.pushBack(5);
    l.pushBack(10);
    l.pushBack(2);
    l.pushBack(1);
    l.print();

    List partitioned;
    partitioned.head = partition(l.head, 5);
    partitioned.print();

    return 0;
}
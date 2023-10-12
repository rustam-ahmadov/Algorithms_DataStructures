#include "../Linked_List/list.h"

bool equal(List::node *n1, List::node *n2)
{
    return n1 == n2;
}

List::node *getLastNode(List::node *head, int &size)
{
    while (head->next)
    {
        head = head->next;
        size++;
    }
    return head;
}

List::node *getIntersectionNode(List::node *n1, List::node *n2)
{
    if (!n1 || !n2)
        return NULL;

    int n1_size = 1, n2_size = 1;
    if (!equal(getLastNode(n1, n1_size), getLastNode(n2, n2_size)))
        return NULL;

    if (n1_size > n2_size)
        while (n1_size > n2_size)
        {
            n1 = n1->next;
            n1_size--;
        }
    else if (n1_size < n2_size)
        while (n2_size > n1_size)
        {
            n2 = n2->next;
            n2_size--;
        }
    while (n1 && n2)
    {
        if (n1 == n2)
            return n1;
        n1 = n1->next;
        n2 = n2->next;
    }

    return NULL;
}

int main(int argc, char **argv)
{
    List l1;
    l1.pushBack(1);
    l1.pushBack(2);
    l1.pushBack(3);
    l1.pushBack(4);

    List l2;
    l2.head = new List::node;
    l2.head->val = 10;
    l2.head->next = l1.head;

    List::node *node = getIntersectionNode(l1.head, l2.head);
    std::cout << node << std::endl;

    return 0;
}
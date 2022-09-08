#include "../Linked_List/list.h"

// O(n)
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
// O(n)
int printKthToLast(List::node *head, int k)
{
    if (!head)
        return 1;
    int index = printKthToLast(head->next, k);
    if (index == k)
        std::cout << k << "th to last node is " << head->val;
    return ++index;
}

List::node *nthToLast(List::node *head, int k)
{
    List::node *p1 = head;
    List::node *p2 = head;

    for (int i = 0; i < k; i++)
    {
        if (head == nullptr)
            return NULL;
        p1 = p1->next;
    }

    while (p1)
    {
        p1 = p1->next;
        p2 = p2->next;
    }

    return p2;
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
    List::node *kth = nthToLast(l.head, 4);
    std::cout << kth->val;

    std::cout << '\n';

    return 0;
}

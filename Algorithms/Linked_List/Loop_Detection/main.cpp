#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;

    ListNode(int x) : val(x), next(NULL) {}
};

ListNode *detectCycle(ListNode *head)
{
    ListNode *slower = head;
    ListNode *faster = head;

    while (faster && faster->next)
    {
        slower = slower->next;
        faster = faster->next->next;

        if (slower == faster)
            break;
    }

    if (!faster || !faster->next)
        return NULL;

    slower = head;
    while (slower != faster)
    {
        slower = slower->next;
        faster = faster->next;
    }
    return faster;
}

int main(int argc, char **argv)
{
    ListNode *head = new ListNode(1);
    ListNode *tail = head;

    tail->next = new ListNode(2);
    tail = tail->next;

    tail->next = new ListNode(3);
    tail = tail->next;

    ListNode *collide = tail;

    tail->next = new ListNode(4);
    tail = tail->next;

    tail->next = new ListNode(5);
    tail = tail->next;

    tail->next = new ListNode(6);
    tail = tail->next;

    tail->next = collide;

    std::cout << detectCycle(head);

    return 0;
}
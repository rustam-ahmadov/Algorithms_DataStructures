#include "../Linked_List/list.h"

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *addSum(ListNode *l1, ListNode *l2)
{
    ListNode *tail = nullptr;
    ListNode *head = nullptr;
    bool check = false;

    while (l1 || l2)
    {
        int sum = 0, digit1 = 0, digit2 = 0;
        if (l1)
            digit1 = l1->val;
        if (l2)
            digit2 = l2->val;

        if (check)
            sum = digit1 + digit2 + 1;
        else
            sum = digit1 + digit2;

        if (sum >= 10)
        {
            sum = sum % 10;
            check = true;
        }
        else
            check = false;

        if (head)
        {
            tail->next = new ListNode(sum);
            tail = tail->next;
        }
        else
        {
            head = new ListNode(sum);
            tail = head;
        }

        if (l1)
            l1 = l1->next;
        if (l2)
            l2 = l2->next;
    }
    if (check)
    {
        tail->next = new ListNode(1);
    }
    return head;
}

// 156 = 651
// 295 = 592
//     +
// 3421 1243

int main(int argc, char **argv)
{

    ListNode *l1 = new ListNode(9);
    l1->next = new ListNode(9);
    l1->next->next = new ListNode(9);

    ListNode *l2 = new ListNode(9);
    l2->next = new ListNode(9);
    l2->next->next = new ListNode(9);
    l2->next->next->next = new ListNode(9);
    l2->next->next->next->next = new ListNode(9);

    ListNode *sum = addSum(l1, l2);

    while (sum)
    {
        std::cout << sum->val << ' ';
        sum = sum->next;
    }

    std::cout << '\n';

    return 0;
}
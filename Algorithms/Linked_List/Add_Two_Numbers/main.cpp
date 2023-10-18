#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
{
    ListNode *res = nullptr;
    ListNode *temp = nullptr;
    int sum = 0, left = 0;
    while (l1 || l2 || left)
    {
        if (res == nullptr)
        {

            res = new ListNode(0);
            temp = res;
        }
        else
        {
            temp->next = new ListNode(0);
            temp = temp->next;
        }

        if (l1)
        {
            sum += l1->val;
            l1 = l1->next;
        }
        if (l2)
        {
            sum += l2->val;
            l2 = l2->next;
        }
        sum += left;

        if (!(sum / 10))
        {
            temp->val = sum;
            left = 0;
        }
        else
        {
            temp->val = sum % 10;
            left = 1;
        }

        sum = 0;
    }
    return res;
}

int main()
{

    ListNode *head1 = new ListNode(2);
    head1->next = new ListNode(4);
    head1->next->next = new ListNode(3);

    ListNode *head2 = new ListNode(5);
    head2->next = new ListNode(6);
    head2->next->next = new ListNode(4);

    addTwoNumbers(head1, head2);

    return 0;
}
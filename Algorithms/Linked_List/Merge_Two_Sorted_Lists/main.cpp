#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *mergeTwoListMySolution(ListNode *list1, ListNode *list2)
{
    ListNode *res = nullptr;
    ListNode *temp = nullptr;
    int min_num = 0;
    bool isNotEnd = true;
    while (true)
    {
        if (list1 && list2)
        {
            if (list1->val < list2->val)
            {
                min_num = list1->val;
                list1 = list1->next;
            }
            else
            {
                min_num = list2->val;
                list2 = list2->next;
            }
        }
        else if (list1)
        {
            min_num = list1->val;
            list1 = list1->next;
        }
        else if (list2)
        {
            min_num = list2->val;
            list2 = list2->next;
        }
        else
            break;

        if (res == nullptr)
        {
            res = new ListNode(min_num);
            temp = res;
        }
        else
        {
            temp->next = new ListNode(min_num);
            temp = temp->next;
        }
    }

    return res;
}

ListNode *mergeTwoLists(ListNode *list1, ListNode *list2)
{
    ListNode *res = nullptr, *temp = nullptr;
    temp = res = new ListNode();

    while (list1 && list2)
    {
        if (list1->val < list2->val)
        {
            temp->next = new ListNode(list1->val);
            list1 = list1->next;
        }
        else
        {
            temp->next = new ListNode(list2->val);
            list2 = list2->next;
        }
        temp = temp->next;
    }

    if (list1)
        temp->next = list1;
    else if (list2)
        temp->next = list2;

    return res->next;
}

int main()
{

    ListNode *head1 = new ListNode(1);
    head1->next = new ListNode(3);
    head1->next->next = new ListNode(5);

    ListNode *head2 = new ListNode(1);
    head2->next = new ListNode(4);
    head2->next->next = new ListNode(6);
    head2->next->next->next = new ListNode(9);

    mergeTwoLists(head1, head2);

    return 0;
}
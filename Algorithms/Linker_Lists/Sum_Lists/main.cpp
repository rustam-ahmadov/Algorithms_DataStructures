#include "../Linked_List/list.h"

// int reverse(int num)
// {
//     int reversed = 0;
//     while (num)
//     {
//         reversed = reversed * 10 + num % 10;
//         num /= 10;
//     }
//     return reversed;
// }

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

int reverse(int num)
{
    int reversed = 0;
    while (num)
    {
        reversed = reversed * 10 + num % 10;
        num /= 10;
    }
    return reversed;
}
ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
{

    int head1_num = 0, head2_num = 0;
    while (l1)
    {
        head1_num = head1_num * 10 + l1->val;
        l1 = l1->next;
    }
    while (l2)
    {
        head2_num = head2_num * 10 + l2->val;
        l2 = l2->next;
    }

    head1_num = reverse(head1_num);
    head2_num = reverse(head2_num);

    int sum = head1_num + head2_num;

    ListNode *node = NULL;
    while (sum)
    {
        ListNode *temp = node->next;
        if (node)
        {
            temp = new ListNode(sum % 10);
            temp->next = NULL;
        }
        else
            node = new ListNode(sum % 10);
        sum /= 10;
    }
    return node;
}

int main(int argc, char **argv)
{
    // List l1;
    // l1.pushBack(7);
    // l1.pushBack(1);
    // l1.pushBack(6);

    // List l2;
    // l2.pushBack(5);
    // l2.pushBack(9);
    // l2.pushBack(2);

    // List reversed_list_sum = addList(l1.head, l2.head);
    // reversed_list_sum.print();

    ListNode *head1 = new ListNode(2);
    head1->next = new ListNode(2);
    head1->next->next = new ListNode(3);

    ListNode *head2 = new ListNode(5);
    head2->next = new ListNode(6);
    head2->next->next = new ListNode(4);

    head2 = addTwoNumbers(head1, head2);
    while (head2)
    {
        std::cout << head2->val << ' ';
        head2 = head2->next;
    }
    return 0;
}
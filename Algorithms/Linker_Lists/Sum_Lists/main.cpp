#include "../Linked_List/list.h"

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

List addList(List::node *head1, List::node *head2)
{
    List reversed_list_sum;
    int head1_num = 0, head2_num = 0;
    while (head1)
    {
        head1_num = head1_num * 10 + head1->val;
        head1 = head1->next;
    }
    while (head2)
    {
        head2_num = head2_num * 10 + head2->val;
        head2 = head2->next;
    }

    head1_num = reverse(head1_num);
    head2_num = reverse(head2_num);

    int sum = head1_num + head2_num;

    while(sum)
    {

        reversed_list_sum.pushBack(sum%10);
        sum/=10;
    }
    return reversed_list_sum;

}

int main(int argc, char **argv)
{
    List l1;
    l1.pushBack(7);
    l1.pushBack(1);
    l1.pushBack(6);

    List l2;
    l2.pushBack(5);
    l2.pushBack(9);
    l2.pushBack(2);
    
    List reversed_list_sum = addList(l1.head,l2.head);
    reversed_list_sum.print();
    return 0;
}
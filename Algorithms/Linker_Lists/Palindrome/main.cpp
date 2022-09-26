#include "../Linked_List/list.h"
#include <vector>

// O(n) my Solution
bool palindromeList(List::node *head)
{
    std::vector<int> v;
    while (head)
    {
        v.push_back(head->val);
        head = head->next;
    }

    int n = v.size();
    if (n < 2)
        return true;

    for (int i = 0, j = n - 1; i < n / 2, j >= n / 2; i++, j--)
    {
        if (v[i] != v[j])
            return false;
    }
    return true;
}

bool equal(List::node *head1, List::node *head2,int half)
{
    while ((head1 && head2 ) && half)
    {
        if (head1->val != head2->val)
            return false;
        head1 = head1->next;
        head2 = head2->next;
    }
    return true;
}

List::node *reverseAndClone(List::node *head, int &size)
{
    List::node *reversed = NULL;

    while (head)
    {
        List::node *temp = new List::node;
        temp->val = head->val;

        temp->next = reversed;
        reversed = temp;

        head = head->next;
        size++;
    }

    return reversed;
}
// CCI p217
bool palindrome(List::node *head)
{
    int half = 0;
    List::node *reversed = reverseAndClone(head, half);
    return equal(head, reversed, half);
}

int main(int argc, char **argv)
{
    List l;

    l.pushBack(0);
    l.pushBack(1);
    l.pushBack(2);

    l.pushBack(2);
    l.pushBack(1);
    l.pushBack(0);

    List l1;

    l1.pushBack(0);
    l1.pushBack(1);
    l1.pushBack(2);

    l1.pushBack(3);

    l1.pushBack(2);
    l1.pushBack(1);
    l1.pushBack(0);

    std::cout << palindromeList(l.head);
    std::cout << '\n';
    std::cout << palindromeList(l1.head);

    std::cout << '\n';

    std::cout<<palindrome(l.head);
    std::cout << '\n';
    std::cout<<palindrome(l1.head);

    std::cout << '\n';

    return 0;
}
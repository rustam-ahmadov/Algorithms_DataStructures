#include "list.h"
#include <unordered_map>

// O(n) finds duplicates in linked list and mark them by increasing
// the value of key(which is the value of linked list)by using Hash_Table(unordered_map)
void findDuplicates(std::unordered_map<int, int> &table, List::node *head)
{
    List::node *temp = head;
    // find duplicates by increasing
    while (temp)
    {
        table[temp->val]++;
        temp = temp->next;
    }
}

// O(n) removes duplicates from linked list by checking if value of node(key) is > 1
void removeDuplicates(List::node *&head)
{
    std::unordered_map<int, int> table;
    List::node *temp = head, *before_temp = head;
    if (!head)
        return;
    findDuplicates(table, head);
    while (temp)
    {
        if (table[temp->val] == 1)
        {
            before_temp = temp;
            temp = temp->next;
        }
        else if (table[temp->val] > 1)
        {
            table[temp->val]--;
            if (temp != head)
            {
                before_temp->next = temp->next;
                delete temp;
                temp = before_temp->next;
            }
            else
            {
                head = head->next;
                delete temp;
                temp = head;
                before_temp = head;
            }
        }
    }
}
// O(n^2) remove duplicates without using additional space (table)
void removeDuplicates(List::node *&head, bool n_2)
{
    List::node *temp = head;
    List::node *before_i = head;
    List::node *i = head->next;
    while (temp)
    {
        i = temp->next;        
        before_i = temp;
        while (i)
        {
            if (temp->val == i->val)
            {
                before_i->next = i->next;
                delete i;
                i = before_i->next;
            }
            else
            {
                before_i = i;
                i = i->next;
            }
        }
        temp = temp->next;        
    }
}

int main(int argc, char **argv)
{
    List l;

    l.pushBack(1);
    l.pushBack(1);
    l.pushBack(2);

    l.pushBack(4);
    l.pushBack(4);

    l.pushBack(5);

    l.pushBack(6);
    l.pushBack(6);
    l.pushBack(6);
    l.pushBack(6);
    l.pushBack(6);

    l.pushBack(7);

    l.print();
    removeDuplicates(l.head, true);
    l.print();

    return 0;
}
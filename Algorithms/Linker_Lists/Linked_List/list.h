#include <iostream>

class List
{
public:
    struct node
    {
        int val;
        node *next;
    };
    node *head = nullptr;
    node *tail = nullptr;
    int size = 0;



    int getSize()
    {
        return size;
    }

    void print()
    {
        if (!size&&!head)
            return;

        node *temp = head;
        while (temp)
        {
            std::cout << temp->val << ' ';
            temp = temp->next;
        }
        std::cout<<'\n';
    }

    void pushBack(int val)
    {
        if (!head)
        {
            head = new node;
            head->val = val;
            head->next = nullptr;
            tail = head;
        }
        else
        {
            tail->next = new node;
            tail->next->val = val;
            tail->next->next = nullptr;

            tail = tail->next;
        }
        size++;
    }
};
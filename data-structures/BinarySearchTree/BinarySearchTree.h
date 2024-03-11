#include <iostream>
#include <conio.h>
#include <stack>

class BinarySearchTree
{
private:
    struct node
    {
        int value;
        node *left = nullptr;
        node *right = nullptr;

        node(int value)
        {
            this->value = value;
            left = nullptr;
            right = nullptr;
        }
    };
    node *root = nullptr;

    void makeEmpty(node *p)
    {
        if (!p)
            return;
        else
        {
            makeEmpty(p->left);
            makeEmpty(p->right);
        }
        delete p;
        p = nullptr;
    }

    node *remove(int value, node *p)
    {
        if (p == nullptr)
            return p;
        else if (value < p->value)
            p->left = remove(value, p->left);
        else if (value > p->value)
            p->right = remove(value, p->right);
        else
        {
            // Case 1: No child
            if (p->left == nullptr && p->right == nullptr)
            {
                delete p;
                p = nullptr;
            }
            // Case 2: one child (Right)
            else if (p->left == nullptr)
            {
                node *temp = p;
                p = p->right;
                delete temp;
            } // Case 2: one child (Left)
            else if (p->right == nullptr)
            {
                node *temp = p;
                p = p->left;
                delete temp;
            }
            // Case 4: two child
            else if (p->left && p->right)
            {
                node *min = findMin(p->right);
                p->value = min->value;
                p->right = remove(min->value, p->right);
            }
        }
        return p;
    }

    node *findMin(node *root)
    {
        node *temp = root;

        while (temp)
        {
            if (temp->left == nullptr)
                return temp;
            temp = temp->left;
        }
    }

    // just for display
    std::stack<node *> s;
    void display(node *p, bool isRepeated)
    {
        if (!isRepeated)
            s.push(p);

        system("cls");
        if (p != nullptr)
        {
            std::cout << "         " << p->value;
            std::cout << '\n';
            if (p->left != nullptr)
            {

                std::cout << "" << p->left->value;
            }
            if (p->right != nullptr)
            {
                std::cout << "              " << p->right->value;
            }

            int choice = 0;
            choice = _getch();
            // left subtree
            if (choice == 97)
            {
                display(p->left, false);
            }
            // right subtree
            else if (choice == 100)
                display(p->right, false);
            else if (choice == 119)
            {

                if (s.size() > 1)
                {
                    s.pop();
                    display(s.top(), true);
                }
                else
                {
                    display(p, false);
                }
            }
        }
        return;
    }

public:
    void makeEmpty()
    {
        makeEmpty(root);
    }

    int remove(int value)
    {
        remove(value, root);
        return 0;
    }

    node *findMin()
    {
        return findMin(root);
    }

    // return -1 if there is already such value
    // in case of success returns value wich has been inserted
    int insert(int value)
    {
        if (root == nullptr)
        {
            root = new node(value);
            return value;
        }

        node *temp = root;
        while (temp)
        {
            if (temp->value == value)
                return -1;
            // if value greater go to the right
            if (value > temp->value)
            {
                // if right child is null then create new child
                if (temp->right == nullptr)
                {
                    temp->right = new node(value);
                    return value;
                }
                else
                    temp = temp->right;
            }
            // if value is less go to the left
            else
            {
                // if left child is null then create new child
                if (temp->left == nullptr)
                {
                    temp->left = new node(value);
                    return value;
                }
                else
                    temp = temp->left;
            }
        }
    }

    void display()
    {
        display(root, false);
    }
};
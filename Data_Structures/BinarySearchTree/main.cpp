#include <iostream>

class BinarySearchTree
{
    struct node
    {
        int value;
        node *left;
        node *right;
        int height;
    };
    node *root;

    void makeEmpty(node *p)
    {
        if (!root)
            return;
        else
        {
            makeEmpty(root->left);
            makeEmpty(root->right);
        }
        delete p;
    }

public:
    void makeEmpty()
    {
        makeEmpty(root);
    }
};

int main(int argc, char **argv)
{
    return 0;
}
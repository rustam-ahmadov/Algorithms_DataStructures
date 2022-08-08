#include <iostream>

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

// First approach
//  O(n^2)
bool subTreeLesser(TreeNode *, int);
bool subTreeGreater(TreeNode *root, int val)
{
    if (!root)
        return true;
    if (root->val > val && subTreeGreater(root->right, val) && subTreeGreater(root->left, val))
        return true;
    return false;
}
bool subTreeLesser(TreeNode *root, int val)
{
    if (!root)
        return true;
    if (root->val < val && subTreeLesser(root->left, val) && subTreeLesser(root->left, val))
        return true;
    return false;
}


bool isValidBST(TreeNode *root)
{
    if (!root)
        return true;

    if (subTreeGreater(root->right, root->val) && subTreeLesser(root->left, root->val) && isValidBST(root->left) && isValidBST(root->right))
        return true;
    return false;
}

// Second approach
// O(n)
bool searchTree(TreeNode *root, int min_value, int max_value)
{
    if (!root)
        return true;
    if (root->val > min_value && root->val < max_value &&
        searchTree(root->left, min_value, root->val - 1) &&
        searchTree(root->right, root->val + 1, max_value))
        return true;

    return false;
}

int main(int argc, char **argv)
{

    TreeNode *root = new TreeNode(50);

    // Right side
    root->right = new TreeNode(100);
    root->right->left = new TreeNode(75);
    root->right->right = new TreeNode(150);
    // Left side

    root->left = new TreeNode(25);
    root->left->right = new TreeNode(35);
    root->left->left = new TreeNode(15);

    std::cout << searchTree(root, INT_MIN, INT_MAX);

    return 0;
}
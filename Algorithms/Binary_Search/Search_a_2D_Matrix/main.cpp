#include <iostream>
#include <vector>

using namespace std;

bool searchMatrix(vector<vector<int>> &matrix, int target)
{
    int i_m = 0;
    const int m = matrix.size();
    int n = matrix[0].size();
    for (int i = 1; i < m; i++)
    {
        if (target <= matrix[i][n - 1] && target >= matrix[i][0])
        {
            i_m = i;
            break;
        }
    }

    int l = 0, r = n - 1, mid = r / 2;
    while (l <= r)
    {
        if (target == matrix[i_m][mid])
            return true;

        if (target > matrix[i_m][mid])
            l = mid + 1;
        else
            r = (l + r) / 2 - 1;

        mid = (r + l) / 2;
    }
    return false;
}

int main()
{
    vector<vector<int>> matrix = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
    std::cout<<searchMatrix(matrix, 5);
}
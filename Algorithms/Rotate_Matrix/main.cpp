#include <iostream>
// rows and columns size
const int R = 4;
const int C = 4;


//time O(n) space O(n) 
void rotate(int matrix[][C], int row, int column)
{
    int rotated_matrix[row][column];

    for (int i = 0; i < column; i++)
        for (int j = row - 1, new_matrix_column = 0; j >= 0; j--, new_matrix_column++)
            rotated_matrix[i][new_matrix_column] = matrix[j][i];

    for (int i = 0; i < row; i++)
        for (int j = 0; j < column; j++)
            matrix[i][j] = rotated_matrix[i][j];
}
void print(int matrix[][C], int row, int column)
{
    for (int i = 0; i < row; i++)
    {
        for (int j = 0; j < column; j++)
        {
            std::cout << matrix[i][j] << ' ';
        }
        std::cout << '\n';
    }
}

int main(int argc, char **argv)
{

    int matrix[R][C] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16};

    rotate(matrix, R, C);
    print(matrix, R, C);
    return 0;
}
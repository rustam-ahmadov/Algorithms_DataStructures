#include <iostream>

void nullifyRow(int matrix[][4], int C, int row)
{
    for (int i = 0; i < C; i++)
        matrix[row][i] = 0;
}
void nullifyColumn(int matrix[][4], int R, int column)
{
    for (int i = 0; i < R; i++)
        matrix[i][column] = 0;
}

void setZeros(int matrix[][4], int R, int C)
{
    bool row[R];
    bool column[C];

    // Store the row and column index with value 0
    for (int i = 0; i < R; i++)
    {
        for (int j = 0; j < C; j++)
        {
            if (matrix[i][j] == 0)
            {
                row[i] = true;
                column[j] = true;
            }
        }
    }

    // Nullify rows
    for (int i = 0; i < R; i++)
        if (row[i] == true)
            nullifyRow(matrix, C, i);

    // Nullify columns
    for (int i = 0; i < C; i++)
        if (column[i] == true)
            nullifyColumn(matrix, R, i);
}

void printMatrix(int matrix[][4], int R, int C)
{
    for (int i = 0; i < R; i++)
    {
        for (int j = 0; j < C; j++)
            std::cout << matrix[i][j] << ' ';
        std::cout << '\n';
    }
}

int main(int argc, char **argv)
{
    int matrix[3][4] = {1, 2, 0, 4, 
    5, 6, 7, 8, 0, 9, 10};

    setZeros(matrix,3,4);
    printMatrix(matrix,3,4);

    return 0;
}
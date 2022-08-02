#include <iostream>

// First two functions are just for demo
void fill_array(int *array, const int ARRAY_LENGTH);
void print_array(const int *array, const int ARRAY_LENGTH);

void sort_array(int *array, const int ARRAY_LENGTH);

// I use BINARY SEARCH in this function  and return true
// if array contains number otherwise return false
bool is_array_contains(const int *array, const int ARRAY_LENGTH, int NUMBER);

int main()
{
    srand(time(NULL));

    // Just for demo
    const int NUMBER = 8;
    const int ARRAY_LENGTH = 8;
    int *p_arr = new int[ARRAY_LENGTH];

    fill_array(p_arr, ARRAY_LENGTH);
    print_array(p_arr, ARRAY_LENGTH);

    // Then sort and print again to see the changes
    sort_array(p_arr, ARRAY_LENGTH);
    print_array(p_arr, ARRAY_LENGTH);

    bool is_contains = is_array_contains(p_arr, ARRAY_LENGTH, NUMBER);

    if (is_contains)
        std::cout << "Yeap, there is a such number" << std::endl;
    else
        std::cout << "Nope, there isn't a such number" << std::endl;

    return 0;
}

void fill_array(int *array, const int ARRAY_LENGTH)
{

    for (int i = 0; i < ARRAY_LENGTH; i++)   
        array[i] = i + 1;
    rand() % (ARRAY_LENGTH * 2 ) + 1; 
    
}

void print_array(const int *array, const int ARRAY_LENGTH)
{
    std::cout << "Array length: " << ARRAY_LENGTH << std::endl;
    std::cout << "Array elements:" << std::endl;
    for (int i = 0; i < ARRAY_LENGTH; i++)
    {
        std::cout << i << ":" << array[i] << std::endl;
    }
    std::cout << std::endl;
}

void sort_array(int *array, const int ARRAY_LENGTH)
{
    for (int i = 0; i < ARRAY_LENGTH; i++)
    {
        for (int j = 0; j < ARRAY_LENGTH - 1; j++)
        {
            if (array[j] > array[j + 1])
            {
                int temp = array[j];
                array[j] = array[j + 1];
                array[j + 1] = temp;
            }
        }
    }
}


//O(log n)
bool is_array_contains(const int *array, const int ARRAY_LENGTH, const int NUMBER)
{
    bool is_contains = false;

    int scope_first_index = 0;
    int scope_last_index = ARRAY_LENGTH - 1;
    int scope_middle_index = (scope_first_index + scope_last_index) / 2;

    while (scope_first_index <= scope_last_index)
    {
        if (NUMBER == array[scope_middle_index])
        {
            is_contains = true;
            break;
        }
        else if (NUMBER > array[scope_middle_index])
            scope_first_index = scope_middle_index + 1;
        else
            scope_last_index = scope_middle_index - 1;

        scope_middle_index = (scope_first_index + scope_last_index) / 2;
    }
    return is_contains;
}

void printArray(int[] array)
{
    for (int i = 0; i < array.Length; i++)
    {
        System.Console.Write(array[i] + " ");
    }
}



void quickSort(int[] array, int lowIndex, int highIndex)
{
    if (lowIndex < highIndex)
    {
        int pivot = partition(array, lowIndex, highIndex);
        quickSort(array, lowIndex, pivot - 1);
        quickSort(array, pivot + 1, highIndex);
    }
}
int partition(int[] array, int lowIndex, int highIndex)
{
    int pivot = array[highIndex];
    int i = lowIndex - 1;

    for (int j = lowIndex; j < highIndex; j++)
    {

        if (array[j] <= pivot)
        {
            i++;
            swap(array, j, i);
        }
    }
    i++;
    swap(array, i, highIndex);
    return i;
}

static void swap(int[] array, int index1, int index2)
{
    int temp = array[index1];
    array[index1] = array[index2];
    array[index2] = temp;
}


const int ARRAY_SIZE = 8;

Random rand = new Random();
int[] numbers = new int[ARRAY_SIZE]{51,2,13,80,15,30,60,49};



System.Console.WriteLine("Before:");
printArray(numbers);

quickSort(numbers, 0, numbers.Length - 1);

System.Console.WriteLine("\n\nAfter:");
printArray(numbers);

//O(n * n)
int[] SelectionSort(int[] Array)
{
    int min = 0;
    int temp = 0;

    for (int i = 0; i < Array.Length; i++)
    {
        min = i;

        for (int j = i + 1; j < Array.Length; j++)
        {
            if (Array[j] < Array[min])
                min = j;
        }

        if (i != min)
        {
            temp = Array[i];
            Array[i] = Array[min];
            Array[min] = temp;
        }
    }
    return Array;
}



void PrintArray(int[] Array)
{
    foreach (var num in Array)
    {
        System.Console.WriteLine(num);
    }
}

int[] Array = new int[] { 10, 2, 8, 3, 6, 1, 9, 5, 7, 4, 0 };
SelectionSort(Array);
PrintArray(Array);
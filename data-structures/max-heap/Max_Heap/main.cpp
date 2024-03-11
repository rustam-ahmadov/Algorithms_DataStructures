#include <iostream>
using namespace std;

void swap(int *x, int *y)
{
  int temp = *x;
  *x = *y;
  *y = temp;
}

class max_heap
{
public:
  max_heap(int capacity)
  {
    _size = 0;
    _capacity = capacity;
    _arr = new int[capacity];
  }

  // A recursive method to heapify a subtree with the root at given index
  // This method assumes that the subtrees are already heapified
  void heapify(int i)
  {
    int l = left(i);
    int r = right(i);
    int biggest = i;

    if (l < _size && _arr[l] > _arr[i])
      biggest = l;
    if (r < _size && _arr[r] > _arr[biggest])
      biggest = r;
    if (biggest != i)
    {
      swap(&_arr[i], &_arr[biggest]);
      heapify(biggest);
    }
  }
  // Removes max element(root) from max_heap and return it
  // The time Complexity of this Operation is O(log N) as this operation needs to maintain
  // the heap property (by calling heapify()) after removing the root.
  int extractMax()
  {
    if (!_size)
      return INT_MIN;
    if (_size == 1)
    {
      _size--;
      return _arr[0];
    }

    // store the max val and remove it from heap
    int root = _arr[0];
    _arr[0] = _arr[_size - 1]; // swapping root with the last added element
    _size--;
    heapify(0);
  }

  void deleteKey(int i)
  {
    increaseKey(i, INT_MAX);
    extractMax();
  }

  void increaseKey(int i, int new_val)
  {
    _arr[i] = new_val;
    while (i && _arr[parent(i)] < _arr[i])
    {
      swap(&_arr[i], &_arr[parent(i)]);
      i = parent(i);
    }
  }

  void insertKey(int k)
  {
    if (_size == _capacity)
    {
      cout << "\nOverflow: Could not insert key\n";
      return;
    }
    _size++;
    int i = _size - 1;
    _arr[i] = k;

    // Fix the max heap property at the end
    while (i && _arr[i] > _arr[parent(i)])
    {
      swap(&_arr[i], &_arr[parent(i)]);
      i = parent(i);
    }
  }

  int parent(int i) { return (i - 1) / 2; }

  // index of left child of node at index i
  int left(int i) { return (2 * i) + 1; }

  // index of right child of node at index i
  int right(int i) { return (2 * i) + 2; }

  int getMax() { return _arr[0]; }

private:
  int *_arr;
  int _size = 0;
  int _capacity = 10;
};

int main()
{
  max_heap mh(11);
  mh.insertKey(3);
  mh.insertKey(2);
  mh.insertKey(10);
  std::cout << mh.getMax();
  mh.insertKey(11);
  std::cout << mh.getMax();
  mh.deleteKey(0);
  std::cout << mh.getMax();

  return 0;
}
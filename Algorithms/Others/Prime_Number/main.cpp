#include <iostream>


//O(log n)
bool isPrime(int number)
{
  for(int i = 2; i * i <= number; i++)
  {
    if(number % 2 == 0)
      return false;
  }
  return true;
}

int main()
{

  int number = 10;
  std::cout << isPrime(137);

  return 0;
}
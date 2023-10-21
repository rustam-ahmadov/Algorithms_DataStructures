#include <iostream>
#include <vector>
#include <algorithm>


using namespace std;

//O(n(log n) + m * n) = (m * n) where m = max number(we just check from 1 to max number)
int minEatingSpeedBruteForse(vector<int> &piles, int h)
{
    const int n = piles.size();
    int k = 0;
    if (h < n)
        return k;

    sort(piles.begin(), piles.end());

    int max = piles[n - 1];
    int left_h = h;
    for (int i = 1; i <= max; i++)
    {
        for (int j = 0; j < n; j++)
        {
            int p = piles[j];
            p -= i;
            left_h--;
            while (p > 0)
            {
                p -= i;
                left_h--;
            }
            
            if (n - (j + 1) > left_h)
                break;

            if (n - (j + 1) == 0 && 0 == left_h)
                return i;
        }
        left_h = h;
    }

    return k;
}

int main()
{

    vector<int> piles = {30,11,23,4,20};
    int h = 5;

    cout << minEatingSpeed(piles, h);
    return 0;
}
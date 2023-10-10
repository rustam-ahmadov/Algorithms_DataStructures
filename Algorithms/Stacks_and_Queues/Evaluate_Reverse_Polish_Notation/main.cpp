#include <iostream>
#include <stack>
#include <vector>

using namespace std;

int getNumFromStr(string str)
{
    int n = str.size();
    int min = 1;
    if (str.front() == '-')
    {
        min = -1;
        str = str.substr(1);
        n--;
    }
    int res = 0;
    for (int i = 0; i < n; i++)
        res = res * 10 + str[i] - '0';
    return res * min;
}

int getRes(int v1, int v2, char operation)
{
    switch (operation)
    {
    case '*':
        return v1 * v2;
    case '/':
        return v1 / v2;
    case '+':
        return v1 + v2;
    case '-':
        return v1 - v2;
    }
    return 0;
}

int evalRPN(vector<string> &tokens)
{
    const int n = tokens.size();
    if (!n)
        return 0;

    stack<int> nums;
    int num = 0;
    char operation = ' ';
    for (int i = 0; i < n; i++)
    {
        if (tokens[i] != "*" && tokens[i] != "/" && tokens[i] != "+" && tokens[i] != "-")
            num = getNumFromStr(tokens[i]);
        else
        {
            operation = tokens[i][0];
            int second = nums.top();
            nums.pop();
            int first = nums.top();
            nums.pop();
            num = getRes(first, second, operation);
        }
        nums.push(num);
    }
    return nums.top();
}

int main()
{
    vector<string> tokens = {"3", "-4", "+"};
    std::cout << evalRPN(tokens);

    return 0;
}
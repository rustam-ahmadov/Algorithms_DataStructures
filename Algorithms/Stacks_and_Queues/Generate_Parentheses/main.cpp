#include <iostream>
#include <stack>
#include <vector>

using namespace std;

void generate(int n, int o, int c, string p, vector<string> &parns)
{
    if (o == n && c == n)
        parns.push_back(p);
    else
    {
        if (o < n)
            generate(n, o + 1, c, p + '(', parns);
        if (c < o)
            generate(n, o, c + 1, p + ')', parns);
    }
}

vector<string> generateParenthesis(int n)
{
    vector<string> parns;
    string p = "";
    int o = 0, c = 0;
    generate(n, o, c, p, parns);
    return parns;
}

int main()
{

    vector<string> parens = generateParenthesis(3);
    for (int i = 0; i < parens.size(); i++)
        std::cout << parens[i] << ' ';

    return 0;
}
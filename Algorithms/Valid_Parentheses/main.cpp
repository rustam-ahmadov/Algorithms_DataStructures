#include <iostream>
#include <string>
#include <stack>

bool isMatch(const char first, const char second);
bool isParenthesisOpen(const char c);

bool isValid(std::string &s)
{
    std::stack<char> st;
    int str_length = s.size();
    char parenthesis = ' ';
    for (int i = 0; i < str_length; i++) // iterate over each and every elements
    {
        parenthesis = s[i];
        if (isParenthesisOpen(parenthesis))
            st.push(parenthesis);
        else
        {
            if (st.empty() || !isMatch(st.top(), parenthesis))
                return false;
            st.pop();
        }
    }
    return st.empty();
}

bool isMatch(const char first, const char second)
{
    if (first == '(' && second == ')')
        return true;
    if (first == '[' && second == ']')
        return true;
    if (first == '{' && second == '}')
        return true;
    return false;
}

bool isParenthesisOpen(const char c)
{
    switch (c)
    {
    case '{':
        return true;
    case '[':
        return true;
    case '(':
        return true;
    default:
        return false;
    }
}
int main()
{
    std::string check = "[{}]";
    bool result = isValid(check);

    // bool result = isValid("(){[]}");
    // bool result = isValid("))");

    std::cout << result;
    return 0;
}
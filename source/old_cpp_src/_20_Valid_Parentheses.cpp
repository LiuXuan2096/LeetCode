class Solution {
public:
    bool isValid(string s) {
        stack<char> stack;

        for (char ch : s)
        {
            if (ch == '(' || ch == '[' || ch == '{')
            {
                /* code */
                stack.push(ch);
            }
            else
            {
                if (stack.empty())
                {
                    return false;
                }
                if (ch == ')')
                {
                    /* code */
                    if (stack.top() == '(')
                    {
                        /* code */
                        stack.pop();
                    }
                    else
                    {
                        return false;
                    }       
                }
                if (ch == ']')
                {
                    /* code */
                    if (stack.top() == '[')
                    {
                        /* code */
                        stack.pop();
                    }
                    else
                    {
                        return false;
                    }       
                }
                if (ch == '}')
                {
                    /* code */
                    if (stack.top() == '{')
                    {
                        /* code */
                        stack.pop();
                    }
                    else
                    {
                        return false;
                    }       
                }
                
            }
            
        }
        return stack.empty();
    }
};
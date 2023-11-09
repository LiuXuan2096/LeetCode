class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> numberStack; // 存放参与运算的数字的栈

        for (string s : tokens)
        {
            if (s.size() == 1 && (s == "+" || s == "-" || s == "*" || s == "/"))
            {
                /* 当前字符串是运算符，执行运算*/
                int right = numberStack.top();
                numberStack.pop();
                int left = numberStack.top();
                numberStack.pop();
                if (s == "+")
                {
                    /* code */
                    numberStack.push(left + right);
                }
                else if (s == "-")
                {
                    /* code */
                    numberStack.push(left - right);
                }
                else if (s == "*")
                {
                    /* code */
                    numberStack.push(left * right);
                }
                else if (s == "/")
                {
                    /* code */
                    numberStack.push(left / right);
                }             
            }
            else
            {
                // 当前字符串是数字，进入数字栈
                numberStack.push(stoi(s));
            }
            
        }
        return numberStack.top();
    }
};
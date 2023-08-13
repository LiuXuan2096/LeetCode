class MyQueue {
public:
    MyQueue() {
    
    }
    
    void push(int x) {
        s1.push(x);
    }
    
    int pop() {
        int res;
        if (s2.empty())
        {
            /* code */
            while (!s1.empty())
            {
                s2.push(s1.top());
                s1.pop();
            }
            res = s2.top();
            s2.pop();
        }
        else
        {
            res = s2.top();
            s2.pop();
        }
        return res;
    }
    
    int peek() {
        int res;
        if (s2.empty())
        {
            /* code */
            while (!s1.empty())
            {
                s2.push(s1.top());
                s1.pop();
            }
            res = s2.top();
        }
        else
        {
            res = s2.top();
        }
        return res;
    }
    
    bool empty() {
        return (s1.empty() && s2.empty());
    }
private:
    stack<int> s1;
    stack<int> s2;
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
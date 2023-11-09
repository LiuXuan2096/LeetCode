class MyStack {
public:
    MyStack() {

    }
    
    void push(int x) {
        input.push(x);
        while (!output.empty())
        {
            input.push(output.front());
            output.pop();
        }
        queue tmp = input;
        input = output;
        output = tmp;
    }
    
    int pop() {
        int res = output.front();
        output.pop();
        return res;
    }
    
    int top() {
        return output.front();
    }
    
    bool empty() {
        return output.empty() && input.empty();
    }
private:
    queue<int> input;  // 指向入队队列
    queue<int> output; // 指向出队队列
};

/**
 * Your MyStack object will be instantiated and called as such:
 * MyStack* obj = new MyStack();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->top();
 * bool param_4 = obj->empty();
 */
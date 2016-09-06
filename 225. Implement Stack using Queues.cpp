class Stack {
public:
    // Push element x onto stack.
    queue<int> q;
    void push(int x) {
        q.push(x);
    }

    // Removes the element on top of the stack.
    void pop() {

        if(q.size() !=1){
                queue<int> tmp;

        while(q.size()!=1){
            tmp.push(q.front());
            q.pop();
        }
        q = tmp;
        }else
            q.pop();
    }

    // Get the top element.
    int top() {
        return q.back();
    }

    // Return whether the stack is empty.
    bool empty() {
        return q.size()==0;
    }
};
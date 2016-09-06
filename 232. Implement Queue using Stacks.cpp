class Queue {
private:
    std::stack<int> stack;
    std::stack<int> stack_reversed;
public:
    // Push element x to the back of queue.
    void push(int x) {
        while(!stack.empty()){
            stack_reversed.push(stack.top());
            stack.pop();
        }
        stack_reversed.push(x);
        while(!stack_reversed.empty()){
            stack.push(stack_reversed.top());
            stack_reversed.pop();
        }
    }

    // Removes the element from in front of queue.
    void pop(void) {
        stack.pop();
    }

    // Get the front element.
    int peek(void) {
        return stack.top();
    }

    // Return whether the queue is empty.
    bool empty(void) {
        return stack.empty();
    }
};
class MyLinkedList {
public:

    // 定义链表节点结构体
    struct LinkedNode
    {
        int val;
        LinkedNode* next;
        LinkedNode(int val):val(val), next(nullptr) {}

    };

    // 初始化链表
    MyLinkedList() {
        _dummyHead = new LinkedNode(0); // 虚拟头结点
        _size = 0;
    }

    // 获取到第index个节点数值，如果index是非法数值直接返回-1，注意index是从0开始的，第0个节点就是头结点
    int get(int index) {
        if (index < 0 || index > _size - 1)
        {
            return -1;
        }
        LinkedNode* cur = _dummyHead->next;
        while (index--)
        {
            cur = cur->next;
        }
        return cur->val;
    }
    
    // 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
    void addAtHead(int val) {
        LinkedNode* newHead = new LinkedNode(val);
        newHead->next = _dummyHead->next;
        _dummyHead->next = newHead;
        _size++;
    }

    // 将值为 val 的节点追加到链表的最后一个元素。
    void addAtTail(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        LinkedNode* cur = _dummyHead;
        while (cur->next != nullptr)
        {
            cur = cur->next;
        }
        cur->next = newNode;
        _size++;
    }
    
    // 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，
    // 则该节点将附加到链表的末尾。
    // 如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
    void addAtIndex(int index, int val) {
        if (index > _size)
        {
            return;
        }
        LinkedNode* newNode = new LinkedNode(val);
        LinkedNode* cur = _dummyHead;
        if (index <= 0)
        {
            newNode->next = _dummyHead->next;
            _dummyHead->next = newNode;
            
        }
        else if (index == _size)
        {
            while (index--)
            {
                cur = cur->next;
            }
            cur->next = newNode;
            
        }
        else
        {
            while (index--)
            {
                cur = cur->next;
            }
            newNode->next = cur->next;
            cur->next = newNode;
           
        }
        _size++;
        return;
    }

    // 如果索引 index 有效，则删除链表中的第 index 个节点。
    void deleteAtIndex(int index) {
        if (index < 0 || index > _size - 1)
        {
            return;
        }
        LinkedNode* cur = _dummyHead;
        while (index--)
        {
            cur = cur->next;
        }
        LinkedNode* deletedNode = cur->next;
        cur->next = deletedNode->next;
        delete deletedNode;
        _size--;
        return;
    }

private:
    int _size;
    LinkedNode* _dummyHead;
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList* obj = new MyLinkedList();
 * int param_1 = obj->get(index);
 * obj->addAtHead(val);
 * obj->addAtTail(val);
 * obj->addAtIndex(index,val);
 * obj->deleteAtIndex(index);
 */
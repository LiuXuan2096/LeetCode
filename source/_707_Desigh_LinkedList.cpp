class MyLinkedList {
public:

    // ��������ڵ�ṹ��
    struct LinkedNode
    {
        int val;
        LinkedNode* next;
        LinkedNode(int val):val(val), next(nullptr) {}

    };

    // ��ʼ������
    MyLinkedList() {
        _dummyHead = new LinkedNode(0); // ����ͷ���
        _size = 0;
    }

    // ��ȡ����index���ڵ���ֵ�����index�ǷǷ���ֱֵ�ӷ���-1��ע��index�Ǵ�0��ʼ�ģ���0���ڵ����ͷ���
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
    
    // ������ĵ�һ��Ԫ��֮ǰ���һ��ֵΪ val �Ľڵ㡣������½ڵ㽫��Ϊ����ĵ�һ���ڵ㡣
    void addAtHead(int val) {
        LinkedNode* newHead = new LinkedNode(val);
        newHead->next = _dummyHead->next;
        _dummyHead->next = newHead;
        _size++;
    }

    // ��ֵΪ val �Ľڵ�׷�ӵ���������һ��Ԫ�ء�
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
    
    // �������еĵ� index ���ڵ�֮ǰ���ֵΪ val  �Ľڵ㡣��� index ��������ĳ��ȣ�
    // ��ýڵ㽫���ӵ������ĩβ��
    // ��� index ���������ȣ��򲻻����ڵ㡣���indexС��0������ͷ������ڵ㡣
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

    // ������� index ��Ч����ɾ�������еĵ� index ���ڵ㡣
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
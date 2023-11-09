class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* dummyHead = new ListNode(0);  // 设置一个虚拟头结点
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        int size = 0;
        while (cur->next != nullptr)
        {
            size++;
            cur = cur->next;
        }
        int length = size - n; // 得到要删除的节点前一个节点离头结点的距离
        cur = dummyHead;
        while (length--)
        {
            cur = cur->next;
        }
        ListNode* tmp = cur->next;
        cur->next = tmp->next;
        delete tmp;
        return dummyHead->next;
    }
};
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* dummyHead = new ListNode(0);  // ����һ������ͷ���
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        int size = 0;
        while (cur->next != nullptr)
        {
            size++;
            cur = cur->next;
        }
        int length = size - n; // �õ�Ҫɾ���Ľڵ�ǰһ���ڵ���ͷ���ľ���
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
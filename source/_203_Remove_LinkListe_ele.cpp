class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* dummyHead = new ListNode(0); // ��������ͷ���
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        // ѭ���ж��Ƿ�ɾ��,ѭ�������������Ǳ�������������
        while (cur->next != nullptr)
        {
            if (cur->next->val == val)
            {
                ListNode* tmp = cur->next;
                cur->next = cur->next->next;
                delete tmp;
            }
            else
            {
                cur = cur->next;
            }
 
        }
        head = dummyHead->next;
        delete dummyHead;
        return head;
        
    }
};
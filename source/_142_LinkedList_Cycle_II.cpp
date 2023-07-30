class Solution {
public:
    ListNode* detectCycle(ListNode* head) {

        if (head == nullptr)
        {
            return head;
        }

        ListNode* fast = head;
        ListNode* slow = head;
        
        // �ж������Ƿ��л�
        while (fast != nullptr && fast->next != nullptr && fast->next->next != nullptr)
        {
            // ����fast�ڵ��slow�ڵ�
            fast = fast->next->next;
            slow = slow->next;
            if (fast == slow)
            {
                fast = head;
                while (fast != slow)
                {
                    fast = fast->next;
                    slow = slow->next;
                }
                return fast;
            }
        }
        return nullptr;
    }
};
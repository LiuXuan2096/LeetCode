class Solution {
public:
    ListNode* detectCycle(ListNode* head) {

        if (head == nullptr)
        {
            return head;
        }

        ListNode* fast = head;
        ListNode* slow = head;
        
        // 判断链表是否有环
        while (fast != nullptr && fast->next != nullptr && fast->next->next != nullptr)
        {
            // 更新fast节点和slow节点
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
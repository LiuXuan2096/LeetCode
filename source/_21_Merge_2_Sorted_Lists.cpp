class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode* p = l1;
        ListNode* q = l2;
        ListNode* newHead = nullptr;  // 新链表的头节点
        ListNode* last = nullptr;  // 新链表的尾节点
        if (l1 == nullptr && l2 == nullptr)
        {
            return l1;
        }
        else if (l1 == nullptr)
        {
            return l2;
        }
        else if (l2 == nullptr)
        {
            return l1;
        }
        else
        {
            if (p->val < q->val)
            {
                newHead = l1;
                last = l1;
                p = p->next;
            }
            else
            {
                newHead = l2;
                last = l2;
                q = q->next;
            }
            while (p != nullptr && q != nullptr)
            {
                if (p->val < q->val)
                {
                    last->next = p;
                    last = p;
                    p = p->next;
                }
                else
                {
                    last->next = q;
                    last = q;
                    q = q->next;
                }
            }
            if (p == nullptr)
            {
                last->next = q;
            } 
            else 
            {
                last->next = p;
            }
            return newHead;
        }
    }
};
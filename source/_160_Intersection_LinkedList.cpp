class Solution {
public:
    ListNode* getIntersectionNode(ListNode* headA, ListNode* headB) {
        int lengthA = 0;
        int lengthB = 0;
        ListNode* curA = headA;
        ListNode* curB = headB;
        while (curA != nullptr)
        {
            curA = curA->next;
            lengthA++;
        }
        while (curB != nullptr)
        {
            curB = curB->next;
            lengthB++;
        }
        curA = headA;
        curB = headB;
        if (lengthA > lengthB)
        {
            int offset = lengthA - lengthB;
            while (offset--)
            {
                curA = curA->next;
            }
            while (curA != nullptr)
            {
                if (curA == curB) {
                    return curA;
                }
                curA = curA->next;
                curB = curB->next;
            }
            return nullptr;
        }
        else
        {
            int offset = lengthB - lengthA;
            while (offset--)
            {
                curB = curB->next;
            }
            while (curB != nullptr)
            {
                if (curB == curA) {
                    return curB;
                }
                curA = curA->next;
                curB = curB->next;
            }
            return nullptr;
        }
        

    }
};
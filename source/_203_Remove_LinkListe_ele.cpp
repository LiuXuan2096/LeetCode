class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* dummyHead = new ListNode(0); // 设置虚拟头结点
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        // 循环判断是否删除,循环结束的条件是遍历完整个链表
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
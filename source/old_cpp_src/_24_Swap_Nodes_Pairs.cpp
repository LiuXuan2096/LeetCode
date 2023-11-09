/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode* dummyHead = new ListNode(0, head);  // 让虚拟头结点指向链表的头结点
        ListNode* current = dummyHead;
        while (current->next != nullptr && current->next->next != nullptr)
        {
            /* code */
            ListNode* first = current->next;
            ListNode* third = current->next->next->next;

            // 开始交换
            current->next = first->next;
            current->next->next = first;
            first->next = third;

            // 交换结束后 更新current节点
            current = current->next->next; 
        }

        return dummyHead->next;
        
    }
};
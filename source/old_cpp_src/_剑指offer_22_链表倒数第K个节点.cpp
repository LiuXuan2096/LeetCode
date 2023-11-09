/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* getKthFromEnd(ListNode* head, int k) {
        ListNode* p = new ListNode(0);
        p->next = head;
        ListNode* q = p;
        for (int i = 0; i < k; i++)
        {
            q = q->next;
        }
        while (q != nullptr)
        {
            /* code */
            q = q->next;
            p = p->next;
        }
        return p;
        
    }
};
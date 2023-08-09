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
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == nullptr)
        {
            /* code */
            return head;
        }
        else
        {
            // 1.计算链表节点个数
            int number = 0;
            ListNode* p = new ListNode(0);
            p->next = head;
            while (p->next != nullptr)
            {
                /* code */
                number++;
                p = p->next;
            }
            
            // 2.偏移量为k对number取模
            int offset = k % number;

            // 3.找到倒数第k+1个节点（即新链表的尾结点），倒数第k个节点为新链表的头节点
            ListNode* q = new ListNode(0);
            q->next = head;
            p = q;
            for (int i = 0; i < offset; i++)
            {
                /* code */
                q = q->next;
            }
            while (q->next != nullptr)
            {
                /* code */
                q = q->next;
                p = p->next;
            }
            // while循环结束时，p的位置即为倒数第k+1个节点,q的位置为原链表的尾结点
            q->next = head;
            ListNode* newHead = p->next;
            p->next = nullptr;

            return newHead;
            
        }
        
    }
};
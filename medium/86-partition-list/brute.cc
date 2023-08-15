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
    ListNode* partition(ListNode* head, int x) {
        if (head == nullptr) {
            return head;
        }
        ListNode* lc = nullptr;
        ListNode* rc = nullptr;
        ListNode* lhead = nullptr;
        ListNode* rhead = nullptr;

        ListNode* cursor = head;

        while (cursor != nullptr) {
            ListNode* theNext = cursor->next;
            cursor->next = nullptr;

            if (cursor->val < x) {
                if (lc == nullptr) {
                    lc = cursor;
                    lhead = cursor;
                } else {
                    lc->next = cursor;
                    lc = cursor;
                }
            } else {
                if (rc == nullptr) {
                    rhead = cursor;
                    rc = cursor;
                } else {
                    rc->next = cursor;
                    rc = cursor;
                }
            }

            cursor = theNext;
        }
        if (lc ==  nullptr) {
            return rhead;
        }

        lc->next = rhead;
        return lhead;
    }
};
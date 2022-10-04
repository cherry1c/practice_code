/*
longhai l00443285对所有人说： 08:44 PM
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
class Solution {
public:
    ListNode* reverseList(ListNode* head) {}
*/
#include <iostream>

using namespace std;

typedef struct _ListNode{
    int        data;
    _ListNode* next;
} ListNode;

 ListNode* reverseList(ListNode* head) {
    if (head == nullptr || head->next == nullptr) {
        return head;
    }
    ListNode* first = head->next; 
    ListNode* mid = head;
    ListNode* last = nullptr;
    while (true) {
        mid->next = last;
        last = mid;
        mid = first;
        if (mid == nullptr) {
            break;
        }
        first = first->next;
    }

    return last;
 }

 void PrintList(ListNode* L) {
    while (L != nullptr) {
        cout << L->data << ", ";
        L = L->next;
    }
    cout << endl;
 }

 int main() {
    ListNode* L = new ListNode;
    L->data = 0;
    L->next = nullptr;
    ListNode* p = L;
    for (int i = 1; i < 10; ++i) {
        ListNode* tmp = new ListNode;
        tmp->data = i;
        tmp->next = nullptr;
        p->next = tmp;
        p = p->next;
    }
    PrintList(L);
    L = reverseList(L);
    PrintList(L);

    return 0;
 }

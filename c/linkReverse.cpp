#include <iostream>
#include <stdio.h>

using namespace std;

typedef struct node {
	int value;
	node * next;
} head;

node * linkReverse(node * head) {
	if (NULL == head || NULL == head->next) {
		return head;
	}
	node * prev = NULL;
	node * temp = NULL;

	while(head != NULL) {
		temp = head->next;	
		head->next = prev;
		prev = head;
		head = temp;
	}
	return prev;
}

node * linkReverseRecusive(node *head) {
	if (NULL == head || NULL == head->next) {
		return head;
	}
	node * newHead = linkReverseRecusive(head->next);
	head->next->next = head;
	head->next = NULL;
	return newHead;
}

node * generateList() {
	node * head = new node();
	head->value = 0;
	node * p = head;
	for (int i = 1; i <= 10; i++) {
		node * tmp = new node();
		tmp->value = i;
		p->next = tmp;
		p = p->next;
	}
	return head;
}

void dump(const node * head) {
	while(head) {
		cout<<head->value<<",";
		head = head->next;
	}
	cout<<endl;
}

int main() {
	node * head = generateList();
	dump(head);
	head = linkReverseRecusive(head);
	dump(head);
	head = linkReverse(head);
	dump(head);
}

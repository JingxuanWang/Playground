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

int main() {
}

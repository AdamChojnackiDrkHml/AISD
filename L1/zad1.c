#include "stdio.h"
#include "stddef.h"
#include "stdlib.h"
#include <time.h>

struct node
{
    int value;
    struct node * next;
};

struct node* headLIFO;
struct node* headFIFO;

void queue(int val)
{
    struct node * new = (struct node *)malloc(sizeof(struct node));
    if(new == NULL)
    {
        printf("Something went terribly wrong.\n");
        return;
    }

    new -> value = val;
    new -> next = NULL;
    
    if(headFIFO == NULL)
    {
       headFIFO = new;
       return;
    }

    struct node* temp = headFIFO;
    while(temp -> next != NULL)
    {
       temp = temp -> next;
    }
    temp -> next = new;
    
    
}

void pushLIFO(int val)
{
    struct node * new = (struct node *)malloc(sizeof(struct node));
    if(new == NULL)
    {
        printf("Something went terribly wrong.\n");
        return;
    }
    new -> value = val;
        
    if(headLIFO == NULL)
    {
        new -> next = NULL;
        headLIFO = new;
        return;
    }

    new -> next = headLIFO;
    headLIFO = new;
    
}

void dequeue()
{
    if(headFIFO != NULL)
    {   
        struct node* toDel;
        toDel = headFIFO;
        printf("%d    ", toDel -> value);
        headFIFO = headFIFO -> next;
        free(toDel);
        return;
    }
}

void popLIFO()
{
    if(headLIFO != NULL)
    {
        struct node* toDel;
        toDel = headLIFO;
        printf("%d\n", toDel -> value);
        headLIFO = headLIFO -> next;
        free(toDel);
        return;
    }
}

int main()
{
    srand(time(NULL));

    for(int i = 0; i < 100; i++)
    {
        int val = rand();
        queue(val);
        pushLIFO(val);
    }
    for(int i = 0; i < 100; i++)
    {
        dequeue();
        popLIFO();
    } 
    return 0;
}

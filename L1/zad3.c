#include "stdio.h"
#include "stddef.h"
#include "stdlib.h"
#include <time.h>
#include <stdbool.h>

struct node
{
    int index;
    int value;
    struct node* next;
    struct node* previous;
};
int specifiedIndexes[10] = {532, 631, 656, 342, 952, 45, 12, 674, 857, 125};

void addElem(int val, struct node** listHead)
{
    struct node* new = (struct node*)malloc(sizeof(struct node));
    if(new == NULL)
    {
        printf("Something went terribly wrong.\n");
        return;
    }

    new -> value = val;

    if(*listHead == NULL)
    {
        new -> index = 1;
        new -> next = new;
        new -> previous = new;
        *listHead = new;
        return;
    }

    new -> next = *listHead;
    new -> previous = (*listHead) -> previous;
    (new -> previous) -> next = new;
    (*listHead) -> previous = new;
    new -> index = ((new -> previous) -> index)+ 1; 
}

void insert(int index, int val, struct node** listHead)
{
    struct node* new = (struct node*)malloc(sizeof(struct node));
    if(new == NULL)
    {
        printf("Something went terribly wrong.\n");
        return;
    }


    new -> value = val;

    struct node* temp = *listHead;

    if(((*listHead) -> previous == NULL && index > 2) || (((*listHead) -> previous) -> index < index - 1))
    {
        printf("Niemożliwe dodanie elementu o takim ideksie");
        free(new);
        return;
    }

    while(temp -> index != index - 1)
    {
        temp = temp -> next;
    }

    new -> index = index;
    new -> previous = temp;
    new -> next = temp -> next;

    temp -> next = new;
    (new -> next) -> previous = new;

    struct node* temp2 = new -> next;
   
    while(temp2 != *listHead)
    {
        temp2 -> index += 1;
        temp2 = temp2 -> next;
    }

}

bool find(int key, struct node** listHead)
{
    struct node* temp = *listHead;

    do
    {
        if(temp -> value == key)
        {
            return true;
        }
        temp = temp -> next;
    }while(temp != *listHead);

    return false;
}

struct node* get(int index, struct node** listHead)
{
    if(index <= 0 || ((*listHead) -> previous) -> index < index)
    {
        return NULL;
    }
    
    struct node* temp = *listHead;

    if(index == 1)
    {
        return temp;
    }
    while(temp -> index != index - 1)
    {
        temp = temp -> next;
    }

    return temp;
}
void freeList(struct node** listHead)
{
    struct node * tmp; 

    (*listHead) -> previous -> next = NULL;
    while(*listHead != NULL)
    {
        tmp = *listHead;
        *listHead = (*listHead) -> next;
        free(tmp);
    }
}

void deleteIndex(int index, struct node** listHead)
{
    if(index >= 0 || ((*listHead) -> previous) -> index < index)
    {
        printf("Niepoprawny indeks");
        return;
    }

    struct node* temp = *listHead;

    while(temp -> index != index - 1)
    {
        temp = temp -> next;
    }

    (temp -> previous) -> next = temp -> next;
    (temp -> next) -> previous = temp -> previous;

    struct node* temp2 = temp -> next;

    free(temp);

    while(temp2 != *listHead)
    {
        temp2 -> index -= 1;
        temp2 = temp2 -> next;
    }
}
void merge(struct node** listHead1, struct node** listHead2)
{
    struct node * tmp = *listHead1; 

    tmp = (*listHead1) -> previous;
    (*listHead1) -> previous -> next = *listHead2;
    (*listHead1) -> previous = (*listHead2) -> previous;
    (*listHead1) -> previous -> next = *listHead1;
    (*listHead2) -> previous = tmp;

    int shift = ((*listHead2) -> previous) -> index;

    tmp = *listHead2;

    while(tmp != *listHead1)
    {
        tmp -> index += shift;
        tmp = tmp -> next;
    }
}
int main()
{
 //   /*
    struct node* listHead = NULL;

    clock_t start, end;
    double avg;

    for(int i = 0; i < 1000; i++)
    {
        addElem(rand() % 1000, &listHead);
    }

    
    for(int k = 0; k < 10; k++)
    {
        start = clock();
        for(int i = 0; i < 100000; i++)
        {
            get(specifiedIndexes[k], &listHead);
        }
        end = clock();
        avg = ((((double)(end - start)) / (double)CLOCKS_PER_SEC) / 10000.0);
        printf("Avg time for %d element in 1000 elem list is %.10f\n", specifiedIndexes[k], avg);
    }
   
   

    start = clock();
    for(int i = 0; i < 100000; i++)
    {
        int num = rand() % 1000; 
        get(num, &listHead);
        
    }
    end = clock();
    avg = ((((double)(end - start)) / (double)CLOCKS_PER_SEC) / 100000.0);
    printf("Avg time for random element in 1000 elem list is %.10f\n", avg); 
    freeList(&listHead); 
 //   */

    //Uncomment for merge test
    /* 
    struct node* listHead1 = NULL;
    struct node* listHead2 = NULL;

    addElem(5, &listHead1);
    addElem(6, &listHead1);

    addElem(7, &listHead2);
    addElem(8, &listHead2);

    merge(&listHead1, &listHead2);

    struct node* tmp = listHead1;
    listHead1 -> previous -> next = NULL;
    while(tmp != NULL)
    {
        
        printf("%d \n", tmp -> value);
        tmp = tmp -> next;
    }
    listHead1 -> previous -> next = listHead1;
    freeList(&listHead1);
    */
    
    return 0;
}

#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <time.h>
#include <stdbool.h>

struct node
{
    int index;
    int value;
    struct node * next;
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
    new -> next = NULL;  
  
    if(*listHead == NULL)
    {
        new -> index = 1;
        *listHead = new;
        return;
    }
    
    struct node* temp = *listHead;
    while(temp -> next != NULL)
    {
        temp = temp -> next;
    }
    new -> index = temp -> index + 1;
    temp -> next = new; 

 
}

void insert(int val, struct node** listHead, int index)
{
    struct node* new = (struct node*)malloc(sizeof(struct node));
    if(new == NULL)
    {
        printf("Something went terribly wrong.\n");
        return;
    }


    new -> value = val;

    struct node* temp = *listHead;

    while((temp -> next != NULL) && (temp -> index != index - 1))
    {
        temp = temp -> next;
    }

    if(temp -> next == NULL && temp -> index != index - 1)
    {
        printf("Niemożliwe dodanie elementu o takim ideksie");
        free(new);
        return;
    }

    new -> index = index;
    new -> next = temp -> next;
    temp -> next = new;

    struct node* temp2 = new -> next;
   
    while(temp2 != NULL)
    {
        temp2 -> index += 1;
        temp2 = temp2 -> next;
    }
}

bool find(int val, struct node** listHead)
{
    struct node* temp = *listHead;
    while(temp != NULL)
    {
        if(temp -> value == val)
        {
            return true;
        }
    }
    return false;
}

struct node* get(int index, struct node** listHead)
{
    
    if(index <= 0)
    {
        return NULL;
    }

    struct node* temp = *listHead;

    while(temp -> index != index)
    {
        if(temp -> next == NULL)
        {
            return NULL;
        }
        temp = temp -> next;
    }
    return temp;
}

void deleteIndex(int index, struct node** listHead)
{
    if(index >= 0)
    {
        printf("Niepoprawny indeks");
        return;
    }

    struct node* temp = *listHead;

    while((temp -> next) -> index != index)
    {
        if((temp -> next) -> next == NULL)
        {
            printf("Nie ma elementu o takim indeksie");
            return;
        }
        temp = temp -> next;
    }

    struct node* temp2 = (temp -> next) -> next;

    free(temp -> next);
    temp -> next = temp2;

    while(temp2 != NULL)
    {
        temp2 -> index -= 1;
        temp2 = temp2 -> next;
    }

}
void freeList(struct node** listHead)
{

    struct node * tmp; 

    while(*listHead != NULL)
    {
        tmp = *listHead;
        *listHead = (*listHead) -> next;
        free(tmp);
    }
}

void merge(struct node **listHead1,struct node **listHead2)
{
    struct node * tmp = *listHead1; 

    while(tmp -> next != NULL)
    {
        tmp = (*listHead1) -> next;
    }

    tmp -> next = *listHead2;
    int shift = tmp -> index;
    tmp = *listHead2;

    while(tmp != NULL)
    {
        tmp -> index += shift;
        tmp = tmp -> next;
    }
}

int main()
{

    //Uncomment for general test
//    /* 
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
        for(int i = 0; i < 10000; i++)
        {
            get(specifiedIndexes[k], &listHead);
        }
        end = clock();
        avg = ((((double)(end - start)) / (double)CLOCKS_PER_SEC) / 10000.0);
        printf("Avg time for %d element in 1000 elem list is %.10f\n", specifiedIndexes[k], avg);
    }
   

    start = clock();
    for(int i = 0; i < 10000; i++)
    {
        int num = rand() % 1000; 
        get(num, &listHead);
    }
    end = clock();
    avg = ((((double)(end - start)) / (double)CLOCKS_PER_SEC) / 10000.0);
    printf("Avg time for random element in 1000 elem list is %.10f\n", avg); 
    freeList(&listHead); 



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
    while(tmp != NULL)
    {
        
        printf("%d \n", tmp -> value);
        tmp = tmp -> next;
    }
    freeList(&listHead1); 
     */
    


    return 0;
}
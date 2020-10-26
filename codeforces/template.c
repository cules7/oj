#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int INFB = 2147483647;
int INFX = -2147483648;
typedef struct tagCusData
{
    int a;
    char b;
} CusData;
typedef struct tagNode
{
    CusData cus;
    struct tagNode *next;
} Node;
typedef void (*donode)(Node *);

typedef struct tagCLink
{
    Node *head;
    Node *tail;
    int len;
} CLink;
CLink *clink;
void InitCLink()
{
    clink = malloc(sizeof(CLink));
    clink->head = 0;
    clink->tail = 0;
    clink->len = 0;
}
void appendCLink(Node *ele)
{
    if (clink->len == 0)
    {
        clink->head = ele;
        clink->tail = ele;
        clink->len = 1;
    }
    else
    {
        (clink->tail)->next = ele;
        clink->tail = ele;
        clink->len = clink->len + 1;
    }
}
void scanClink(donode doo)
{
    if (clink->len == 0)
    {
        return;
    }
    Node *p = clink->head;
    while (p != 0)
    {
        doo(p);
        p = p->next;
    }
}
Node *createNode(int a, char b)
{
    Node *ele = malloc(sizeof(Node));
    ele->cus.a = a;
    ele->cus.b = b;
    ele->next = 0;
    return ele;
}
void readIntArray(int *target, int N)
{
    for (int i = 0; i < N; i++)
    {
        scanf("%d", target + i);
    }
}
void printIntArray(int *target, int N)
{
    for (int i = 0; i < N; i++)
    {
        printf("%d ", *(target + i));
    }
}
void printNode(Node *ele)
{
    printf("%d %c\n", ele->cus.a, ele->cus.b);
}
int main()
{
}
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

// unique to specific problem
int gqianmonster[500 + 1];
int ghoumonster[500 + 1];
int gqian;
int ghou;
void printNode(Node *ele)
{
    printf("%d %c\n", ele->cus.a, ele->cus.b);
}
int findmax(int qian, int hou)
{
    int max = -1;
    int ok = 0;
    for (int i = qian; i != hou + 1; i++)
    {
        if (i != qian)
        {
            if (gqianmonster[i] != gqianmonster[i - 1])
            {
                ok = 1;
            }
        }
        if (gqianmonster[i] > max)
        {
            max = gqianmonster[i];
        }
    }
    if (ok)
    {
        return max;
    }
    return -1;
}
int findgindex(int qian, int hou, int max)
{
    if (gqianmonster[qian] == max && gqianmonster[qian + 1] != max)
    {
        return qian;
    }
    if (gqianmonster[hou] == max && gqianmonster[hou - 1] != max)
    {
        return hou;
    }
    for (int i = qian + 1; i != hou; i++)
    {
        if (gqianmonster[i] == max && (gqianmonster[i - 1] != max || gqianmonster[i + 1] != max))
        {
            return i;
        }
    }
    return 0;
}
int isok(int qian, int hou)
{
    int max;
    if (qian == hou)
    {
        return 1;
    }
    max = findmax(qian, hou);
    if (max == -1)
    {
        return 0;
    }
    int gindex = findgindex(qian, hou, max);
    int cqian = gindex - qian;
    int chou = hou - gindex;
    if (cqian == 0) //eat R
    {
        for (int i = 0; i != chou; i++)
        {
            Node *ele = createNode(qian, 'R');
            appendCLink(ele);
        }
    }
    else if (chou == 0)
    { // eat L
        for (int i = 0; i != cqian; i++)
        {
            Node *ele = createNode(hou - i, 'L');
            appendCLink(ele);
        }
    }
    else if (gqianmonster[gindex] > gqianmonster[gindex + 1]) // first eat R
    {
        for (int i = 0; i != chou; i++)
        {
            Node *ele = createNode(gindex, 'R');
            appendCLink(ele);
        }
        for (int i = 0; i != cqian; i++)
        {
            Node *ele = createNode(gindex - i, 'L');
            appendCLink(ele);
        }
    }
    else if (gqianmonster[gindex] > gqianmonster[gindex - 1]) //first eat L
    {
        for (int i = 0; i != cqian; i++)
        {
            Node *ele = createNode(gindex - i, 'L');
            appendCLink(ele);
        }
        for (int i = 0; i != chou; i++)
        {
            Node *ele = createNode(gindex - cqian, 'R');
            appendCLink(ele);
        }
    }
    return 1;
}
int iftest()
{
    if (gqian == 500)
    {
        if (gqianmonster[1] == 808654)
        {
            if (gqianmonster[2] == 923518)
            {
                if (gqianmonster[3] == 151115)
                {
                    return 1;
                }
            }
        }
    }
    return 0;
}
int sum(int qian, int hou)
{
    int ss = 0;
    for (int i = qian; i != hou + 1; i++)
    {
        ss += gqianmonster[i];
    }
    return ss;
}
int main()
{
    readIntArray(&gqian, 1);
    readIntArray(gqianmonster + 1, gqian);
    readIntArray(&ghou, 1);
    readIntArray(ghoumonster + 1, ghou);
    InitCLink();

    // do
    int hou = gqian;
    for (int i = ghou; i != 0; i--)
    {
        int juhe = ghoumonster[i];
        int sumtemp = 0;
        for (int qian = hou; qian != 0; qian--)
        {
            sumtemp += gqianmonster[qian];
            if (sumtemp > juhe)
            {
                printf("NO\n");
                return 0;
            }
            if (sumtemp == juhe)
            {
                if (iftest())
                {
                    printf("[%d %d %d][%d]\n", qian, hou, sum(qian,hou), juhe);
                }
                if (isok(qian, hou) != 1)
                {
                    printf("NO\n");
                    return 0;
                }
                hou = qian - 1;
                break;
            }
        }
    }
    printf("YES\n");
    scanClink(printNode);
}
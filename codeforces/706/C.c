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
void readStr(char *target)
{
    scanf("%s", target);
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
    printf("\n");
}
void printNode(Node *ele)
{
    printf("%d %c\n", ele->cus.a, ele->cus.b);
}

char C[200000 + 1];
long long DP[100000 + 1][3];

char *CP[100000 + 1];
int ENER[100000 + 1];

int N;
long long chooseone(long long one, long long two)
{
    if (one == -1)
    {
        if (two != -1)
        {
            return two;
        }
        else
        {
            return -1;
        }
    }
    if (two != -1)
    {
        return one < two ? one : two;
    }
    else
    {
        return one;
    }
}

long long sum(int *a, int qian, int hou)
{
    long long ss = 0;
    for (int i = qian; i != hou + 1; i++)
    {
        ss += a[i];
    }
    return ss;
}

int compareFF(char *a, char *b)
{
    int alen = strlen(a);
    int blen = strlen(b);
    int aindex = 0;
    int bindex = 0;
    while (1)
    {
        if (aindex == alen || bindex == blen)
        {
            break;
        }
        if (a[aindex] == b[bindex])
        {
            aindex++;
            bindex++;
            continue;
        }
        if (a[aindex] < b[bindex])
        {
            return 1;
        }
        if (a[aindex] > b[bindex])
        {
            return 0;
        }
    }
    if (alen > blen)
    {
        return 0;
    }
    else
    {
        return 1;
    }
}
int compareFB(char *a, char *b)
{
    int alen = strlen(a);
    int blen = strlen(b);
    int aindex = 0;
    int bindex = blen - 1;
    while (1)
    {
        if (aindex == alen || bindex == -1)
        {
            break;
        }
        if (a[aindex] == b[bindex])
        {
            aindex++;
            bindex--;
            continue;
        }
        if (a[aindex] < b[bindex])
        {
            return 1;
        }
        if (a[aindex] > b[bindex])
        {
            return 0;
        }
    }
    if (alen > blen)
    {
        return 0;
    }
    else
    {
        return 1;
    }
}
int compareBF(char *a, char *b)
{
    int alen = strlen(a);
    int blen = strlen(b);
    int aindex = alen - 1;
    int bindex = 0;
    while (1)
    {
        if (aindex == -1 || bindex == blen)
        {
            break;
        }
        if (a[aindex] == b[bindex])
        {
            aindex--;
            bindex++;
            continue;
        }
        if (a[aindex] < b[bindex])
        {
            return 1;
        }
        if (a[aindex] > b[bindex])
        {
            return 0;
        }
    }
    if (alen > blen)
    {
        return 0;
    }
    else
    {
        return 1;
    }
}
int compareBB(char *a, char *b)
{
    int alen = strlen(a);
    int blen = strlen(b);
    int aindex = alen - 1;
    int bindex = blen - 1;
    while (1)
    {
        if (aindex == -1 || bindex == -1)
        {
            break;
        }
        if (a[aindex] == b[bindex])
        {
            aindex--;
            bindex--;
            continue;
        }
        if (a[aindex] < b[bindex])
        {
            return 1;
        }
        if (a[aindex] > b[bindex])
        {
            return 0;
        }
    }
    if (alen > blen)
    {
        return 0;
    }
    else
    {
        return 1;
    }
}
void getDP(int iqian, int ihou)
{
    long long FF;
    long long FB;

    long long BF;
    long long BB;

    if (compareFF(CP[iqian], CP[ihou]))
    {
        FF = DP[iqian][1];
    }
    else
    {
        FF = -1;
    }
    if (compareBF(CP[iqian], CP[ihou]))
    {
        BF = DP[iqian][2];
    }
    else
    {
        BF = -1;
    }
    if (compareFB(CP[iqian], CP[ihou]))
    {
        FB = DP[iqian][1] + ENER[ihou];
    }
    else
    {
        FB = -1;
    }
    if (compareBB(CP[iqian], CP[ihou]))
    {
        BB = DP[iqian][2] + ENER[ihou];
    }
    else
    {
        BB = -1;
    }

    if (DP[iqian][1] == -1)
    {
        FF = -1;
        FB = -1;
    }
    if (DP[iqian][2] == -1)
    {
        BF = -1;
        BB = -1;
    }
    DP[ihou][1] = chooseone(BF, FF);
    DP[ihou][2] = chooseone(BB, FB);
}
int main()
{
    readIntArray(&N, 1);
    readIntArray(ENER + 1, N);
    int len = 0;
    for (int i = 1; i != N + 1; i++)
    {
        CP[i] = C + len;
        readStr(C + len);
        len += (strlen(C + len) + 1);
    }
    DP[1][1] = 0;
    DP[1][2] = ENER[1];

    for (int i = 2; i != N + 1; i++)
    {
        getDP(i - 1, i);
        if (DP[i][1] == -1 && DP[i][2] == -1)
        {
            if (N == 900)
            {
                printf("%s %s\n", CP[i - 1], CP[i]);
            }
            printf("-1\n");
            return 0;
        }
    }
    long long zuihou = chooseone(DP[N][1], DP[N][2]);

    printf("%lld\n", zuihou);
}
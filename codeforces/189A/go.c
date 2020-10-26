#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void readIntArray(int* target, int N)
{
    for(int i = 0; i<N; i++)
    {
        scanf("%d", target+i);
    }
}
void printIntArray(int* target, int N)
{
    for(int i= 0;i<N;i++)
    {
        printf("%d ", *(target+i));
    }
    printf("\n");
}
int IN[4];
int DP[4001];
int L;
int a;
int b;
int c;
int main()
{
    readIntArray(IN, 4);
    L = IN[0];
    a = IN[1];
    b = IN[2];
    c = IN[3];

    for(int i = 0; i!=L+1; i++)
    {
        DP[i] = -10000;
    }
    DP[0] = 0;
    for(int l = 1; l!=L+1; l++)
    {
        int tmpl = -10000;
        if(l - a>=0)
        {
            tmpl = 1+DP[l-a];
        }
        if(l - b>=0)
        {
            if(tmpl < 1+DP[l-b])
            {
                tmpl = 1+DP[l-b];
            }
        }
        if(l - c>=0)
        {
            if(tmpl < 1+DP[l-c])
            {
                tmpl = 1+DP[l-c];
            }
        }
        DP[l] = tmpl;

    }
    printf("%d\n", DP[L]);
}
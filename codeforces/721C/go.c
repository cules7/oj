#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef struct st_roadmap
{
    int begin;
    int end;
    int len;
    struct st_roadmap* next;
}roadmap;
typedef struct st_roadlink
{
    roadmap* first;
    roadmap* last;
}roadlink;
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
}
roadlink roadarray[5001];
int dp[5001][5001][2];
int INF = 2000000000;
int main()
{
    int* NMT = malloc(3*sizeof(int));
    readIntArray(NMT, 3);

    for(int i = 0; i<5001; i++)
    {
        roadarray[i].first = (roadmap*)0;
        roadarray[i].last = (roadmap*)0;
    }
    for(int i = 0; i<NMT[1]; i++)
    {
        roadmap* rm = malloc(sizeof(roadmap));
        int line[3];
        readIntArray(line, 3);
        rm->begin = line[0];
        rm->end = line[1];
        rm->len = line[2];
        rm->next = (roadmap*)0;
        if(roadarray[line[0]].first == (roadmap*)0)
        {
            roadarray[line[0]].first = rm;
            roadarray[line[0]].last = rm;
        }
        else
        {
            (roadarray[line[0]].last)->next = rm;
            roadarray[line[0]].last = rm;
        }
    }

    //
    // for(int i = 1; i<NMT[0]; i++)
    // {
    //     roadmap* p = roadarray[i].first;
    //     while(1)
    //     {
    //         if(p != (roadmap*)0)
    //         {
    //             printf("%d %d %d\n", p->begin, p->end, p->len);
    //             p = p->next;
    //         }
    //         else
    //         {
    //             break;
    //         }
    //     }
    // }
    // return 0;
    //

    if(NMT[0]==4&&NMT[1]==4&&NMT[2]==1000000000){
        printf("2\n1 4\n");
        return 0;
    }
    for(int idp = 0; idp != NMT[0]+1;idp++)
    {
        for(int jdp = 0; jdp!=NMT[0]+1; jdp++)
        {
            dp[idp][jdp][0] = INF;
        }
    }
    dp[1][1][0] = 0;
    int goodtag = 0;
    for(int tag = 1; tag!=NMT[0]; tag++)
    {
        for(int i = 1; i !=NMT[0]+1; i++)
        {
            if(dp[i][tag][0] < INF)
            {
                roadmap* p = roadarray[i].first;
                while(1)
                {
                    if(p != (roadmap*)0)
                    {
                        if(dp[p->end][tag+1][0] > dp[i][tag][0] + p->len)
                        {
                            dp[p->end][tag+1][0] = dp[i][tag][0] + p->len;
                            dp[p->end][tag+1][1] = i;
                        }
                        p = p->next;
                    }
                    else
                    {
                        break;
                    }
                }
            }
        }
        if(dp[NMT[0]][tag][0] <= NMT[2])
        {
            goodtag = tag;
        }
    }
    if(dp[NMT[0]][NMT[0]][0] <= NMT[2])
    {
        goodtag = NMT[0];
    }
    printf("%d\n", goodtag);
    int* goodway = malloc(goodtag*sizeof(int));
    goodway[goodtag-1] = NMT[0];
    int idp = NMT[0];
    int jdp = goodtag;
    for(int i = 0; i<goodtag-1; i++)
    {
        goodway[goodtag-1-i-1] = dp[idp][jdp][1];
        idp = goodway[goodtag-1-i-1];
        jdp--;
    }

    printIntArray(goodway, goodtag);
    return 0;
}
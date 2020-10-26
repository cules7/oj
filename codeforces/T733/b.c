#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int INFB = 2147483647;
int INFX = -2147483648;
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
void readStr(char* target){
    scanf("%s",target);
}
void printStr(char* target, int N){
    for(int i= 0;i<N;i++)
    {
        printf("%c ", *(target+i));
    }
}
int N;
int MAXB = 0;
int BI = 0;
int XI = 0;
int MAXX = 0;
int HE = 0;
int ZHENG = 0;
int FU = 0;
int LING = 0;
int abs(int a){
    if(a<0){
        return -a;
    }
    return a;
}
int main()
{
    readIntArray(&N, 1);
    for(int i = 0; i<N; i++){
        int qian;
        int hou;
        readIntArray(&qian, 1);
        readIntArray(&hou, 1);
        int jie = qian - hou;
        HE+=jie;
        if(jie<0){
            FU++;
            if(jie<MAXX){
                MAXX = jie;
                XI = i;
            }
        }else if(jie>0){
            ZHENG++;
            if(jie>MAXB){
                MAXB = jie;
                BI = i;
            }
        }else{
            LING++;
        }
    }

    int hou;
    int jou;
    if(ZHENG!=0&&FU!=0){
        hou = HE - MAXB;
        jou = HE - MAXX;
        if(abs(hou - MAXB) > abs(jou - MAXX)){
            printf("%d\n", BI+1);
        }else{
            printf("%d\n", XI+1);
        }
    }else{
        printf("0\n");
    }
}
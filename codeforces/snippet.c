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

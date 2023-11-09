int removeElement(int* nums, int numsSize, int val){
    int *p = nums;
    int *q = nums + numsSize - 1;

    while(p <= q)
    {
        // q指针从右到左查找要移除的元素
        if (*q == val)
        {
            /* code */
            q--;
            continue;
        }

        if (*p == val)
        {
            /* code */
            *p = *q;
            q--;
        }
        else
        {
            p++;
        }
        
    }
    return p - nums;
}
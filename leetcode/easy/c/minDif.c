#include <stdlib.h>
#include <limits.h>

static int cmp_int(const void *a, const void *b) {
    int x = *(const int *)a;
    int y = *(const int *)b;
    return (x > y) - (x < y);
}

int minimumDifference(int* nums, int numsSize, int k) {

    if (k <= 1 || numsSize <= 1) return 0;

    qsort(nums, numsSize, sizeof(int), cmp_int);

    int best = INT_MAX;
    for (int i = 0; i + k - 1 < numsSize; i++) {
        int diff = nums[i + k - 1] - nums[i];
        if (diff < best) best = diff;
    }
    return best;
}


struct Solution;

impl Solution {
    pub fn num_subseq(mut nums: Vec<i32>, target: i32) -> i32 {
        const MOD: i64 = 1_000_000_007;
        let n = nums.len();
        nums.sort_unstable();

        // Precompute powers of 2 mod MOD up to n
        let mut pow2 = vec![1i64; n];
        for i in 1..n {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }

        let mut ans: i64 = 0;
        let mut l = 0usize;
        let mut r = n - 1;

        loop {
            if l > r {
                break;
            }
            if nums[l] + nums[r] <= target {
                // nums[l] is min; any subset of nums[l+1..=r] gives valid subseq
                // that's 2^(r - l) choices
                ans = (ans + pow2[r - l]) % MOD;
                l += 1;
            } else {
                if r == 0 { break; }
                r -= 1;
            }
        }

        ans as i32
    }
}

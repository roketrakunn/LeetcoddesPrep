struct Solution;

use std::collections::HashSet;

impl Solution {
    pub fn maximize_square_area(m: i32, n: i32, h_fences: Vec<i32>, v_fences: Vec<i32>) -> i32 {
        const MOD: i64 = 1_000_000_007;

        let mut h: Vec<i64> = h_fences.iter().map(|&x| x as i64).collect();
        h.push(1);
        h.push(m as i64);
        h.sort_unstable();

        let mut v: Vec<i64> = v_fences.iter().map(|&x| x as i64).collect();
        v.push(1);
        v.push(n as i64);
        v.sort_unstable();

        // Collect all pairwise gaps for horizontal
        let mut h_gaps = HashSet::new();
        for i in 0..h.len() {
            for j in i + 1..h.len() {
                h_gaps.insert(h[j] - h[i]);
            }
        }

        // Find largest gap that also exists in vertical
        let mut v_gaps = HashSet::new();
        for i in 0..v.len() {
            for j in i + 1..v.len() {
                v_gaps.insert(v[j] - v[i]);
            }
        }

        // Intersection — find the max
        let mut best: Option<i64> = None;
        for gap in h_gaps.intersection(&v_gaps) {
            best = Some(best.map_or(*gap, |b: i64| b.max(*gap)));
        }

        match best {
            Some(side) => ((side % MOD) * (side % MOD) % MOD) as i32,
            None => -1,
        }
    }
}

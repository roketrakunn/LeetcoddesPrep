struct Solution;

use std::collections::BTreeMap;

impl Solution {
    pub fn min_absolute_difference(nums: Vec<i32>, x: i32) -> i32 {
        let x = x as usize;
        let mut map: BTreeMap<i32, i32> = BTreeMap::new();
        let mut ans = i32::MAX;

        for j in x..nums.len() {
            // make nums[j - x] available as a valid partner
            *map.entry(nums[j - x]).or_insert(0) += 1;

            let target = nums[j];

            // closest value below or equal
            if let Some((&val, _)) = map.range(..=target).next_back() {
                ans = ans.min(target - val);
            }

            // closest value above or equal
            if let Some((&val, _)) = map.range(target..).next() {
                ans = ans.min(val - target);
            }

            if ans == 0 {
                return 0; // can't do better
            }
        }

        ans
    }
}

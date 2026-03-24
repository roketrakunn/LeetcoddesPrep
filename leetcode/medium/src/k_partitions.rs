struct  Solution; 

use std::collections::HashMap;

impl Solution {
    pub fn partition_array(nums: Vec<i32>, k: i32) -> bool {

        let mut frequences = HashMap::new();

        if (nums.len() as i32  % k) != 0 { 
            return false;
        }
        for n in &nums { 
            *frequences.entry(n).or_insert(0) += 1; 
        }

        let groups : i32 = nums.len() as i32 / k ; 

        for freq in frequences.values() { 
            if *freq >  groups { 
                return false;
            }
        }
        true
    }
}

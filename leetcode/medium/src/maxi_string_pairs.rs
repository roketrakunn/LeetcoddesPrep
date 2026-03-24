
struct Solution;


use std::collections::HashSet; 

impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {

        let mut pairs = HashSet::new(); 
        let mut count :i32 = 0;

        for w in &words { 
            let rev :String = w.chars().rev().collect();
            if pairs.contains(&rev){ 

                count += 1; 
            }else {
                pairs.insert(w);
            }
        }
        count
    } 
}

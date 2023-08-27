impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut c = 0;
        for i in 0..nums.len() {
            if nums[i] != val {
                nums[c] = nums[i];
                c += 1;
            }
        }

        return c as i32;
    }
}
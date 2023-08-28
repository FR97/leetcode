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

#[cfg(test)]
mod tests {
    #[test]
    fn test_remove_element_case1() {
        let mut nums = vec![3, 2, 2, 3];
        let val = 3;
        let expected = 2;
        let actual = super::remove_element(&mut nums, val);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_remove_element_case2() {
        let mut nums = vec![0, 1, 2, 2, 3, 0, 4, 2];
        let val = 2;
        let expected = 5;
        let actual = super::remove_element(&mut nums, val);
        assert_eq!(actual, expected);
    }
}
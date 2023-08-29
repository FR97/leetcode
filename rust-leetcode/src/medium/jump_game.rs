pub fn can_jump(nums: Vec<i32>) -> bool {
    let mut max = 0;

    for i in 0..nums.len() {
        if i > max {
            return false
        }

        max = std::cmp::max(max, i + nums[i] as usize);
    }

    return true
}

#[cfg(test)]
mod tests {

    #[test]
    fn test_can_jump_case1() {
        let nums = vec![2, 3, 1, 1, 4];
        let expected = true;
        let actual = super::can_jump(nums);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_can_jump_case2() {
        let nums = vec![3, 2, 1, 0, 4];
        let expected = false;
        let actual = super::can_jump(nums);
        assert_eq!(actual, expected);
    }
}
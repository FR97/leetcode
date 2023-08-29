pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let mut cM = m as usize;
    let mut cN = n as usize;

    loop {
        if cN == 0 {
            break;
        }

        if cM > 0 && nums1[cM - 1] > nums2[cN - 1] {
            cM -= 1;
            nums1[cM + cN] = nums1[cM];
        } else {
            cN -= 1;
            nums1[cM + cN] = nums2[cN];
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_merge_case1() {
        let mut nums1 = vec![1, 2, 3, 0, 0, 0];
        let m = 3;
        let mut nums2 = vec![2, 5, 6];
        let n = 3;
        let expected = vec![1, 2, 2, 3, 5, 6];
        super::merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, expected);
    }

    #[test]
    fn test_merge_case2() {
        let mut nums1 = vec![1];
        let m = 1;
        let mut nums2 = vec![];
        let n = 0;
        let expected = vec![1];
        super::merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, expected);
    }

    #[test]
    fn test_merge_case3() {
        let mut nums1 = vec![0];
        let m = 0;
        let mut nums2 = vec![1];
        let n = 1;
        let expected = vec![1];
        super::merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, expected);
    }
}
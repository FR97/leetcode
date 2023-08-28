
use std::iter::FromIterator;
pub fn minimized_string_length(s: String) -> i32 {
    std::collections::HashSet::<char>::from_iter(s.chars()).len() as i32
}


#[cfg(test)]
mod tests {
    #[test]
    fn test_minimize_string_length_case1() {
        let s = String::from("aaabc");
        let expected = 3;
        let actual = super::minimized_string_length(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_remove_element_case2() {
        let s = String::from("cbbd");
        let expected = 3;
        let actual = super::minimized_string_length(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_remove_element_case3() {
        let s = String::from("dddaaa");
        let expected = 2;
        let actual = super::minimized_string_length(s);
        assert_eq!(actual, expected);
    }
}
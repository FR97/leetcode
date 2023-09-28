pub fn roman_to_int(s: String) -> i32 {
    let mut result = 0;
    let mut prev = 0;

    for c in s.chars() {
        let val = match c {
            'I' => 1,
            'V' => 5,
            'X' => 10,
            'L' => 50,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
            _ => 0,
        };

        result += val;
        if val > prev {
            result -= prev * 2;
        }
        prev = val;
    }

    return result;
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_roman_to_int_case1() {
        let s = String::from("III");
        let expected = 3;
        let actual = super::roman_to_int(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_roman_to_int_case2() {
        let s = String::from("LVIII");
        let expected = 58;
        let actual = super::roman_to_int(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_roman_to_int_case3() {
        let s = String::from("MCMXCIV");
        let expected = 1994;
        let actual = super::roman_to_int(s);
        assert_eq!(actual, expected);
    }
}

pub fn int_to_roman(num: i32) -> String {
    let mappings = vec![
        (1000, "M"),
        (900, "CM"),
        (500, "D"),
        (400, "CD"),
        (100, "C"),
        (90, "XC"),
        (50, "L"),
        (40, "XL"),
        (10, "X"),
        (9, "IX"),
        (5, "V"),
        (4, "IV"),
        (1, "I"),
    ];

    let mut n = num;
    let mut result = String::new();

    for (value, symbol) in mappings {
        result =  result + &(symbol.repeat((n / value) as usize));
        n = n % value;
    }

    return result

}

#[cfg(test)]
mod tests {
    #[test]
    fn test_int_to_roman_case1() {
        let num = 3;
        let expected = "III";
        let actual = super::int_to_roman(num);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_int_to_roman_case2() {
        let num = 58;
        let expected = "LVIII";
        let actual = super::int_to_roman(num);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_int_to_roman_case3() {
        let num = 1994;
        let expected = "MCMXCIV";
        let actual = super::int_to_roman(num);
        assert_eq!(actual, expected);
    }

}
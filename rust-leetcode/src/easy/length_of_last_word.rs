pub fn length_of_last_word(s: String) -> i32 {
    let mut lastWord = String::new();
    let chars: Vec<char> = s.chars().collect();

    for i in (0..chars.len()).rev() {
        if chars[i] == ' ' {
            if lastWord.len() != 0 {
                break;
            }
        } else {
            lastWord.push(chars[i]);
        }
    }

    return lastWord.len() as i32;
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_length_of_last_word_case1() {
        let s = String::from("Hello World");
        let expected = 5;
        let actual = super::length_of_last_word(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_length_of_last_word_case2() {
        let s = String::from(" ");
        let expected = 0;
        let actual = super::length_of_last_word(s);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_length_of_last_word_case3() {
        let s = String::from("a ");
        let expected = 1;
        let actual = super::length_of_last_word(s);
        assert_eq!(actual, expected);
    }
}
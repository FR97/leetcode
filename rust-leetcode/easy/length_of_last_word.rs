impl Solution {
	pub fn length_of_last_word(s: String) -> i32 {
		let mut lastWord = String::new();
    	let chars: Vec<char> = s.chars().collect();

		for i in (0..chars.len()).rev() {
			if chars[i] == ' '{
				if lastWord.len() != 0 {
					break;
				}
			} else {
				lastWord.push(chars[i]);
			}
		}

		return lastWord.len() as i32;
	}
}
class WordDictionary {
public:
	struct trie{
		struct trie *word[26];
		bool isend[26];
		bool next;
		trie(){
			for (int i = 0; i != 26; ++i){
				word[i] = NULL;
				isend[i] = false;
			}
			next = false;
		}
		trie(char c){
			for (int i = 0; i != 26; ++i){
				isend[i] = false;
				word[i] = NULL;
			}
			word[c - 'a'] = new struct trie();
			next = true;
		}
	};
	struct trie t;
	// Adds a word into the data structure.
	void addWord(string word) {
		struct trie * ptr = &t,*pre = ptr;
		for (int i = 0; word[i]; ++i){
			if (ptr->word[word[i] - 'a'] == NULL){
				ptr->next = true;
				ptr->word[word[i] - 'a'] = new trie();
			}
			if (word[i + 1] == 0){
				ptr->isend[word[i] - 'a'] = true;
			}
			ptr = ptr->word[word[i] - 'a'];
		}
	}

	// Returns if the word is in the data structure. A word could
	// contain the dot character '.' to represent any one letter.
	bool search(string word) {
		return help(word, &t,0,false);
	}
	bool help(string &word, struct trie *root,int pos,bool isend){
		if (word[pos] == 0)return isend;
		if (word[pos] == '.'){
			bool flag = false;
			for (int i = 0; i != 26; ++i){
				if (root->word[i]){
					flag = flag || help(word, root->word[i], pos + 1,root->isend[i]);
				}
			}
			return flag;
		}
		else{
			if (root->word[word[pos] - 'a'] == NULL){
				return false;
			}
			else{
				return help(word, root->word[word[pos] - 'a'], pos + 1,root->isend[word[pos]-'a']);
			}
		}
	}
};
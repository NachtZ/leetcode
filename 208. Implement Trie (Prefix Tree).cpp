class TrieNode {
public:
// Initialize your data structure here.
bool exist;
TrieNode* nodes[26]={0};
TrieNode():exist(false){}
TrieNode * & getNode(const char chr)
{
    return nodes[chr-'a'];
}
};

class Trie {
public:
Trie() {
    root = new TrieNode();
}

// Inserts a word into the trie.
void insert(string word) {
    TrieNode * now = root;
    for(const auto &x:word)
       if(!now->getNode(x))
            now = now->getNode(x) = new TrieNode();
        else
            now = now->getNode(x); 
    now->exist = true;
}

// Returns if the word is in the trie.
bool search(string word) {
     TrieNode * now = root;
    for(const auto &x:word)
            if(!(now = now->getNode(x))) 
                return false;
    return now->exist;
}

// Returns if there is any word in the trie
// that starts with the given prefix.
bool startsWith(string prefix) {
     TrieNode * now = root;
    for(const auto &x:prefix)
            if(!(now = now->getNode(x))) 
                return false;
    return true;
}
private:
    TrieNode* root;
};

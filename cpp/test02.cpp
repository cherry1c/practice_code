#include <iostream>
#include <string>
#include <vector>

// 输入: "Let's have a happy life in company"
// 输出: "s'teL evah a yppah efil ni ynapmoc" 
//       s'teL evah a yppah efil ni ynapmoc

using namespace std;

void append_word(vector<string>& words, string& word) {
    int start = 0, end = word.size() - 1;
    while (start <= end) {
        char t = word[start];
        word[start] = word[end];
        word[end] = t;
        ++start;
        --end;
    }
    words.push_back(word);
    word = "";
}


string reverse(const string& s) {
    if (s.size() == 0 ) {
        return s;
    }
    
    vector<string> words;
    string word;
    for (int i = 0; i < s.size(); ++i) {
        if (s[i] != ' ') {
            word.push_back(s[i]);
        } else if (word.size() != 0) {
            append_word(words, word);
        }
    }
    if (word.size() != 0 ) {
        append_word(words, word);
    }
    string ans;
    for (int i = 0; i < words.size(); ++i) {
        ans.append(words[i]);
        ans.push_back(' ');
    }
    return ans;
}

int main() {
    cout << reverse("Let's have a happy life in company") << endl;

    return 0;
}



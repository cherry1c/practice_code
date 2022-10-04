#include <iostream>
#include <sstream>
#include <map>
#include <vector>


using namespace std;

void PrintSql01() {
    ostringstream oss;
    string start_time_mark("2022-07-06 00:00:00");
    int feed_id_mark = 12345678;

    oss << "select feed_id, feed_model, updated_at from square_feeds where ";
    oss << "updated_at > \"" << start_time_mark << "\" or ";
    oss << "( updated_at = \"" << start_time_mark << "\" and feed_id > " << feed_id_mark << ") ";
    oss << "order by updated_at, feed_id limit 50";
    cout << oss.str() << endl;
}

void PrintSql02() {
    ostringstream oss;
    oss << "select feed_id, " << "status" << ", " << "pre_status" << " from index_feeds where feed_id in ( ";
    oss << "1,2,3";
    oss << ") and is_deleted = 0 order by feed_id";
    cout << oss.str() << endl;
}

void PrintSql03() {
    ostringstream oss;
    oss << "select feed_id, feed_time from index_feeds where author_id = " << 123 << " and is_deleted=0 " << " and feed_id < " << 12345678;

    oss << " and ( (source = " << 0 << " and " << "status" << " in (" << "1, 3" << "))";
    oss << " or (source = ";
    oss << 1;
    oss << " and " << "status" << " in (1, 3)) ) ";

    
    oss << " order by feed_id desc limit " << 500;
    cout << oss.str() << endl;
}


void printMap() {
    map<int, vector<int>> m;
    m[0] = {1,2,3,4};
    m[1] = {6,7,8,9};
    for (auto it = m.begin(); it != m.end(); ++it) {
        if ((it->second).empty()) {
            continue;
        }
        cout << it->first << ": ";
        cout << (it->second)[0];
        for (int i = 1; i != (it->second).size(); ++i) {
            cout << ", " << (it->second)[i];
        }
        cout << endl;
    }
}

void split(const string& s, vector<string>& tokens, const string& delimiters = " ")
{
    string::size_type lastPos = s.find_first_not_of(delimiters, 0);
    string::size_type pos = s.find_first_of(delimiters, lastPos);
    while (string::npos != pos || string::npos != lastPos) {
        tokens.emplace_back(s.substr(lastPos, pos - lastPos));
        lastPos = s.find_first_not_of(delimiters, pos);
        pos = s.find_first_of(delimiters, lastPos);
    }
}

void testStringSplit() {
    string s("123");
    vector<string> v;
    split(s, v, ",");
    for (auto it = v.begin(); it != v.end(); ++it) {
        cout << *it << " ";
    }
    cout << endl;
}


void testStringToint() {
    int a;
    string value(" 123 ");
    stringstream ss;
    ss.str(value);
    ss >> a;
    cout << a << endl;
}

template<typename T>
void printVector(const vector<T>& v) {
    auto it = v.begin();
    cout << *it;
    ++it;
    for (; it != v.end(); ++it) {
        cout << ", " << *it;
    }
    cout << endl;
}

// 去掉字符串收尾空格
static void ClearHeadTailSpace(vector<string> &v) 
{
    if (v.empty())
    {
        return;
    }
    for (auto str = v.begin(); str != v.end(); ++str)
    {
        if (str->empty())   
        {  
            continue; 
        }  

        str->erase(0,str->find_first_not_of(" "));
        str->erase(str->find_last_not_of(" ") + 1);
    }
    
    return;  
}



int main() {
    vector<string>  s{"123   ", "    456", "789"};
    printVector(s);
    ClearHeadTailSpace(s);
    printVector(s);

    return 0;
}
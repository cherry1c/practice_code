#include <iostream>
#include <sstream>


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


int main() {
    PrintSql03();

    return 0;
}
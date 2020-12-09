#include <cstdio>
#include <string>

#include "rocksdb/db.h"
#include "rocksdb/slice.h"
#include "rocksdb/options.h"

using namespace std;
using namespace rocksdb;

const std::string PATH = "/home/yixingwei/Share/RocksDB/Data/temp"; //rocksDB的数据存储目录绝对路径

int main(){
    DB* db;
    Options options;
    options.create_if_missing = true;
    Status status = DB::Open(options, PATH, &db);
    assert(status.ok());
    Slice key("test01");
    Slice value("success");
    
    std::string get_value;
    status = db->Put(WriteOptions(), key, value);
    if(status.ok()){
        status = db->Get(ReadOptions(), key, &get_value);
        if(status.ok()){
            printf("value is %s\n", get_value.c_str());
        }else{
            printf("get failed\n"); 
        }
    }else{
        printf("put failed\n");
    }

    delete db;
}


#include <hashtable.h>
#include <iostream>
#include "HashTable.h"

int main(int argv, char **argc)
{

    HashTable<std::string, std::string> hash_table(30);
    hash_table["rustam"]="29";
    hash_table["ur"]="2";
    hash_table["ru"]="23";
    hash_table["rustam"]="30";
    hash_table["afat"]="76";
    hash_table["azada"]="53";
    hash_table["fuad"]="56";
    hash_table["elmira"]="30";

    hash_table.print();

    return 0;
}
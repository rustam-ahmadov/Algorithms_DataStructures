#pragma once




class String {

public:

	String();
	String(const char*) ;
	String(const String&);
	
	int getLen();
	char* getText();
	
	void setStr(const char*);
	
	String& operator=(const String&other);
	
	bool operator==(const String& other);
	bool operator!=(const String& other);

	String operator+(const String&) ;
	
	void print();
	
	~String();

private:
	char* str;
	int length=0;
};
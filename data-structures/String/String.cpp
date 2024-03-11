#include"String.h"
#include<iostream>

String::String() {
	str = nullptr;
}


String::String(const char*str) {
	length = strlen(str);
	this->str = new char[length+1];
	strcpy_s(this->str,length+1,str);

}


String::String(const String& other) :length(other.length) {

	str = new char[length + 1];
	strcpy_s(str, length + 1, other.str);
}

char * String::getText() {
	return str;
}

void String::setStr(const char*str) {
	delete[]this->str;
	length = strlen(str);
	this->str = new char[length + 1];
	strcpy_s(this->str, length + 1, str);
}


int String::getLen() {
	return length;

}


String& String::operator=(const String&other) {
	length = other.length;
	if (str!=nullptr)
	{
		delete[]str;
		str = new char[length + 1];
		strcpy_s(str, length + 1, other.str);
	}
	else
	{
		str = new char[length + 1];
		strcpy_s(str, length + 1, other.str);
	}
	return *this;
}

bool String::operator==(const String& other) {

	if (length==other.length)
	{
		for (int i = 0; i < length; i++)
		{
			if (str[i]!=other.str[i])
			{
				return false;
			}
		}
		return true;
	}
	return false;
}

bool String::operator!=(const String& other) {

	if (length!=other.length)
	{
	  return true;
		
	}

	for (int i = 0; i < length; i++)
	{
		if (str[i] != other.str[i])
		{
			return true;
		}
	}
	return false;

}


String String::operator+(const String& other) {
	String newStr;
	int dlinna = 0;
	for (size_t i = 0; ; i++)
	{
		if (other.str[i] == 0)
		{
			dlinna = i ;
			break;
		}
	}
	
	dlinna += length;
	newStr.str = new char[dlinna + 1];


	for (int i = 0; i < length; i++)
	{
		newStr.str[i] = this->str[i];
	}

	for (int i = length; i < dlinna; i++)
	{
		newStr.str[i] = other.str[i - length];
	}

	newStr.str[dlinna] = '\0';
	newStr.length = dlinna;
	
	

	return newStr;
}



void String:: print() { 
	std::cout << str<<"\n";
}

String::~String() {

	delete[]str;
}
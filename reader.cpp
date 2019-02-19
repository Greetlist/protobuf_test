#include "./cpp_out/person.pb.h"
#include <iostream>
#include <fstream>

using namespace std;

void printPersonInfo(::test::Person& person) {
	cout << "The person name is : " << person.name() << endl;
	cout << "The person id is : " << person.id() << endl;
	if (person.has_email()) {
		cout << "The person email is : " << person.email() << endl;
	}
	int size = person.phones_size();
	for (int i = 0; i < size; ++i) {
		::test::Person_PhoneNumber phone = person.phones(i);
		cout << "The " << i << "'s phone is : " << phone.number() << endl;
		cout << "The " << i << "'s phone type is : " << phone.type() << endl;
	}
}

int main(int argc, char** argv) {
	GOOGLE_PROTOBUF_VERIFY_VERSION;
	::test::Person person;
	fstream reader("./log.txt", ios::in | ios::binary);
	person.ParseFromIstream(&reader);
	printPersonInfo(person);
	::google::protobuf::ShutdownProtobufLibrary();
	return 0;
}

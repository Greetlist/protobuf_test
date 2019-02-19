#include "./cpp_out/person.pb.h"
#include <string>
#include <iostream>
#include <fstream>

using namespace std;

int main(int argc, char** argv) {
	GOOGLE_PROTOBUF_VERIFY_VERSION;
	::test::Person person;
	string name("Greetlist");
	int id = 123;
	string email("lrt12250914@sina.com");
	person.set_name(name);
	person.set_id(reinterpret_cast<::google::protobuf::int32>(id));
	person.set_email(email);
	for (int i = 0; i < 3; ++i) {
		::test::Person_PhoneNumber* curPhoneNumber = person.add_phones();
		curPhoneNumber->set_number("18585146697");
		curPhoneNumber->set_type(::test::Person_PhoneType_MOBILE);
	}
	fstream writer("./log.txt", ios::out | ios::binary | ios::trunc);
	person.SerializeToOstream(&writer);
	::google::protobuf::ShutdownProtobufLibrary();
	return 0;
}

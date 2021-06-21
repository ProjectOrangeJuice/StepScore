import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import '../../assets/constants.dart' as Constants;

Future<bool> register(String username) async {
  var url = Uri.http(Constants.API_URL, '/api/auth/register');

  // Await the http get response, then decode the json-formatted response.
  var response = await http.post(url, body: username);
  if (response.statusCode == 200) {
    return true;
  } else if (response.statusCode == 409) {
    return false;
  } else {
    throw ("bad response -> " + response.body);
  }
}

Future<bool> login(String username) async {
  var url = Uri.http(Constants.API_URL, '/api/auth/login');
  final _storage = FlutterSecureStorage();

  // Await the http get response, then decode the json-formatted response.
  var response = await http.post(url, headers: {"username": username});
  if (response.statusCode == 200) {
    await _storage.write(key: "username", value: username);
    return true;
  } else if (response.statusCode == 401) {
    return false;
  } else {
    throw ("bad response -> " + response.body);
  }
}

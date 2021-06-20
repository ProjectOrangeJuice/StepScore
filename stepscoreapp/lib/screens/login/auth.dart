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

import 'package:flutter/material.dart';
import 'auth.dart';

class WelcomeScreen extends StatelessWidget {
  final usernameController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Padding(
            padding: EdgeInsets.all(10),
            child: ListView(
              children: <Widget>[
                Container(
                    alignment: Alignment.center,
                    padding: EdgeInsets.all(10),
                    child: Text(
                      'Login',
                      style: TextStyle(
                          color: Colors.orange,
                          fontWeight: FontWeight.w500,
                          fontSize: 30),
                    )),
                Container(
                  padding: EdgeInsets.all(10),
                  child: TextField(
                    controller: usernameController,
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'User Name',
                    ),
                  ),
                ),
                Container(
                    child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: <Widget>[
                    ElevatedButton(
                      child: Text('Login'),
                      style: ElevatedButton.styleFrom(primary: Colors.orange),
                      onPressed: () {
                        //   print(nameController.text);
                        //  print(passwordController.text);
                      },
                    ),
                    ElevatedButton(
                      child: Text('Register'),
                      onPressed: () {
                        register(usernameController.text)
                            .then((reg) {})
                            .catchError((e) {
                          showAlertDialog(
                              context, "Error registering", e.toString());
                          print(e);
                        });

                        print(usernameController.text);
                        //  print(passwordController.text);
                      },
                    )
                  ],
                ))
              ],
            )));
  }
}

showAlertDialog(BuildContext context, String title, String content) {
  // set up the button
  Widget okButton = FlatButton(
    child: Text("OK"),
    onPressed: () {
      Navigator.of(context).pop();
    },
  );

  // set up the AlertDialog
  AlertDialog alert = AlertDialog(
    title: Text(title),
    content: Text(content),
    actions: [
      okButton,
    ],
  );
  // show the dialog
  showDialog(
    context: context,
    builder: (BuildContext context) {
      return alert;
    },
  );
}

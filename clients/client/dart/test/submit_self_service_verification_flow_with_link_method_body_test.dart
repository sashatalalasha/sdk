import 'package:test/test.dart';
import 'package:ory_client/ory_client.dart';

// tests for SubmitSelfServiceVerificationFlowWithLinkMethodBody
void main() {
  final instance = SubmitSelfServiceVerificationFlowWithLinkMethodBodyBuilder();
  // TODO add properties to the builder and call build()

  group(SubmitSelfServiceVerificationFlowWithLinkMethodBody, () {
    // Sending the anti-csrf token is only required for browser login flows.
    // String csrfToken
    test('to test the property `csrfToken`', () async {
      // TODO
    });

    // Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
    // String email
    test('to test the property `email`', () async {
      // TODO
    });

    // Method supports `link` only right now.
    // String method
    test('to test the property `method`', () async {
      // TODO
    });

  });
}

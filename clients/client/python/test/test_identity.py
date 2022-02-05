"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v0.0.1-alpha.61
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.from_ory_client_model_identity_credentials_import_identity_credentials import FromOryClientModelIdentityCredentialsImportIdentityCredentials
from ory_client.model.from_ory_client_model_identity_state_import_identity_state import FromOryClientModelIdentityStateImportIdentityState
from ory_client.model.from_ory_client_model_recovery_address_import_recovery_address import FromOryClientModelRecoveryAddressImportRecoveryAddress
from ory_client.model.from_ory_client_model_verifiable_identity_address_import_verifiable_identity_address import FromOryClientModelVerifiableIdentityAddressImportVerifiableIdentityAddress
from ory_client.model.globals_identity_credentials_identity_credentials import GlobalsIdentityCredentialsIdentityCredentials
from ory_client.model.globals_identity_state_identity_state import GlobalsIdentityStateIdentityState
from ory_client.model.globals_recovery_address_recovery_address import GlobalsRecoveryAddressRecoveryAddress
from ory_client.model.globals_verifiable_identity_address_verifiable_identity_address import GlobalsVerifiableIdentityAddressVerifiableIdentityAddress
globals()['from ory_client.model.identity_credentials import IdentityCredentials'] = from ory_client.model.identity_credentials import IdentityCredentials
globals()['from ory_client.model.identity_state import IdentityState'] = from ory_client.model.identity_state import IdentityState
globals()['from ory_client.model.recovery_address import RecoveryAddress'] = from ory_client.model.recovery_address import RecoveryAddress
globals()['from ory_client.model.verifiable_identity_address import VerifiableIdentityAddress'] = from ory_client.model.verifiable_identity_address import VerifiableIdentityAddress
globals()['globals()['IdentityCredentials'] = IdentityCredentials'] = globals()['IdentityCredentials'] = IdentityCredentials
globals()['globals()['IdentityState'] = IdentityState'] = globals()['IdentityState'] = IdentityState
globals()['globals()['RecoveryAddress'] = RecoveryAddress'] = globals()['RecoveryAddress'] = RecoveryAddress
globals()['globals()['VerifiableIdentityAddress'] = VerifiableIdentityAddress'] = globals()['VerifiableIdentityAddress'] = VerifiableIdentityAddress
from ory_client.model.identity import Identity


class TestIdentity(unittest.TestCase):
    """Identity unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testIdentity(self):
        """Test Identity"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Identity()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()

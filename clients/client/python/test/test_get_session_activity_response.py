"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.2.17
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.session_activity_datapoint import SessionActivityDatapoint
globals()['SessionActivityDatapoint'] = SessionActivityDatapoint
from ory_client.model.get_session_activity_response import GetSessionActivityResponse


class TestGetSessionActivityResponse(unittest.TestCase):
    """GetSessionActivityResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testGetSessionActivityResponse(self):
        """Test GetSessionActivityResponse"""
        # FIXME: construct object with mandatory attributes with example values
        # model = GetSessionActivityResponse()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
